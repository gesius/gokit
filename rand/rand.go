package rand

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

type RND struct {
	n int32
}

func NewRND() *RND {
	return &RND{n: 0}
}

func (r *RND) Rand() *RND {
	binary.Read(rand.Reader, binary.LittleEndian, &r.n)
	return r
}

func (r *RND) ToString() string {
	return strconv.Itoa(r.n)
}

func (r *RND) RandomString(p string, s string) {
	var b bytes.Buffer
	n := NewRND().Rand().ToString()
	buffer.WriteString(p)
	buffer.WriteString(r.Rand().ToString())
	buffer.WriteString(":")
	buffer.WriteString(s)

	return buffer.String()
}
