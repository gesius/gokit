package seq

import (
	"sync/atomic"
)

type SEQNO_UINT64 struct {
	N uint64
}

func (s *SEQNO) incr() {
	atomic.AddUint64(&s.value, 1)
}

func (s *SEQNO) decr() {
	atomic.AddUint64(&s.value, -1)
}

func NewSEQNO(name string) *SEQNO {

	s := &SEQNO{N: 0}

	return s
}
