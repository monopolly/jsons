package jsons

import (
	"fmt"
	"strings"

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

//"name.last"
func String(a []byte, path ...string) string {
	return gjson.GetBytes(a, keys(path)).String()
}

func Int(a []byte, path ...string) int {
	return int(gjson.GetBytes(a, keys(path)).Int())
}
func Int64(a []byte, path ...string) (i int64) {
	return gjson.GetBytes(a, keys(path)).Int()
}

func Bool(a []byte, path ...string) bool {
	return gjson.GetBytes(a, keys(path)).Bool()
}

func Exist(a []byte, path ...string) bool {
	return gjson.GetBytes(a, keys(path)).Exists()
}

func Float(a []byte, path ...string) float64 {
	return gjson.GetBytes(a, keys(path)).Float()
}
func Bytes(a []byte, path ...string) (res []byte) {
	//res, _, _, _ = jsonparser.Get(a, key)
	return []byte(gjson.GetBytes(a, keys(path)).Raw)
}

func Interface(a []byte, path ...string) interface{} {
	return gjson.GetBytes(a, keys(path)).Value()
}

func ArrayInt(a []byte, path ...string) (r []int) {
	for _, x := range gjson.GetBytes(a, keys(path)).Array() {
		r = append(r, int(x.Int()))
	}
	return
}

func ArrayInt64(a []byte, path ...string) (r []int64) {
	for _, x := range gjson.GetBytes(a, keys(path)).Array() {
		r = append(r, x.Int())
	}
	return
}

func ArrayString(a []byte, path ...string) (r []string) {
	for _, x := range gjson.GetBytes(a, keys(path)).Array() {
		r = append(r, x.String())
	}
	return
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

func Delete(a []byte, keys ...string) (b []byte) {
	for _, x := range keys {
		b, _ = sjson.DeleteBytes(a, x)
	}
	return
}

func Set(a []byte, key string, v interface{}) (b []byte) {
	b, _ = sjson.SetBytes(a, key, v)
	return
}

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
	b, _ := ffjson.Marshal(v)
	fmt.Println(string(b))
}
