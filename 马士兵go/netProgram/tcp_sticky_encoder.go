package netProgram

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// Encoder 定义编码器 发送端
type Encoder struct {
	w io.Writer //编码结束后写入的目标
}

// NewEncoder 编码器函数
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode 将结果写入到w io.Writer
func (enc *Encoder) Encode(message string) error {
	// 1获取message长度
	l := int32(len(message))
	buf := new(bytes.Buffer)
	// 2在数据包中写入长度
	if err := binary.Write(buf, binary.LittleEndian, &l); err != nil {
		return err
	}
	// 3将数据主题Body写入
	if _, err := buf.Write([]byte(message)); err != nil {
		return err
	}

	// 4利用io.Writer发送数据
	if _, err := enc.w.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

// Decoder 定义解码器 接收端
type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}
func (dec *Decoder) Decode(message *string) error {
	// 1读取前4个字节 header
	header := make([]byte, 4)
	hn, err := dec.r.Read(header)
	if err != nil {
		return err
	}
	if hn != 4 {
		return errors.New("header is not enough")
	}
	// 2将前四个字节转换为int32类型
	var l int32
	buf := bytes.NewBuffer(header)
	if err := binary.Read(buf, binary.LittleEndian, &l); err != nil {
		return err
	}
	// 3读取body
	body := make([]byte, l)
	bn, err := dec.r.Read(body)
	if err != nil {
		return err
	}
	if bn != int(l) {
		return errors.New("body is not enough")
	}
	// 4设置Message
	*message = string(body)
	return nil
}
