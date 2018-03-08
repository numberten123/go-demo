package go_web

import (
    "fmt"
    "log"
    "net/http"
    "test/go_config"
    "test/go_tcp"
)

const (
    anError = `error msg:%s`
)

type statistics struct {
    numbers []float64
    mean    float64
    median  float64
}

func Init_web() {
    http.HandleFunc("/order", dealOrder)
    http.HandleFunc("/send_msg", sendMsg)
    ip := go_config.GetValue("web_ip")
    port := go_config.GetValue("web_port")
    info := fmt.Sprintf(`%v:%v`,ip, port)
    if err := http.ListenAndServe(info, nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}

//订单回调接口
func dealOrder(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm() // Must be called before writing response
    if err != nil {
        fmt.Fprintf(writer, anError, err)
    } else {
        orderID := request.Form["order_id"]
        Amount := request.Form["amount"]
        result := finish_order(orderID, Amount)
        fmt.Fprint(writer, result)
    }
}

func sendMsg(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm() // Must be called before writing response
    if err != nil {
        fmt.Fprintf(writer, anError, err)
    } else {
        msg := request.Form["msg"]
        go_tcp.SendMsg(to_string(msg))   //发送信息
        fmt.Fprint(writer, "ok")
    }
}

//完成订单
func finish_order(orderID []string, Amount []string) (string) {
    return fmt.Sprintf(`order:%s, amount:%s`,orderID, Amount)
}

func to_string(msg []string) (string){
    var r string
    for _,v := range msg{
        r += v
    }
    return r
}