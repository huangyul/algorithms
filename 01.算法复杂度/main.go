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

func main() {}
