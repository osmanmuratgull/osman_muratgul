package main

import "fmt"

func totalone(a int, b int) int {

	return a + b
}

func totaltwo(a, b, c int) int {
	return a + b + c
}

func totalthree(a, b, c, d int) int {

	return a + b + c/d
}

func main() {

	res := totalone(1, 2)
	fmt.Println("1+2 =", res)

	res = totaltwo(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res = totalthree(1, 2, 3, 2)
	fmt.Println("1+2+3/2 = ", res)

}
