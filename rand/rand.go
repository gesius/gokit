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
	return strconv.Itoa(int(r.n))
}

func (r *RND) RandomString(p string, s string) string {
	var b bytes.Buffer

	b.WriteString(p)
	b.WriteString(r.Rand().ToString())
	b.WriteString(":")
	b.WriteString(s)

	return b.String()
}
