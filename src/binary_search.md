# 二分查找

二分查找是一种针对有序集合的查找算法，也叫折半查找算法。其查找思想类似于分治，每次都通过跟区间的中间元素对比，将带查找的区间缩小为之前的一般，直到找到要查找的元素，或者区间被缩小为0。

## O(logn)的查找速度

假设数据规模为n，查找区间的变化满足等比数列`n，n/2，n/4，n/8，...，n/(2^k)，...`。其中`n/(2^k)=1`时，k的值就是总共缩小的次数。而每一次缩小操作只涉及两个元素的比较，所以经过了k次区间缩小操作，时间复杂度就是`O(k)`。通过`n/(2^k)=1`，我们可以求得`k=logn/log2`，所以时间复杂度就是`O(logn)`。

时间复杂度为`O(logn)`的算法有时甚至比时间复杂度为`O(1)`的算法效率更高。在用`O`标记法表示算法的时间复杂度的时候会省略掉常数、系数和低阶。对于常数级时间复杂度的算法来说，`O(1)`有可能表示的是一个非常大的常量值，比如`O(1000)`、`O(10000)`。

## 二分查找的递归与非递归实现

最简单的情况就是有序数组中不存在重复元素，用二分查找找到值等于给定值的数据。

递归实现如下

```Golang
func binarySearch(nums []int, value int) int {
	return recursion(nums, 0, len(nums)-1, value)
}

func recursion(nums []int, left, right int, value int) int {
	if left > right {
		return -1
	}

	mid := left + ((right - left) >> 1)
	if nums[mid] == value {
		return mid
	} else if nums[mid] < value {
		return recursion(nums, mid+1, right, value)
	} else {
		return recursion(nums, left, mid-1, value)
	}
}
```

非递归实现如下

```Golang
func binarySearch(nums []int, value int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == value {
			return mid
		} else if nums[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

### 非递归实现中容易出错的地方

1. 循环终止的条件

    终止条件为`low<=high`而不是`low<high`。因为`low<=high`可以正确的处理区间只有一个元素的场景，而`low<high`会漏掉这个场景下的数据比较。

2. mid的取值

    `mid = (low + high) / 2`在`low`或者`high`很大的时候可能会有数据溢出的风险，更好的写法为`mid = low + (high - low)/2`。为了提升性能还可以改写为`mid = low + ((high - low)>>1)`，计算机处理位运算的速度总比除法来得快。

3. low和high的更新

    `low=mid+1`，`high=mid-1`。这里如果`low=mid`或者`high=mid`就可能会发生死循环。比如当`high=3`，`low=3`，`nums[3]!=value`时代码就无法退出了。

## 二分查找应用场景的局限性

1. 二分查找依赖顺序表结果，也就是依赖数组

    由于二分查找的实现依赖于随机访问，数组的随机访问时间复杂度为`O(1)`，而链表的随机访问时间复杂度为`O(n)`。

2. 二分查找针对的是有序数据

    如果数据无序则需要先做排序，而这部分的时间复杂度至少为`O(nlogn)`。所以，如果数组的插入、删除操作很频繁，那么就需要频繁地在二分查找之前排序，提高了查找成本。

3. 数据量太小或太小不适合二分查找

    如果数据量很小，二分查找和直接遍历并不会有太多差别。有一个例外：元素间比较很复杂，这时候可以采用二分查找来减少比较次数。

    由于二分查找依赖于数组，如果数据量很大就会有巨大的内存开销。

## 必知必会

实现一个有序数组的二分查找算法

查找有序数组中第一个值等于给定值的元素

查找有序数组中最后一个值等于给定值的元素

查找有序数组中第一个值大于等于给定值的元素

查找有序数组中最后一个值小于等于给定值的元素
