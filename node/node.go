package node

import (
	"log"

	"github.com/Distributed-Computing-Denver/lamport/config"
)

// Runner runs until signalled to stop by sigCh
type Runner interface {
	Run(sigCh chan bool)
}

type node struct {
	conf config.Config
}

// New creates a Runner that can be used to
// run a Lamport node
func New(conf config.Config) Runner {
	return node{conf: conf}
}

// Start starts a new lamport node using the supplied
// Runner
func (n node) Run(sigCh chan bool) {
	sig := <-sigCh
	if sig {
		log.Println("Signal received, terminating Lamport")
		sigCh <- true
		return
	}
}
