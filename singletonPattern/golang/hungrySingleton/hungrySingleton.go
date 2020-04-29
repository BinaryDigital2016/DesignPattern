package hungrySingleton

type HungrySingleton struct {
}

var hungry = new(HungrySingleton)

func GetInstance() *HungrySingleton {
	return hungry
}
