package rand

import (
	"crypto/rand"
	"encoding/binary"
	"runtime"
	"sync"
)

type RND struct{ n int32 }

func NewRND() *RND {
	return &RND{n: 0}
}

func (r *RND) Rand() *RND {
	binary.Read(rand.Reader, binary.LittleEndian, &r.n)
	return r
}

func (r *RND) ToString() string {
	if r == nil {
		return nil
	}
	return strconv.Itoa(r.n)
}

func (r *RND) RandomString(p string s string) {
	n := NewRND().Rand().ToString()
	buffer.WriteString(p)
	buffer.WriteString(r.Rand().ToString())
	buffer.WriteString(":")
	buffer.WriteString(s)

	return buffer.String()
}