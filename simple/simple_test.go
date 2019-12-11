package simple

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	var test3 string = "t"
	var test4 []byte = []byte("hello")
	var test5 float64 = 1.12234234
	s := New(map[string]interface{}{
		"test":  1.1,
		"test2": 16,
		"test3": test3,
		"test4": test4,
		"test5": test5,
	})
	assert.NotNil(t, s)
	result := s.Get("test3")
	fmt.Println(reflect.TypeOf(result))

	i, err := strconv.Atoi(fmt.Sprintf("%v", result))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i)

	result4 := s.GetString("test4")
	fmt.Println(result4)

	result5 := s.GetString("test5")
	fmt.Println(result5)
}
