package go_mysql

import (  
    "database/sql"  
    "fmt"
    "test/go_config"
    _ "github.com/go-sql-driver/mysql"
)  
  
func InitMysql() {
    info := fmt.Sprintf(`%v:%v@tcp(%v)/%v?charset=%v`,
        go_config.GetValue("dbusername"), 
        go_config.GetValue("dbpassword"),
        go_config.GetValue("dbhostsip"), 
        go_config.GetValue("dbname"),"utf8")
    db, err := sql.Open("mysql", info)
    checkErr(err)  
    rows, err := db.Query("SELECT role_id,name,icon FROM role where role_id<10")  
    checkErr(err)
    for rows.Next() {  
        var role_id int  
        var name string  
        var icon string  
        err = rows.Scan(&role_id, &name, &icon)  
        checkErr(err)  
        fmt.Println(role_id)  
        fmt.Println(name)  
        fmt.Println(icon)  
    }
    db.Close()
}  
  
func checkErr(err error) {  
    if err != nil {  
        panic(err)  
    }  
}  