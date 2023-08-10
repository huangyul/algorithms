package main

// 随机访问
func f1() {
	arr := [5]int{1, 2, 3, 4, 5}
	print(arr[3]) // 4
}

// 访问越界
func f2() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i <= 4; i++ {
		arr[i] = 0
		print("hello world")
	}
}

func main() {
	// f1()
	f2()
}
