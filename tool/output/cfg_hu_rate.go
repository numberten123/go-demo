// !!! DO NOT EDIT !!!
// auto generated from Excel files
package go_config

import(
	"test/go_include"
	"test/go_log"
)

func Cfg_hu_rate(key uint32) (go_include.Cfg_hu_rate) {
	switch key{
		case 1:
			return go_include.Cfg_hu_rate{
				Type:1,
				Rate:1,
				Min_gen:0,
			}
		case 2:
			return go_include.Cfg_hu_rate{
				Type:2,
				Rate:2,
				Min_gen:0,
			}
		case 3:
			return go_include.Cfg_hu_rate{
				Type:3,
				Rate:4,
				Min_gen:0,
			}
		case 4:
			return go_include.Cfg_hu_rate{
				Type:4,
				Rate:4,
				Min_gen:0,
			}
		case 5:
			return go_include.Cfg_hu_rate{
				Type:5,
				Rate:4,
				Min_gen:0,
			}
		case 6:
			return go_include.Cfg_hu_rate{
				Type:6,
				Rate:4,
				Min_gen:0,
			}
		case 7:
			return go_include.Cfg_hu_rate{
				Type:7,
				Rate:8,
				Min_gen:0,
			}
		case 8:
			return go_include.Cfg_hu_rate{
				Type:8,
				Rate:8,
				Min_gen:0,
			}
		case 9:
			return go_include.Cfg_hu_rate{
				Type:9,
				Rate:8,
				Min_gen:1,
			}
		case 10:
			return go_include.Cfg_hu_rate{
				Type:10,
				Rate:16,
				Min_gen:0,
			}
		case 11:
			return go_include.Cfg_hu_rate{
				Type:11,
				Rate:16,
				Min_gen:0,
			}
		case 12:
			return go_include.Cfg_hu_rate{
				Type:12,
				Rate:16,
				Min_gen:0,
			}
		case 13:
			return go_include.Cfg_hu_rate{
				Type:13,
				Rate:16,
				Min_gen:0,
			}
		case 14:
			return go_include.Cfg_hu_rate{
				Type:14,
				Rate:32,
				Min_gen:0,
			}
		case 15:
			return go_include.Cfg_hu_rate{
				Type:15,
				Rate:32,
				Min_gen:0,
			}
		case 16:
			return go_include.Cfg_hu_rate{
				Type:16,
				Rate:32,
				Min_gen:1,
			}
		case 17:
			return go_include.Cfg_hu_rate{
				Type:17,
				Rate:64,
				Min_gen:4,
			}
		case 18:
			return go_include.Cfg_hu_rate{
				Type:18,
				Rate:256,
				Min_gen:4,
			}
		default:
			go_log.Error("Cfg_store not match:", key)
			return go_include.Cfg_hu_rate{}
	}
}
