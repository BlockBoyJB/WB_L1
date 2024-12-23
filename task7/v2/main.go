package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// В этом решении используется мапа из пакета sync. Непонятно, зачем я решил это сделать, если под капотом у нее тот же мьютекс
	// Идея - хранить в мапе по значению *int64 и атомарно менять его
	m := &sync.Map{}

	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(m *sync.Map, key int) {
			defer wg.Done()
			// в мапе ключ - int, значение - *int64
			v, _ := m.LoadOrStore(key, new(int64)) // если у нас значения по ключу нет, то записываем новое значение-ссылку на int64 в памяти
			atomic.AddInt64(v.(*int64), 1)         // атомарно увеличиваем значение
		}(m, i%10)
	}

	wg.Wait()
	for i := 0; i < 10; i++ {
		v, _ := m.Load(i)
		fmt.Printf("%d:%d ", i, *((v).(*int64))) // вывод как в обычной мапе хотел, но значение по ключу выглядит страшно...
	}
}