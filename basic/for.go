package main

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)

	sum_c := 1
	for sum_c < 1000 {
		sum_c += sum_c
	}
	println(sum_c)
}
