package seq

import (
	"github.com/gesius/gokit/atomic"
	"testing"
)

func TestNewSEQNO(t *testing.T) {
	var v *SEQNO

	v = NewSEQNO("gokit.atomic.seqno.0")

	if v == nil {
		t.Error("Expected: NOT NIL, got: ", v)
	}
	if v.Get() != 0 {
		t.Error("Expected: 0, got: ", v)
	}

}

func Testincr(t *testing.T) {
	var v *SEQNO

	v = NewSEQNO("gokit.atomic.seqno.0")

	if v == nil {
		t.Error("Expected: NOT NIL, got: ", v)
	}

	if v.Get() != 0 {
		t.Error("Expected: 0, got: ", v)
	}

	v.Incr()

	if v.Get() != 1 {
		t.Error("Expected: 1, got: ", v)
	}

	v.Incr()
	v.Incr()
	v.Decr()

	if v.Get() != 2 {
		t.Error("Expected: 2, got: ", v)
	}
}
