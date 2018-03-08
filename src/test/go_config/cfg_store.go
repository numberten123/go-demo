// !!! DO NOT EDIT !!!
// auto generated from Excel files
package go_config

import(
	"test/go_include"
	"test/go_log"
)

func Cfg_store(key uint32) (go_include.Cfg_store) {
	switch key{
		case 1:
			return go_include.Cfg_store{
				Id:1,
				Shop_type:1,
				Shop_num:60,
				Pay:6,
			}
		case 2:
			return go_include.Cfg_store{
				Id:1,
				Shop_type:1,
				Shop_num:120,
				Pay:12,
			}
		default:
			go_log.Info("Cfg_store not match:", key)
			return go_include.Cfg_store{}
	}
}
