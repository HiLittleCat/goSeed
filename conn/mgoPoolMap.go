package conn

func init() {
	mgoPools = make(map[string]MgoPool)
}

var mgoPools map[string]MgoPool

func MgoSet(key string, p *MgoPool) {
	mgoPools[key] = *p
}

func GetMgoPool(key string) MgoPool {
	return mgoPools[key]
}
