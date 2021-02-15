package main

func main()  {
}

//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]
func merge(nums1 []int, m int, nums2 []int, n int)  {
	var l, r, s int
	l = m - 1
	r = n - 1
	s = m + n - 1
	for l >= 0 && r >= 0 {
		if nums1[l] < nums2[r] {
			nums1[s] = nums2[r]
			r --
		} else {
			nums1[s] = nums1[l]
			l --
		}
		s --
	}
	if r >= 0 {
		for i:= 0; i <= r; i ++ {
			nums1[i] = nums2[i]
		}
	}
	return
}

