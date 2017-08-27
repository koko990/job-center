package etcdsync

import (
	"testing"

	"fmt"
	"time"
	"sync"

	"github.com/gocomb/job-center/server/util"
)

func TestMutex_Lock(t *testing.T) {

	m := LockFactory("/etcdsync", 123, []string{"http://10.110.18.26:2379"})
	if m == nil {
		util.Logger.SetInfo("etcdsync.NewMutex failed")
	}
	err := m.Lock()
	if err != nil {
		util.Logger.SetInfo("etcdsync.Lock failed")
	} else {
		util.Logger.SetInfo("etcdsync.Lock OK")
	}

	util.Logger.SetInfo("Get the lock. Do something here.")

	err = m.Unlock()
	if err != nil {
		util.Logger.SetInfo("etcdsync.Unlock failed")
	} else {
		util.Logger.SetInfo("etcdsync.Unlock OK")
	}
}
func method1(m *sync.Mutex,done chan struct{}) {
	m.Lock()
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("this is method 1")
	m.Unlock()
	done<- struct{}{}
}
func method2(m *sync.Mutex,done chan struct{}) {
	m.Lock()
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("this is method 2")
	m.Unlock()
	done<- struct{}{}
}
func TestHelloWorld(t *testing.T) {
	m := new(sync.Mutex)
	var method1Done chan struct{}
	var method2Done chan struct{}
	method1Done=make(chan struct{})
	method2Done=make(chan struct{})
	go method1(m,method1Done)
	go method2(m,method2Done)
	<-method1Done
	<-method2Done
}
