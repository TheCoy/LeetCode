/*
 * @lc app=leetcode.cn id=1146 lang=golang
 * @lcpr version=30122
 *
 * [1146] 快照数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type SnapshotArray struct {
	data    map[int][]int
	version int
}

func Constructor(length int) SnapshotArray {
	version := 0
	data := make(map[int][]int, 0)
	data[version] = make([]int, length)
	return SnapshotArray{
		data:    data,
		version: version,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	arr := this.data[this.version]
	arr[index] = val

	this.data[this.version] = arr
}

func (this *SnapshotArray) Snap() int {
	oldArr := this.data[this.version]
	oldVersion := this.version
	newArr := make([]int, len(oldArr))
	copy(newArr, oldArr)
	this.version++
	this.data[this.version] = newArr
	return oldVersion
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return this.data[snap_id][index]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
// @lc code=end

/*
// @lcpr case=start
// ["SnapshotArray","set","snap","set","get"][[3],[0,5],[],[0,6],[0,0]]\n
// @lcpr case=end

*/

