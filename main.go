package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type myReadWriter struct {
	sup int
	inf int
	s   []byte
	i   int
}

func (r *myReadWriter) Read(b []byte) (n int, err error) {
	var aux string
	r.sup = len(b) + r.inf
	if r.sup > len(r.s) && r.inf < len(r.s) {
		r.sup = len(r.s)
	}
	if r.sup > len(r.s) {
		return 0, io.EOF
	}
	aux = string(r.s[r.inf:r.sup])
	r.inf = r.sup
	n = len(aux)
	for i, v := range aux {
		b[i] = byte(v)
	}
	return
}

func (r *myReadWriter) Write(b []byte) (n int, err error) {
	r.s = append(r.s, b...)
	n = len(b)
	r.i = n
	return
}

func main() {
	var rw = &myReadWriter{}
	var str = "Insert text here!!!"
	n, err := io.WriteString(rw, str)
	fmt.Println(n, err)

	fmt.Println(rw)
	buf := make([]byte, 5)
	rw.Read(buf)
	fmt.Println(string(buf))

	fmt.Println(rw)
	all, err := ioutil.ReadAll(rw)
	fmt.Println(string(all), err)
	fmt.Println(rw)

}
