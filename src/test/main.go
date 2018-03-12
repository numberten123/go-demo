package main 

import(
	"test/go_web"		//自定义web监听接口
	// "test/go_log"		//自定义日志
	"test/go_mysql"
	"test/go_tcp"
	"test/go_config"
	"fmt"
)

func main(){
	fmt.Println(go_config.Cfg_area_game(1001))
	fmt.Println(go_config.Cfg_hu_rate(18))
	fmt.Println(go_config.Cfg_hu_rate(19))
	go_mysql.InitMysql()
	go go_tcp.Init_tcp()
	go_web.Init_web()
}