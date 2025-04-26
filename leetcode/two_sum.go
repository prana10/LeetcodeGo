package leetcode

// penerapan hash map untuk mencapai kompleksitas O(n)
func TwoSum(nums []int, target int) []int {
	// hash map untuk menyimpan nilai dan index
	hashMap := make(map[int]int)

	// iterasi melalui array
	for i, num := range nums {
		// cari nilai yang jika ditambahkan dengan num, hasilnya sama dengan target
		complement := target - num

		// cek apakah nilai complement ada di hash map
		if index, ok := hashMap[complement]; ok {
			// jika ada, return index dan i
			return []int{index, i}
		}

		// jika tidak ada, tambahkan nilai num dan index ke hash map
		hashMap[num] = i
	}

	// jika tidak ada pasangan yang memenuhi syarat, return []int{}
	return []int{}
}
