package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Transfer(to *Account, amout int) {
	a.mu.Lock()
	fmt.Printf("锁定账户A，余额：%d\n", a.balance)
	time.Sleep(100 * time.Millisecond)

	to.mu.Lock()
	fmt.Printf("锁定账户B，余额：%d\n", to.balance)

	a.balance -= amout
	to.balance += amout

	to.mu.Unlock()
	a.mu.Unlock()
}

func main() {
	// 科夫曼条件
	// 互斥：sync.Mutex 同一时间只能一个 goroutine 持有
	// 占有且等待：G1 持有 A，同时等待 B
	// 不可剥夺：G2 不释放 B，G1 不能强行拿走
	// 循环等待：G1 等 B，G2 等 A

	accountA := &Account{balance: 1000}
	accountB := &Account{balance: 500}

	// 转账A->B
	go accountA.Transfer(accountB, 100)

	time.Sleep(100 * time.Millisecond)

	// 转账B->A
	go accountB.Transfer(accountA, 50)

	time.Sleep(2 * time.Second)
	fmt.Printf("最终余额 -A：%d，B：%d\n", accountA.balance, accountB.balance)
}
