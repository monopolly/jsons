package jsons

import (
	"bytes"

	"github.com/pquerna/ffjson/ffjson"
)

//Simple marshaling
func Marshal(v interface{}) (r []byte) {
	r, _ = ffjson.Marshal(v)
	return
}

//Simple marshaling
func Unmarshal(data []byte, v interface{}) (err error) {
	return ffjson.Unmarshal(data, v)
}

//Simple json, json, json = []byte()
func MarshalCreateJsonArray(v ...[]byte) (r []byte) {
	var b bytes.Buffer
	b.WriteString("[")
	b.Write(bytes.Join(v, []byte(",")))
	b.WriteString("]")
	return b.Bytes()
}
