package seq

import (
	"sync/atomic"
)

//see: https://github.com/rcrowley/go-metrics/blob/master/counter.go
type SEQNO struct {
	n int64
}

func (s *SEQNO) get() int64 {
	return atomic.LoadInt64(&s.n)
}

func (s *SEQNO) incr() {
	atomic.AddInt64(&s.n, 1)
}

func (s *SEQNO) decr() {
	atomic.AddInt64(&s.n, -1)
}

func NewSEQNO(name string) *SEQNO {

	s := &SEQNO{n: 0}
	s.incr()

	return s
}
