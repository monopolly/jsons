package jsons

import (
	"github.com/pquerna/ffjson/ffjson"
)

//Simple marshaling
func Marshal(v interface{}) (r []byte) {
	r, _ = ffjson.Marshal(v)
	return
}
