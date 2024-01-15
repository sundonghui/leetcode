package randomizedset

import (
	"testing"
)

func Test_randomizedSet(t *testing.T) {
	obj := Constructor()
	t.Log(obj.Insert(2))
	t.Log(obj.Remove(2))
	t.Log(obj.GetRandom())
}
