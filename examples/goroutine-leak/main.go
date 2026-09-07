package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type workItem int
type workResult int

func processWorkItem(w workItem) (workResult, error) {
	time.Sleep(10 * time.Millisecond)
	if w == 5 {
		return 0, errors.New("simulated error")
	}
	return workResult(w * 2), nil
}

type result struct {
	res workResult
	err error
}

func processWorkItems(ws []workItem) ([]workResult, error) {
	// 使用缓冲区解决 goroutine 内存泄漏
	ch := make(chan result, len(ws))
	for _, w := range ws {
		go func() {
			res, err := processWorkItem(w)
			ch <- result{res, err}
		}()
	}

	var results []workResult
	for range len(ws) {
		r := <-ch
		if r.err != nil {
			return nil, r.err
		}
		results = append(results, r.res)
	}
	return results, nil
}

func main() {
	// Start pprof server
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Repeatedly trigger the leak
	for {
		items := []workItem{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		_, err := processWorkItems(items)
		if err != nil {
			log.Printf("Error processing items: %v", err)
		}

		time.Sleep(time.Second)
	}
}

func DoubleSend() {
	ch := make(chan any)
	go func(err error) {
		if err != nil {
			// In case of an error, send nil.
			ch <- nil
			// Return statement is missing
		}
		// Otherwise, continue with normal behaviour.
		// This send is still executed, which causes a leak in the error case.
		ch <- struct{}{}
	}(fmt.Errorf("error"))
	// Receive only one message
	<-ch
}

func EarlyReturn(err error) {
	// 增加缓冲区
	ch := make(chan any)

	// Create a worker goroutine
	go func() {
		// Send something to the channel
		// Leaks if the parent goroutine terminates early
		ch <- struct{}{}
	}()

	if err != nil {
		// The parent goroutine quits too early in case of an error
		// Sender leaks
		return
	}

	// Receive is only executed if there is no error
	<-ch
}

func Timeout(ctx context.Context) {
	// An unbuffered channel is used to coordinate
	// a worker add parent thread
	// 增加缓冲区
	ch := make(chan any)

	// Create worker goroutine
	go func() {
		// Perform some work then signal to the parent thread
		ch <- struct{}{}
	}()

	// Wait for message from worker or context
	// to be cancelled or timed out.
	select {
	case <-ch: //Receive message from workder
	case <-ctx.Done():
		// Sender leaks beacuse there is no
		// future rendezvous over the channel.
	}
}

func noCloseRange(list []any, workers int) {
	// Create channel that distributes work items.
	ch := make(chan any)

	// Create the worker goroutines.
	for i := 0; i < workers; i++ {
		go func() {
			// Each worker pulls items from the channel
			// and then processes it.
			for item := range ch {
				// Process each item
				_ = item
			}
		}()
	}

	// Queue items to the workers by using the channel.
	for _, item := range list {
		// The parent leaks by sending an item if workers == 0
		// or if all the workers panic,but the panic is recovered.
		ch <- item
	}
	// Otherwise, the channel is never closed. so worders
	// leak once there are no more items left to process.
	// close(ch)
}

func MethodContractViolation() {
	items := make([]any, 10)
	// Create a new worker
	w := NewWorker()

	// Start worker
	w.Start()

	// Operate on worker
	for _, item := range items {
		w.AddToQueue(item)
	}
	// Exits without calling `Stop`
	// w.Stop()
}

type worker struct {
	ch   chan any
	done chan any
}

type Worker interface {
	Start()
	Stop()
	AddToQueue(item any)
}

func NewWorker() Worker {
	return &worker{
		ch:   make(chan any),
		done: make(chan any),
	}
}

// Start spawns a background goroutine that extracts items pushed to the queue.
func (w *worker) Start() {
	go func() {
		for {
			select {
			case <-w.ch: // Normal workflow
			case <-w.done:
				return // Shut down
			}
		}
	}()
}

func (w *worker) Stop() {
	// Allows goroutine created by Start to terminate
	close(w.done)
}

func (w *worker) AddToQueue(item any) {
	w.ch <- item
}

type Gossip struct {
	mu     sync.Mutex
	closed bool
}

func (g *Gossip) bootstrap() {
	for {
		g.mu.Lock()
		if g.closed {
			// Missing g.mu.Unlock
			break
		}
		g.mu.Unlock()
	}
}

// Missing unlock
func Cockroach584() {
	g := &Gossip{
		closed: true,
	}
	// ...
	g.bootstrap()
	g.bootstrap() // Causes a leak
}

type node struct {
	status chan chan struct{}
	stop   chan struct{}
	done   chan struct{}
}

func (n *node) Status() struct{} {
	c := make(chan struct{})
	n.status <- c
	return <-c
}

func (n *node) run() {
	for {
		select {
		case c := <-n.status:
			c <- struct{}{}
		case <-n.stop:
			close(n.done)
			return
		}
	}
}

func (n *node) Stop() {
	select {
	case n.stop <- struct{}{}:
	case <-n.done:
		return
	}
	<-n.done
}

func Etcd6857() {
	n := &node{
		status: make(chan chan struct{}),
		stop:   make(chan struct{}),
		done:   make(chan struct{}),
	}
	go n.run()
	go n.Status()
	go n.Stop()
}

type Connection struct {
	closeChan chan bool
}

type idleAwareFramer struct {
	resetChan chan bool
	writeLock sync.Mutex
	conn      *Connection
}

func (i *idleAwareFramer) monitor() {
	var resetChan = i.resetChan
	for range i.conn.closeChan {
		i.writeLock.Lock()
		close(resetChan)
		i.resetChan = nil
		i.writeLock.Unlock()
		break
	}
}

func (i *idleAwareFramer) WriteFrame() {
	i.writeLock.Lock()
	defer i.writeLock.Unlock()
	if i.resetChan == nil {
		return
	}
	i.resetChan <- true
}

func NewIdleAwareFramer() *idleAwareFramer {
	return &idleAwareFramer{
		resetChan: make(chan bool),
		conn: &Connection{
			closeChan: make(chan bool),
		},
	}
}

func Kubernetes6632() {
	i := NewIdleAwareFramer()

	go func() {
		i.conn.closeChan <- true
	}()
	go i.monitor()
	go i.WriteFrame()
}

type Manager struct {
	plugins []int
}

func (pm *Manager) init() {
	var group sync.WaitGroup
	group.Add(len(pm.plugins))
	for _, p := range pm.plugins {
		go func(p int) {
			defer group.Done()
		}(p)
		group.Wait() // Block here
	}
}

func Moby25384() {
	pm := &Manager{
		plugins: []int{1, 2},
	}
	go pm.init()
}
