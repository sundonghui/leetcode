package buildingh2o

import (
	"fmt"
	"sync"
)

// H2O结构体包含：
// 互斥锁mu、
// 计数器hCount和oCount、
// 条件变量oCond以及通道hChan。
type H2O struct {
	mu     sync.Mutex
	hCount int
	oCount int
	oCond  *sync.Cond
	hChan  chan struct{}
}

// NewH2O初始化结构体并创建通道和条件变量。
func NewH2O() *H2O {
	h2o := &H2O{
		hChan: make(chan struct{}, 2),
	}
	h2o.oCond = sync.NewCond(&h2o.mu)
	return h2o
}

// 增加hCount，若满足条件（有足够H和O）则唤醒等待的O线程。
// 等待从通道hChan接收信号后调用releaseHydrogen。
func (h2o *H2O) Hydrogen(releaseHydrogen func()) {
	h2o.mu.Lock()
	h2o.hCount++
	// 检查是否满足条件并唤醒O线程
	if h2o.hCount >= 2 && h2o.oCount >= 1 {
		h2o.oCond.Signal()
	}
	h2o.mu.Unlock()

	// 等待允许通过
	<-h2o.hChan
	releaseHydrogen()
}

// 增加oCount，并等待直到有足够的H线程到达。
// 消耗两个H和一个O，通过通道发送两个信号允许H线程继续执行。
func (h2o *H2O) Oxygen(releaseOxygen func()) {
	h2o.mu.Lock()
	h2o.oCount++
	// 等待直到有足够的H
	for h2o.hCount < 2 {
		h2o.oCond.Wait()
	}
	// 消耗两个H和一个O
	h2o.hCount -= 2
	h2o.oCount--
	h2o.mu.Unlock()

	// 允许两个H通过
	h2o.hChan <- struct{}{}
	h2o.hChan <- struct{}{}
	releaseOxygen()
}

// 模拟输入字符串生成H和O线程，验证输出是否为有效组合。
func main() {
	var result []byte
	var mu sync.Mutex
	water := "OHOOHHHH"
	h2o := NewH2O()

	var wg sync.WaitGroup
	wg.Add(len(water))

	for _, c := range water {
		go func(c rune) {
			defer wg.Done()
			if c == 'H' {
				h2o.Hydrogen(func() {
					mu.Lock()
					result = append(result, 'H')
					mu.Unlock()
				})
			} else {
				h2o.Oxygen(func() {
					mu.Lock()
					result = append(result, 'O')
					mu.Unlock()
				})
			}
		}(c)
	}

	wg.Wait()
	fmt.Println(string(result)) // 示例输出: HHOHHO（可能不同，但有效）
}
