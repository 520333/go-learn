package netProgram

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

type Pool interface {
	Get() (net.Conn, error)  //获取池子中空闲连接
	Put(conn net.Conn) error //放回连接，非关闭
	Release() error          //释放连接池
	Len() int                //有效连接的长度
}

// PoolConfig 连接池配置
type PoolConfig struct {
	InitConnNum int           //初始化连接数，池初始化时的连接数
	MaxConnNum  int           //最大连接数，池中最多支持多少连接
	MaxIdleNum  int           //空闲连接数，池中最多有多少可用的连接
	IdleTimeout time.Duration //空闲连接超时时间，多久后空闲连接后会被释放
	Factory     ConnFactory   //连接工厂

}

// IdleConn 空闲连接类型
type IdleConn struct {
	conn    net.Conn  //连接本身
	putTime time.Time //放回池子时间，用于判断是否空闲超时
}

// TcpPool 连接池结构
type TcpPool struct {
	config        PoolConfig     //配置信息
	openingConNum int            //开放使用的连接数量
	idleList      chan *IdleConn //空闲的连接列表
	addr          string         //连接地址
	mu            sync.RWMutex   //并发安全锁
}

// ConnFactory 连接工厂接口
type ConnFactory interface {
	Factory(addr string) (net.Conn, error) //创建生产连接
	Close(net.Conn) error                  //关闭连接
	Ping(net.Conn) error                   //Ping测试
}

// TcpConnFactory TCP连接工厂类型
type TcpConnFactory struct{}

func (f *TcpConnFactory) Factory(addr string) (net.Conn, error) {
	//校验参数的合理性
	if addr == "" {
		return nil, errors.New("addr is empty")
	}
	//建立连接
	conn, err := net.DialTimeout("tcp", addr, time.Second*5)
	if err != nil {
		return nil, err
	}
	//返回
	return conn, nil
}
func (f *TcpConnFactory) Close(conn net.Conn) error {
	return conn.Close()
}
func (f *TcpConnFactory) Ping(conn net.Conn) error {
	return nil
}

const (
	defaultMaxConnNum  = 100
	defaultInitConnNum = 1
)

//const defaultInitConnNum = 100

// 创建TcpPool对象
func NewTcpPool(addr string, poolConfig PoolConfig) (*TcpPool, error) {
	// 1.校验参数
	if addr == "" {
		return nil, errors.New("addr is empty")
	}
	//校验工厂的存在
	if poolConfig.Factory == nil {
		return nil, errors.New("factory is not exists")
	}
	// 最大连接数
	if poolConfig.MaxConnNum == 0 {
		//a return错误
		//return nil, errors.New("max conn num is zero")
		//b 人为修改一个合理的
		poolConfig.MaxConnNum = defaultMaxConnNum
	}
	// 初始化连接数
	if poolConfig.InitConnNum == 0 {
		poolConfig.InitConnNum = defaultInitConnNum
	} else if poolConfig.InitConnNum > poolConfig.MaxConnNum {
		poolConfig.InitConnNum = poolConfig.MaxConnNum
	}
	// 合理化最大空闲连接数
	if poolConfig.MaxIdleNum == 0 {
		poolConfig.MaxIdleNum = poolConfig.InitConnNum
	} else if poolConfig.MaxIdleNum > poolConfig.MaxConnNum {
		poolConfig.MaxIdleNum = poolConfig.MaxConnNum
	}

	// 2.初始化TcpPool对象
	pool := TcpPool{
		config:        poolConfig,
		openingConNum: 0,
		idleList:      make(chan *IdleConn, poolConfig.MaxIdleNum),
		addr:          addr,
		mu:            sync.RWMutex{},
	}
	// 3.初始化连接 根据InitConnNum的配置来创建连接
	for i := 0; i < poolConfig.InitConnNum; i++ {
		conn, err := pool.config.Factory.Factory(addr)
		if err != nil {
			//释放可能存在的连接
			pool.Release()
			return nil, err
		}
		pool.idleList <- &IdleConn{conn: conn, putTime: time.Now()}
	}
	// 4.返回
	return &pool, nil
}

// Get TcpPool实现 Pool 接口
func (pool *TcpPool) Get() (net.Conn, error) {
	// 1.锁定
	pool.mu.Lock()
	defer pool.mu.Unlock()
	// 2.获取空闲连接若没有则创建连接
	for {
		select {
		// 获取空闲连接
		case idleConn, ok := <-pool.idleList:
			// 判断channel是否被关闭
			if !ok {
				return nil, errors.New("idle list was closed")
			}
			// 判断连接是否超时 pool.config.IdleTimeout,idleConn.putTime
			if pool.config.IdleTimeout > 0 { //设置了超时时间
				//putTime + TimeOut是否在 now之前
				if idleConn.putTime.Add(pool.config.IdleTimeout).Before(time.Now()) {
					//关闭连接 继续查找下一个连接
					_ = pool.config.Factory.Close(idleConn.conn)
					continue
				}
			}
			//判断连接是否可用
			if err := pool.config.Factory.Ping(idleConn.conn); err != nil {
				//ping失败连接不可用
				_ = pool.config.Factory.Close(idleConn.conn)
				continue
			}
			log.Println("get conn from Idle")
			// 找到了可用的空闲连接
			pool.openingConNum++
			return idleConn.conn, nil
		//创建连接
		default:
			//判断是否还可以继续创建
			if pool.openingConNum >= pool.config.MaxConnNum {
				return nil, errors.New("max opening connection exceeded")
			}
			//创建连接
			conn, err := pool.config.Factory.Factory(pool.addr)
			if err != nil {
				return nil, err
			}
			log.Println("get conn from Factory")
			pool.openingConNum++
			return conn, nil
		}
	}
}
func (pool *TcpPool) Put(conn net.Conn) error {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if conn == nil {
		return errors.New("connection is not exists")
	}
	if pool.idleList == nil {
		_ = pool.config.Factory.Close(conn)
		return errors.New("idle list is not exists")
	}
	select {
	case pool.idleList <- &IdleConn{conn: conn, putTime: time.Now()}:
		pool.openingConNum--
		return nil
	default:
		_ = pool.config.Factory.Close(conn)
		return nil
	}
}
func (pool *TcpPool) Release() error {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if pool.idleList == nil {
		return nil
	}
	close(pool.idleList)
	for idleConn := range pool.idleList {
		_ = pool.config.Factory.Close(idleConn.conn)
	}
	return nil
}
func (pool *TcpPool) Len() int {
	return len(pool.idleList)
}
