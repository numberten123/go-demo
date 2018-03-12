// !!! DO NOT EDIT !!!
// auto generated from Excel files
package go_config

import(
	"test/go_include"
	"test/go_log"
)

func Cfg_store(key uint32) (go_include.Cfg_store) {
	var cfg_store = map[uint32]go_include.Cfg_store{
	    1:
		    go_include.Cfg_store{
				Id:1,
				Shop_type:1,
				Shop_num:60,
				Pay:6,
			},
	    2:
		    go_include.Cfg_store{
				Id:2,
				Shop_type:1,
				Shop_num:60,
				Pay:6,
			},
	    3:
		    go_include.Cfg_store{
				Id:3,
				Shop_type:1,
				Shop_num:60,
				Pay:6,
			},
	}
	if result, ok := cfg_store[key]; ok {
        return result
    }
    go_log.Error("Cfg_store not match:", key)
 	return go_include.Cfg_store{}
}
