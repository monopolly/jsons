//testing
package jsons

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRaw(t *testing.T) {
	a := assert.New(t)
	_ = a

	b := []byte(`{"james":"nice"}`)

	c := Create().Add("one", 1).Add("two", "some").Bytes()

	//b = SetRaw(c, "ch", b)
	b = SetRaw(c, "ch", b)
	fmt.Println(string(b))

	ch := Bytes(b, "ch")
	fmt.Println(string(ch))

	fmt.Println(String(ch, "james"))

}

func TestArrays(t *testing.T) {
	a := assert.New(t)
	_ = a

	c := Create().Add("array", []string{"a", "b", "c"}).Bytes()

	c = ArrayStringAppend(c, "array", "a", "b", "c", "d")
	fmt.Println(string(c))

	c = ArrayStringDelete(c, "array", "a", "b")
	fmt.Println(string(c))

	c = ArrayStringAppend(c, "array", "a", "b", "c", "d")
	fmt.Println(string(c))

	//int
	c = Create().Add("array", []int{1, 2, 3}).Bytes()

	c = ArrayIntAppend(c, "array", 1, 2, 3, 4)
	fmt.Println(string(c))

	c = ArrayIntDelete(c, "array", 1, 2)
	fmt.Println(string(c))

}
