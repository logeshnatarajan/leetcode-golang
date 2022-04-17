type MyQueue struct {
    v []int
}


func Constructor() MyQueue {
    return MyQueue{[]int{}}
}


func (this *MyQueue) Push(x int)  {
    this.v=append(this.v,x)
}


func (this *MyQueue) Pop() int {
    temp:=this.v[0]
    this.v=this.v[1:]
    return temp
}


func (this *MyQueue) Peek() int {
    return this.v[0]
}


func (this *MyQueue) Empty() bool {
    if len(this.v)==0{
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
