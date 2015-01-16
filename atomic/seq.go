package atomic

import (
	"sync/atomic"
)

//see: https://github.com/rcrowley/go-metrics/blob/master/counter.go
type SEQNO struct {
	n int64
}

func (s *SEQNO) Get() int64 {
	return atomic.LoadInt64(&s.n)
}

func (s *SEQNO) Incr() {
	atomic.AddInt64(&s.n, 1)
}

func (s *SEQNO) Decr() {
	atomic.AddInt64(&s.n, -1)
}

func NewSEQNO(name string) *SEQNO {

	s := &SEQNO{n: 0}

	return s
}
