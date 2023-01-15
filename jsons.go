package jsons

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Result = gjson.Result

/* {
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
"name.last"          >> "Anderson"
"age"                >> 37
"children"           >> ["Sara","Alex","Jack"]
"children.#"         >> 3
"children.1"         >> "Alex"
"child*.2"           >> "Jack"
"c?ildren.0"         >> "Sara"
"fav\.movie"         >> "Deer Hunter"
"friends.#.first"    >> ["Dale","Roger","Jane"]
"friends.1.last"     >> "Craig" */

func keys(k []string) string {
	return strings.Join(k, ".")
}

// "name.last"
func String(a []byte, path ...string) string {
	return gjson.GetBytes(a, keys(path)).String()
}

func Int(a []byte, path ...string) int {
	return int(gjson.GetBytes(a, keys(path)).Int())
}

func Byte(a []byte, path ...string) byte {
	b := gjson.GetBytes(a, keys(path)).Int()
	if b > 0 && b < 256 {
		return byte(b)
	}
	return 0
}

func Uint8(a []byte, path ...string) uint8 {
	b := gjson.GetBytes(a, keys(path)).Int()
	if b > 0 && b < 256 {
		return uint8(b)
	}
	return 0
}

func Int64(a []byte, path ...string) (i int64) {
	return gjson.GetBytes(a, keys(path)).Int()
}

func Uint64(a []byte, path ...string) (i uint64) {
	return gjson.GetBytes(a, keys(path)).Uint()
}

func Uint(a []byte, path ...string) (i uint) {
	return uint(gjson.GetBytes(a, keys(path)).Uint())
}

func Uint32(a []byte, path ...string) (i uint32) {
	return uint32(gjson.GetBytes(a, keys(path)).Uint())
}

func TimeDuration(a []byte, path ...string) (i time.Duration) {
	return time.Duration(gjson.GetBytes(a, keys(path)).Int())
}

func Bool(a []byte, path ...string) bool {
	return gjson.GetBytes(a, keys(path)).Bool()
}

func Exist(a []byte, path ...string) bool {
	return gjson.GetBytes(a, keys(path)).Exists()
}

func Float64(a []byte, path ...string) float64 {
	return gjson.GetBytes(a, keys(path)).Float()
}

func Float32(a []byte, path ...string) float32 {
	return float32(gjson.GetBytes(a, keys(path)).Float())
}

// Bytes get raw json
func Bytes(a []byte, path ...string) (res []byte) {
	return []byte(gjson.GetBytes(a, keys(path)).Raw)
}

func Interface(a []byte, path ...string) interface{} {
	x := gjson.GetBytes(a, keys(path)).Value()
	switch x.(type) {
	case int:
		return x.(int)
	case int64:
		return x.(int64)
	case int8:
		return x.(int8)
	case int16:
		return x.(int16)
	case int32:
		return x.(int32)
	case uint:
		return x.(uint)
	case uint8:
		return x.(uint8)
	case uint16:
		return x.(uint16)
	case uint32:
		return x.(uint32)
	case uint64:
		return x.(uint64)
	case string:
		return x.(string)
	case []byte:
		return x.([]byte)
	case float64:
		return x.(float64)
	case float32:
		return x.(float32)
	case bool:
		return x.(bool)
	case []string:
		return x.([]string)
	case []int:
		return x.([]int)
	case []int64:
		return x.([]int64)
	case []bool:
		return x.([]bool)
	case []interface{}:
		return x.([]interface{})
	case map[string]string:
		return x.(map[string]string)
	case map[string]int:
		return x.(map[string]int)
	case map[string]bool:
		return x.(map[string]bool)
	case map[string]int64:
		return x.(map[string]int64)
	case map[string]interface{}:
		return x.(map[string]interface{})
	case map[string][]byte:
		return x.(map[string][]byte)
	case map[int]string:
		return x.(map[int]string)
	case map[int]int:
		return x.(map[int]int)
	case map[int]bool:
		return x.(map[int]bool)
	case map[int]int64:
		return x.(map[int]int64)
	default:
		return x
	}
}

func ArrayInt(a []byte, path ...string) (r []int) {
	for _, x := range gjson.GetBytes(a, keys(path)).Array() {
		r = append(r, int(x.Int()))
	}
	return
}

func ArrayIntAppend(a []byte, key string, items ...int) (res []byte) {
	var list []int
	un := make(map[int]bool)
	for _, x := range ArrayInt(a, key) {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}
	for _, x := range items {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}

	return Set(a, key, list)
}

func ArrayIntDelete(a []byte, key string, items ...int) (res []byte) {
	var list []int
	un := make(map[int]bool)
	for _, x := range items {
		un[x] = true
	}

	for _, x := range ArrayInt(a, key) {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}

	return Set(a, key, list)
}

func ArrayInt64(a []byte, path ...string) (r []int64) {
	for _, x := range gjson.GetBytes(a, keys(path)).Array() {
		r = append(r, x.Int())
	}
	return
}

