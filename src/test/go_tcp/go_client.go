package go_tcp

import (
    "fmt"
    "net"
    "test/go_config"
    "test/go_proto"
    "github.com/golang/protobuf/proto"
)

func SendMsg(msg string) (string)  {
    addrInfo := fmt.Sprintf(`%v:%v`,go_config.GetValue("tcp_ip"), go_config.GetValue("tcp_port"))
    tcpAddr, err := net.ResolveTCPAddr("tcp4", addrInfo) //获取一个TCP地址信息,TCPAddr
    checkError(err)
    conn, err := net.DialTCP("tcp", nil, tcpAddr) //创建一个TCP连接:TCPConn
    checkError(err)
    sendmsg := &go_proto.M_1208Tos{
        ToServer:proto.String(msg),
    }
    senddata, err := proto.Marshal(sendmsg)
    checkError(err)
    fmt.Println("c-send",sendmsg)
    _, err = conn.Write(senddata)
    checkError(err)
    buffer := make([]byte, 1024)
    conn.Read(buffer)
    returnmsg := &go_proto.M_1208Toc{}
    err = proto.Unmarshal(buffer, returnmsg)
    fmt.Println("c-recv",returnmsg)
    return *returnmsg.ToClient
}