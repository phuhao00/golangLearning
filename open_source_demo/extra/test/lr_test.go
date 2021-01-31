package test

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func TestLimiterRate(t *testing.T) {

	var lr *rate.Limiter = rate.NewLimiter(rate.Limit(0.2), 1)
	go L1(lr)
	go L2(lr)
	for {

	}
}

func L1(lr *rate.Limiter) {
	for {
		if lr.Allow() {
			fmt.Println("1---", time.Now().Unix())
		}
	}
}

func L2(lr *rate.Limiter) {
	for {
		if lr.Allow() {
			fmt.Println("2---", time.Now().Unix())
		}
	}
}
