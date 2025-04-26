package leetcode

// Kompleksitas
// Waktu: O(n) karena kita hanya melakukan satu kali scan array
// Ruang: O(1) karena kita hanya menggunakan beberapa variabel pointer
func CountSubarrays(nums []int, minK int, maxK int) int64 {
	// result, lastOutOfBound, lastMin, lastMax
	var result int64 = 0 //result untuk menampung hasil
	lastOutOfBound := -1 //lastOutOfBound untuk menampung index terakhir yang keluar dari batas minK dan maxK
	lastMin := -1        //lastMin untuk menampung index terakhir dari nilai minK
	lastMax := -1        //lastMax untuk menampung index terakhir dari nilai maxK

	// sliding window
	// A. kita scan dari kiri ke kanan, untuk mencari setiap posisi yang memenuhi syarat
	// 1. jika posisi terakhir kali sama dengan nilai minK
	// 2. jika posisi terakhir kali sama dengan nilai maxK
	// 3. jika posisi terakhir kali keluar dari batas minK dan maxK
	// B. setiap posisi `i`, dihitung berapa banyak sub array yang valid di posisi terakhir tersebut.

	// fungsi `i` untuk index dari array
	// fungsi `num` untuk nilai dari array
	for i, num := range nums {
		// kalo kita nemuin nilai yang keluar dari batas minK dan maxK, maka kita update lastOutOfBound
		if num < minK || num > maxK {
			lastOutOfBound = i
		}
		// kalo kita nemuin nilai yang sama dengan minK, maka kita update lastMin
		if num == minK {
			lastMin = i
		}
		// kalo kita nemuin nilai yang sama dengan maxK, maka kita update lastMax
		if num == maxK {
			lastMax = i
		}

		// proses rumus
		// kita hitung subb array yang valid di posisi terakhir tersebut.
		// rumus nya = min(lastMin, lastMax) - lastOutOfBound
		// jika nilai array adalah 1,1,1,1
		// maka proses hitung nya seperti ini :
		// iterasi 0 :
		// lastMin = 0, lastMax = 0, lastOutOfBound = -1
		// lastMin bisa 0 karena nilai minK = 1, dan nilai array di index 0 adalah 1
		// lastMax bisa 0 karena nilai maxK = 1, dan nilai array di index 0 adalah 1
		// lastOutOfBound bisa -1 karena nilai array di index 0 adalah 1, dan tidak keluar dari batas minK dan maxK
		// result nya = min(0,0) - (-1) = 1

		// iterasi 1 :
		// lastMin = 1, lastMax = 1, lastOutOfBound = -1
		// lastMin bisa 1 karena nilai minK = 1, dan nilai array di index 1 adalah 1
		// lastMax bisa 1 karena nilai maxK = 1, dan nilai array di index 1 adalah 1
		// lastOutOfBound bisa -1 karena nilai array di index 1 adalah 1, dan tidak keluar dari batas minK dan maxK
		// result nya = min(1,1) - (-1) = 2

		// iterasi 2 :
		// lastMin = 2, lastMax = 2, lastOutOfBound = -1
		// lastMin bisa 2 karena nilai minK = 1, dan nilai array di index 2 adalah 1
		// lastMax bisa 2 karena nilai maxK = 1, dan nilai array di index 2 adalah 1
		// lastOutOfBound bisa -1 karena nilai array di index 2 adalah 1, dan tidak keluar dari batas minK dan maxK
		// result nya = min(2,2) - (-1) = 3

		// iterasi 3 :
		// lastMin = 3, lastMax = 3, lastOutOfBound = -1
		// lastMin bisa 3 karena nilai minK = 1, dan nilai array di index 3 adalah 1
		// lastMax bisa 3 karena nilai maxK = 1, dan nilai array di index 3 adalah 1
		// lastOutOfBound bisa -1 karena nilai array di index 3 adalah 1, dan tidak keluar dari batas minK dan maxK
		// result nya = min(3,3) - (-1) = 4

		// jadi result nya adalah 1 + 2 + 3 + 4 = 10
		if lastMin > lastOutOfBound && lastMax > lastOutOfBound {
			result += int64(min(lastMin, lastMax) - lastOutOfBound)
		}
	}
	return result
}
