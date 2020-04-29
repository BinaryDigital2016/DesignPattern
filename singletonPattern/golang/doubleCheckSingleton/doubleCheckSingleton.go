package doubleCheckSingleton

import (
	"sync"
)

type DoubleCheckSingleton struct {
}

var d *DoubleCheckSingleton
var mu sync.Mutex

func GetInstance() *DoubleCheckSingleton {
	if d == nil {
		mu.Lock()
		defer mu.Unlock()

		if d == nil {
			d = new(DoubleCheckSingleton)
		}
	}

	return d
}
