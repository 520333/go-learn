package models

import (
	"encoding/json"
	"fmt"
	"ginchat/utils"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

// Message 消息
type Message struct {
	gorm.Model
	FormId   int64  // 消息发送者id
	TargetId int64  // 消息接收者id
	Type     int    // 发送类型 1:私聊 2:群聊 3:广播
	Media    int    // 消息类型 文字 图片 音频
	Content  string // 消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int // 其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLock sync.RWMutex

// Chat 需要：发送者ID 接收者ID 消息类型 发送的内容 发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 1.获取参数并校验token
	//token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	isvalida := true
	conn, err := (&websocket.Upgrader{
		// token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.获取连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 3.用户关系
	// 4.userId跟node绑定 并加锁
	rwLock.Lock()
	clientMap[userId] = node
	rwLock.Unlock()
	// 5.完成消息发送逻辑
	go sendProc(node)
	// 6.完成消息接受逻辑
	go recvProc(node)
	sendMsg(userId, []byte("欢迎进入聊天系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("[ws] sendMsg >>> ", "msg:", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] recvProc <<<< ", string(data))
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
	fmt.Println("init goroutine: ")
}

// 完成UDP数据发送协程
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 50, 255),
		Port: 3000,
	})
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case data := <-udpsendChan:
			fmt.Println("udpSendProc: data", string(data))
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// 完成UDP数据接收协程
func udpRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpRecvProc data:", string(buf[0:n]))
		dispatch(buf[0:n])
	}

}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: // 私信
		fmt.Println("dispatch: data", string(data))
		sendMsg(msg.TargetId, data)
		//case 2: // 群发
		sendGroupMsg(msg.TargetId, data) // 发送的群ID，消息内容
		//case 3: // 广播
		//	sentAllMsg()
		//case 4:
		//	fallthrough
		//default:
	}
}
func sendGroupMsg(targetId int64, msg []byte) {
	fmt.Println("开始群发消息:")
	userIds := SearchUserByGroupId(uint(targetId))
	for i := 0; i < len(userIds); i++ {
		sendMsg(int64(userIds[i]), msg)
	}
}

func JoinGroup(userId uint, comId uint) (int, string) {
	contact := Contact{}
	contact.OwnerId = userId
	contact.TargetId = comId
	contact.Type = 2
	community := Community{}
	utils.DB.Where("id = ?", comId).Find(&community)
	if community.Name == "" {
		return -1, "该群没有找到"
	}
	utils.DB.Where("owner_id = ? and target_id = ? and type = 2", userId, comId).Find(&contact)
	if !contact.CreatedAt.IsZero() {
		return -1, "已添加过该群"
	}
	utils.DB.Create(&contact)
	return 0, "加群成功"
}

func sendMsg(userId int64, msg []byte) {
	fmt.Println("sendMsg >>>> userId", userId, "msg:", string(msg))
	rwLock.RLock()
	node, ok := clientMap[userId]
	rwLock.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
