package cqueue

type CQueue struct {
	appendList []int
	headList   []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.appendList = append(this.appendList, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.headList) == 0 {
		if len(this.appendList) == 0 {
			return -1
		}
		for len(this.appendList) > 0 {
			this.headList = append(this.headList, this.appendList[len(this.appendList)-1])
			this.appendList = this.appendList[:len(this.appendList)-1]
		}
	}
	v := this.headList[len(this.headList)-1]
	this.headList = this.headList[:len(this.headList)-1]
	return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