func ArrayInt64Append(a []byte, key string, items ...int64) (res []byte) {
	var list []int64
	un := make(map[int64]bool)
	for _, x := range ArrayInt64(a, key) {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}
	for _, x := range items {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}

	return Set(a, key, list)
}

func ArrayInt64Delete(a []byte, key string, items ...int64) (res []byte) {
	var list []int64
	un := make(map[int64]bool)
	for _, x := range items {
		un[x] = true
	}

	for _, x := range ArrayInt64(a, key) {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}

	return Set(a, key, list)
}

func ArrayString(a []byte, path ...string) (r []string) {
	for _, x := range gjson.GetBytes(a, keys(path)).Array() {
		r = append(r, x.String())
	}
	return
}

// ArrayStringAppend append unique items (no duplicates)
func ArrayStringAppend(a []byte, key string, items ...string) (res []byte) {
	var list []string
	un := make(map[string]bool)
	for _, x := range ArrayString(a, key) {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}
	for _, x := range items {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}

	return Set(a, key, list)
}

func ArrayStringDelete(a []byte, key string, items ...string) (res []byte) {
	var list []string
	un := make(map[string]bool)
	for _, x := range items {
		un[x] = true
	}

	for _, x := range ArrayString(a, key) {
		if un[x] {
			continue
		}
		un[x] = true
		list = append(list, x)
	}

	return Set(a, key, list)
}

func Array(a []byte, key ...string) (r []Result) {
	if len(key) == 0 {
		return gjson.GetManyBytes(a)
	}

	return gjson.GetBytes(a, key[0]).Array()
}

func Get(a []byte, path ...string) (r Result) {
	return gjson.GetBytes(a, keys(path))
}

func ArrayLen(a []byte, path ...string) (count int) {
	return len(gjson.GetBytes(a, keys(path)).Array())
}

func Count(a []byte) int {
	return len(gjson.Parse(string(a)).Map())
}

func Keys(a []byte) (res []string) {
	for x, _ := range gjson.Parse(string(a)).Map() {
		res = append(res, x)
	}
	return
}

func Iterator(a []byte, f func(k string, v gjson.Result)) {
	for kk, vv := range gjson.Parse(string(a)).Map() {
		f(kk, vv)
	}
	return
}

func MapString(a []byte, path ...string) (r map[string]string) {
	if a == nil {
		return
	}
	r = make(map[string]string)
	for k, v := range gjson.GetBytes(a, keys(path)).Map() {
		r[k] = v.String()
	}
	return
}

func MapInterface(a []byte, path ...string) (r map[string]interface{}) {
	if a == nil {
		return
	}
	r = make(map[string]interface{})
	for k, v := range gjson.GetBytes(a, keys(path)).Map() {
		r[k] = v.Value()
	}
	return
}

func MapInt(a []byte, path ...string) (r map[string]int) {
	if a == nil {
		return
	}
	r = make(map[string]int)
	for k, v := range gjson.GetBytes(a, keys(path)).Map() {
		r[k] = int(v.Int())
	}
	return
}

func MapBool(a []byte, path ...string) (r map[string]bool) {
	if a == nil {
		return
	}
	r = make(map[string]bool)
	for k, v := range gjson.GetBytes(a, keys(path)).Map() {
		r[k] = v.Bool()
	}
	return
}

func MapIntString(a []byte, path ...string) (r map[int]string) {
	if a == nil {
		return
	}
	r = make(map[int]string)
	for k, v := range gjson.GetBytes(a, keys(path)).Map() {
		i, _ := strconv.Atoi(k)
		r[i] = v.String()
	}
	return
}

func MapIntInt(a []byte, path ...string) (r map[int]int) {
	if a == nil {
		return
	}
	r = make(map[int]int)
	for k, v := range gjson.GetBytes(a, keys(path)).Map() {
		i, _ := strconv.Atoi(k)
		r[i] = int(v.Int())
	}
	return
}

func Delete(a []byte, keys ...string) (b []byte) {
	for _, x := range keys {
		b, _ = sjson.DeleteBytes(a, x)
	}
	return
}

// Set set non json into key
func Set(a []byte, key string, v interface{}) (b []byte) {
	b, _ = sjson.SetBytes(a, key, v)
	return
}

// SetRaw set json into key
func SetRaw(a []byte, key string, v []byte) (b []byte) {
	b, _ = sjson.SetRawBytes(a, key, v)
	return
}

func Sets(a *[]byte, key string, v interface{}) {
	if a == nil {
		return
	}
	(*a), _ = sjson.SetBytes((*a), key, v)
}

func Print(v interface{}) {
	b, _ := ffjson.Marshal(v) //ffjson.Marshal(v)
	fmt.Println(string(b))
}

func Join(v [][]byte) (res []byte) {
	var w bytes.Buffer
	w.WriteString("[")
	for _, x := range v {
		w.Write(x)
		w.WriteString(",")
	}
	res = w.Bytes()
	res[len(res)-1] = ']'
	return
}
