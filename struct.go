package jsons

import (
	"encoding/json"
)

func Create() (r *Json) {
	r = &Json{}
	r.b = make(map[string]interface{})
	return r
}

func Creates(k string, v interface{}) (r *Json) {
	r = &Json{}
	r.b = make(map[string]interface{})
	r.Add(k, v)
	return
}

type Json struct {
	b     map[string]interface{}
	array []interface{}
}

func (a *Json) Add(k string, v interface{}) *Json {
	a.b[k] = v
	return a
}

func (a *Json) AddRaw(k string, v []byte) *Json {
	m := make(map[string]interface{})
	json.Unmarshal(v, &m)
	a.b[k] = m
	return a
}

func (a *Json) Array(v interface{}) *Json {
	a.array = append(a.array, v)
	return a
}

func (a *Json) Bytes() (res []byte) {
	res, _ = json.Marshal(a.b)
	return
}

func (a *Json) String() string {
	return string(a.Bytes())
}
