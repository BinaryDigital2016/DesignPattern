package lazySingleton

import "sync"

type LazySingleton struct {
}

var l *LazySingleton
var mu = sync.Mutex{}

func GetInstance() *LazySingleton {
	mu.Lock()
	defer mu.Unlock()
	if l == nil {
		l = new(LazySingleton)
	}

	return l
}
