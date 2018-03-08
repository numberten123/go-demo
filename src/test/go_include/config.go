package go_include

type Cfg_store struct {
	Id 			uint32
	Shop_type 	uint32
	Shop_num	uint32
	Pay			uint32
}

type Cfg_hu_rate struct {
	Type 		uint32
	Rate		uint32
	Min_gen		uint32
}

type Cfg_area_game struct {
	Area_id 	string
	Mahjong 	[]uint32
}