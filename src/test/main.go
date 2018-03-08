package main 

import(
	"test/go_web"		//自定义web监听接口
	"test/go_log"		//自定义日志
	"test/go_mysql"
	"test/go_tcp"
	"test/go_config"
	"fmt"
)

func main(){
	go_log.Error("main", int(100000000000000))
	fmt.Println(go_config.Cfg_store(1))
	fmt.Println(go_config.Cfg_area_game(1001))
	fmt.Println(go_config.Cfg_hu_rate(1))
	go_mysql.InitMysql()
	go go_tcp.Init_tcp()
	go_web.Init_web()
}