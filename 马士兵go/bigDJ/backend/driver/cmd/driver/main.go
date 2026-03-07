package main

import (
	"flag"
	"os"

	"driver/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "Driver"
	// Version is the version of the compiled software.
	Version string = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	id = Name + "-" + uuid.NewString()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(cs *conf.Service, logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	reg, err := initServiceRegistry(cs.Consul.Address)
	if err != nil {
		panic(err)
	}

	if err = initTracer(cs.Jaeger.Url); err != nil {
		panic(err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer func() {
		_ = c.Close()
	}()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Service, bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initServiceRegistry(address string) (*consul.Registry, error) {
	// 一 获取consul客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = address
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("new consul client failed, err:%v", err)
	}
	// 二 获取consul注册管理器
	reg := consul.New(consulClient)

	return reg, nil
}

// 初始化Tracer
// @param url string 指定jaeger的api接口 :14268/api/traces
func initTracer(url string) error {
	// 一建立 jaeger客户端，称之为：exporter导出器
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}

	// 2 创建 trace provider
	tracerProvider := traceSDK.NewTracerProvider(
		traceSDK.WithSampler(traceSDK.AlwaysSample()), // 采样设置
		traceSDK.WithBatcher(exporter),                // 使用exporter作为批处理程序
		traceSDK.WithResource(resource.NewSchemaless( // 将当前服务的信息，作为资源告知给TracerProvider
			semconv.ServiceNameKey.String(Name), // 必要配置
			attribute.String("exporter", url),
		)),
	)

	// 3 设置全局 tracer
	otel.SetTracerProvider(tracerProvider)

	return nil
}
