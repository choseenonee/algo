package leetcode

type RandomizedSet struct {
	data map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{data: map[int]struct{}{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.data[val]; ok {
		return false
	}

	rs.data[val] = struct{}{}

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.data[val]; ok {
		delete(rs.data, val)
		return true
	}

	return false
}

func (rs *RandomizedSet) GetRandom() int {
	for key, _ := range rs.data {
		return key
	}

	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
