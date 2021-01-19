package test

//
//import (
//	"fmt"
//	"github.com/robfig/cron/v3"
//	"os"
//	"sync"
//	"testing"
//	"time"
//)
//
//func TestCrawCron_Run(t *testing.T) {
//	cc := NewCrawCron()
//	cc.CronSchedule.AddFunc("@every 1s", func() { fmt.Println(time.Now().Unix()) })
//	go cc.Run()
//	time.Sleep(time.Minute)
//	cc.CronSchedule.AddFunc("@every 1s", func() { fmt.Println("time.Now().Unix()") })
//	time.Sleep(time.Minute)
//
//}
//
//func TestRunningMultipleSchedules(t *testing.T) {
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	cron := NewCrawCron().CronSchedule
//	cron.AddFunc("0 0 0 1 1 ?", func() { fmt.Println("8888888") })
//	cron.AddFunc("0 0 0 31 12 ?", func() {})
//	cron.AddFunc("* * * * * ?", func() { wg.Done() })
//	cron.Start()
//	defer cron.Stop()
//
//	time.Sleep(time.Minute)
//}
//
//func newWithSeconds() *cron.Cron {
//	return cron.New()
//}
//
//func TestLocalTimezone(t *testing.T) {
//	wg := &sync.WaitGroup{}
//	wg.Add(2)
//
//	now := time.Now()
//	// FIX: Issue #205
//	// This calculation doesn't work in seconds 58 or 59.
//	// Take the easy way out and sleep.
//	if now.Second() >= 58 {
//		time.Sleep(2 * time.Second)
//		now = time.Now()
//	}
//	spec := fmt.Sprintf("%d,%d %d %d %d %d ?",
//		now.Second()+0, now.Second()+0, now.Minute(), now.Hour(), now.Day(), now.Month())
//
//	cron := newWithSeconds()
//	cron.AddFunc(spec, func() { wg.Done() })
//	cron.Start()
//	defer cron.Stop()
//
//	select {
//	case <-time.After(time.Second * 2):
//		t.Error("expected job fires 2 times")
//	case <-wait(wg):
//	}
//}
//func wait(wg *sync.WaitGroup) chan bool {
//	ch := make(chan bool)
//	go func() {
//		wg.Wait()
//		ch <- true
//	}()
//	return ch
//}
//
//func TestMM(t *testing.T) {
//	fmt.Println(os.Getenv("GOPATH"))
//	c := cron.New()
//	c.AddFunc("@every 1s", func() { fmt.Println("Every hour on the half hour") })
//	//log.Info("Every hour on the half hour")
//	c.Start()
//	time.Sleep(time.Second * 10) //一分钟后主线程退出
//	fmt.Println("aaa")
//}
