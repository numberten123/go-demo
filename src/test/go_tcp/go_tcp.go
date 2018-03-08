package go_tcp

import (
    "fmt"
    "net"
    "test/go_config"
    "test/go_log"
    "test/go_proto"
    "github.com/golang/protobuf/proto"
)

func Init_tcp() {
    addrInfo := fmt.Sprintf(`%v:%v`,go_config.GetValue("tcp_ip"), go_config.GetValue("tcp_port"))
    ln, err := net.Listen("tcp", addrInfo)
    checkError(err)
    for {
        conn, err := ln.Accept()
        checkError(err)
        go handleConnection(conn)
    }
}

func checkError(err error) {
    if err != nil {
        // fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        go_log.Error("tcp Fatal error: %s", err.Error())
    }
}

func handleConnection(conn net.Conn){
    buffer := make([]byte, 1024)
    conn.Read(buffer)
    recvmsg := &go_proto.M_1208Tos{}
    err := proto.Unmarshal(buffer, recvmsg)
    checkError(err)
    msg := recvmsg.GetToServer() 
    fmt.Println("s-recv",recvmsg)
    sendmsg := &go_proto.M_1208Toc{
        ToClient:proto.String("return:"+msg),
    }
    senddata, err := proto.Marshal(sendmsg)
    conn.Write(senddata)
    conn.Close()
}