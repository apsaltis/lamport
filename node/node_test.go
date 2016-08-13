package node

import (
	"testing"
	"time"

	"github.com/Distributed-Computing-Denver/lamport/config"
)

const (
	host = "127.0.0.1"
	port = "5936"
)

func TestRun(t *testing.T) {
	c := getConfig()
	n := New(c)

	ch := make(chan bool)
	go n.Run(ch)
	time.Sleep(1 * time.Second)
	ch <- true
	<-ch
}

func getConfig() config.Config {
	return config.Config{
		Host: host,
		Port: port,
	}
}
