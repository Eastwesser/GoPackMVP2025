// Напиши  реализацию InMemory кэша
package main

import "sync"

type Cache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

type InMemoryCache struct {
	data map[string]string
	//mu   sync.Mutex
	mu sync.RWMutex // Лучше обычного мьютекса sync.Mutex по скорости чтения и записи

}

// Так как это кэш, к нему будет конкурентный доступ, к нему будут стремиться обратиться много горутин
// Мапа неконкурентна на запись (на чтение ок) непотокобезопасная

//a := 3
//a++
// Три операции: вычитать, прибавить, записать

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		make(map[string]string), // инициализация мапы
		//sync.Mutex{},
		sync.RWMutex{},
	}
}

func (c *InMemoryCache) Set(k string, v string) {
	// "(c *InMemoryCache)" - ресивер
	// "Set(k string, v string)" - контракт

	// Чтобы был конкурентный доступ к мапе, используем мьютексы
	//c.mu.Lock()
	//c.data[k] = v
	//c.mu.Unlock()

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

func (c *InMemoryCache) Get(k string) (string, bool) {
	c.mu.RLock() // RLock() из RW мьютекса нужен для ускорения массового чтения горутинами
	// Не блокируем чтение, блокируем только запись
	defer c.mu.RUnlock() // если проект будет разрастаться, можно через defer открывать доступ

	data, ok := c.data[k] // ok сигнализирует, есль ли ключ в мапе или нет
	if ok {
		println("Cache hit %s", data)
	}
	return data, ok
	//return c.data[k]
	/*
		Почему нельзя return c.data[k]?
		Произойдет перегрузка функций:
			- первая (?) возвращает value,
			- вторая (?) И value, И флажок
	*/
}

func main() {
	//cache := &InMemoryCache{} // вызовет панику, так как мы не вызвали функцию NewInMemoryCache()
	cache := NewInMemoryCache()
	cache.Set("foo", "bar")
	cache.Set("baz", "qux")

	// Как синхроизировать выполнение горутин? WaitGroup!!!
	// Когда завершается main(), все горутины схлопываются, не будут выполнены
	wg := &sync.WaitGroup{}
	wg.Add(4)

	// когда запускается функция - определит планировщик. go означает, что мы кладем функцию в очередь
	go func() {
		defer wg.Done()
		cache.Set("foo", "upd_bar")
		println("1")
	}()

	go func() {
		defer wg.Done()
		cache.Set("baz", "upd_qux")
		println("2")
	}()

	go func() {
		defer wg.Done()
		cache.Get("foo")
		println("3")
	}()

	go func() {
		defer wg.Done()
		cache.Get("bar")
		println("4")
	}()

	wg.Wait() // пока счетчик не будет 0, мы выполняем горутины

	data, ok := cache.Get("foo")
	//if ok {
	//	fmt.Println(data)
	//} else {
	//	fmt.Println("Not found") // с else некрасиво, println = alias for fmt.Println
	//}

	if !ok {
		println("Key: Not found")
		return
	}
	println("Key: foo, Value: ", data)
}
