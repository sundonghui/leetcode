package test

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	list []string
	preList map[rune]struct{}
)

func search(list []string, key rune) []int {	
	result := make([]int, 0, len(list))
	for i, item := range list {
		for _, v := range item {
			if v == key {
				result = append(result, i)
			}
		}
	}
	return result
}

func pre(list []string) (ans map[rune][]int) {
	for i, item := range list {
		for _, v := range item {
			if !isHave(ans[v], i) {
				ans[v] = append(ans[v], i)
			}			
		}
	}
	return
}

func isHave(list []int, key int ) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

// a,b,c,d,e,f
// index: a,b,c,d
// search index: a,b,d
// explain ref

// uuid byte blob
// string
// uuid: 8 byte 合同的唯一标示 随机生成，顺序是不固定的，一定不会重复
// 引入 时间戳+uuid 唯一索引
// 

type cal struct {
	state int
	*sync.Mutex
}


type cal1 struct {
	
}

type lock struct {
	m chan struct{}
}

func (f *lock) locked() {
	f.m<-struct{}{}
	return
}

func (f *lock) unlocked() {
	<-f.m
	return
}

func newLock() *lock {
	return &lock{m:make(chan struct{}, 1)}
}

func bufferChan() int {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 10)
	sum := 0
	var mutex sync.Mutex // 创建互斥锁
	for i := 0 ; i <= 10000; i++  {
		ch<-struct{}{}
		wg.Add(1)		
		go func(i int) {
			fmt.Printf("goroutine num is: %d \n", runtime.NumGoroutine())
			mutex.Lock()
			sum+=i
			mutex.Unlock()
			<-ch
			wg.Done()
		}(i)		
	}
	wg.Wait()
	return sum
}