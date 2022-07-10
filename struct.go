package jsons

import "fmt"

func Create() (r *Json) {
	return &Json{}
}

func Creates(k string, v interface{}) (r *Json) {
	r = &Json{}
	r.Add(k, v)
	return
}

type Json struct {
	count int
	b     []byte
}

func (a *Json) Add(k string, v interface{}) *Json {
	a.b = Set(a.b, k, v)
	return a
}

func (a *Json) AddRaw(k string, v []byte) *Json {
	a.b = SetRaw(a.b, k, v)
	return a
}

func (a *Json) Array(v interface{}) *Json {
	a.b = Set(a.b, fmt.Sprint(a.count), v)
	a.count++
	return a
}

func (a *Json) Bytes() []byte {
	return a.b
}

func (a *Json) String() string {
	return string(a.b)
}
