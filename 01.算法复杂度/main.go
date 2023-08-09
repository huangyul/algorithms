package main

func cal1() {
	i := 0
	sum := 0

	for ; i < 10; i++ {
		sum += i
	}

	print(sum)
}

func cal2(n int) {
	//第一段代码
	for i := 1; i < 100; i++ {
		print(i)
	}

	// 第二段代码
	for i := 0; i < n; i++ {
		print(i)
	}

	// 第三段代码
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			print(i * i)
		}
	}
}

// 最好、最坏时间复杂度
// n表示arr的长度
func f1(arr []int, n int, x int) int {
	i := 0
	pos := -1
	for ; i < n; i++ {
		if arr[i] == x {
			pos = i
		}
	}
	return pos
}

func main() {
	print(f1([]int{1, 2, 3, 4}, 4, 2))
}
