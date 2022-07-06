type MyHashSet struct {
    m map[int]bool
}


func Constructor() MyHashSet {
	return MyHashSet{
		m: make(map[int]bool),
	}
}


func (this *MyHashSet) Add(key int)  {
    this.m[key]=true
}


func (this *MyHashSet) Remove(key int)  {
    delete(this.m,key)
}


func (this *MyHashSet) Contains(key int) bool {
    return this.m[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
