package onceSingleton

import "sync"

type OnceSingleton struct {
}

var o *OnceSingleton
var once sync.Once

func GetInstance() *OnceSingleton {
	once.Do(func() {
		o = new(OnceSingleton)
	})

	return o
}
