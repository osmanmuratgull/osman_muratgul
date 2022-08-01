package main

import "fmt"

func totalOne(a int, b int) int {

	return a + b
}

func totalTwo(a, b, c int) int {
	return a + b + c
}

func totalThree(a, b, c, d int) int {

	//TODO
	//toplama işlemi olarak başlık var fakat içeride bölme yapılıyor 
	//fonksiyon ismi kod ile uyumlu olmalıdır her zaman
	return a + b + c/d
}

func main() {

	res := totalOne(1, 2)
	fmt.Println("1+2 =", res)

	res = totalTwo(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res = totalThree(1, 2, 3, 2)
	fmt.Println("1+2+3/2 = ", res)

	//Algorithm that calculates simple math
	//Osman Muratgul
}
