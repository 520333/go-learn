package info

import (
	"context"
	"open-devops/src/common"
	"open-devops/src/models"
	"open-devops/src/modules/agent/rpc"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func CollectBaseInfo(cli *rpc.RpcCli, logger log.Logger) {
	var (
		err  error
		sn   string
		cpu  string
		mem  string
		disk string
	)
	snShellCloud := `TOKEN=$(curl -s -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 60" --max-time 2); if [ -n "$TOKEN" ]; then curl -s -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/instance-id --max-time 2; else curl -s http://169.254.169.254/latest/meta-data/instance-id --max-time 2; fi`
	//snShellCloud := `curl -s http://169.254.169.254/a/meta-data/instance-id`
	//snShellHost := `dmidecode -s system-serial-number |tail -n 1|tr -d "\n"`
	snShellHost := `cat /sys/class/dmi/id/product_serial | tr -d "\n"`

	cpuShell := `cat /proc/cpuinfo |grep processor |wc -l| tr -d "\n"`
	memShell := `cat /proc/meminfo |grep MemTotal |awk '/MemTotal/{printf "%d", $2/1024/1024 + 0.99}' /proc/meminfo`
	diskShell := `df -m  |grep '/dev/' |grep -v '/var/lib' |grep -v tmpfs |awk '{sum +=$2};END{printf "%d",sum/1024}'`

	sn, err = common.ShellCommand(snShellCloud)
	if err != nil || sn == "" {
		sn, err = common.ShellCommand(snShellHost)
	}
	level.Info(logger).Log("msg", "CollectBaseInfo", "sn", sn)

	cpu, err = common.ShellCommand(cpuShell)
	if err != nil {
		level.Error(logger).Log("msg", "cpuShell.error", "shell", cpuShell, "err", err)
	}
	level.Info(logger).Log("msg", "CollectBaseInfo", "cpu", cpu)

	mem, err = common.ShellCommand(memShell)
	if err != nil {
		level.Error(logger).Log("msg", "memShell.error", "shell", memShell, "err", err)
	}
	level.Info(logger).Log("msg", "CollectBaseInfo", "mem", mem)

	disk, err = common.ShellCommand(diskShell)
	if err != nil {
		level.Error(logger).Log("msg", "memShell.error", "diskShell", diskShell, "err", err)
	}
	level.Info(logger).Log("msg", "CollectBaseInfo", "disk", disk)

	ipAddr := common.GetLocalIp()
	hostName := common.GetHostName()

	hostObj := models.AgentCollectInfo{
		SN:       sn,
		CPU:      cpu,
		Mem:      mem,
		Disk:     disk,
		IpAddr:   ipAddr,
		HostName: hostName,
	}
	cli.HostInfoReport(hostObj)
}

func TickerInfoCollectAndReport(cli *rpc.RpcCli, ctx context.Context, logger log.Logger) error {
	ticker := time.NewTicker(5 * time.Second)

	level.Info(logger).Log("msg", "TickerInfoCollectAndReport.start")
	CollectBaseInfo(cli, logger)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			level.Info(logger).Log("msg", "receive_quit_signal_and_quit")
			return nil
		case <-ticker.C:
			CollectBaseInfo(cli, logger)
		}
	}

}
