package go_config

import (  
    "encoding/json"
    "io/ioutil"
)

func GetValue(getKey string) (interface{}) {
    var result = map[string]string{}
    data, err := ioutil.ReadFile("../config/go.config")
    if err != nil{
        return "error"
    }
    err = json.Unmarshal(data, &result)
    if err != nil{
        return "error"
    }
    if value, found := result[getKey]; found{
        return value
    }else{
        return "error"
    }
}