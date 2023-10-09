package features

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type MapAccess struct {
	sync.Mutex
	access map[int]int
}

func (m *MapAccess) read(k int) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	_ = m.access[k]
}
func (m *MapAccess) write(k int, v int) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.access[k] = 2
}

func RunMutex() {

	m := MapAccess{
		access: map[int]int{},
	}
	var readOps uint64
	var writeOps uint64

	for r := 1; r < 100; r++ {
		go func() {
			for {
				key := rand.Intn(5)
				m.read(key)
				atomic.AddUint64(&readOps, 1)
			}
		}()
	}
	for r := 0; r < 10; r++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				m.write(key, val)
				atomic.AddUint64(&writeOps, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}
