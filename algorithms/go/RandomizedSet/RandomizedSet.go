package randomizedset

import "math/rand"

type RandomizedSet struct {
	s []int
	m map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.s)
	this.s = append(this.s, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	pos, ok := this.m[val]
	if !ok {
		return false
	}
	last := len(this.s) - 1
	this.s[pos] = this.s[last]
	this.m[this.s[pos]] = pos
	this.s = this.s[:last]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.s) <= 0 {
		return 0
	}
	return this.s[rand.Intn(len(this.s))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
