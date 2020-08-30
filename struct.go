package jsons

import "fmt"

func Create() (r *json) {
	return &json{}
}

type json struct {
	count int
	b     []byte
}

func (a *json) Add(k string, v interface{}) *json {
	a.b = Set(a.b, k, v)
	return a
}

func (a *json) Array(v interface{}) *json {
	a.b = Set(a.b, fmt.Sprint(a.count), v)
	a.count++
	return a
}

func (a *json) Bytes() []byte {
	return a.b
}
