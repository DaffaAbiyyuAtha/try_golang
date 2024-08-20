package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func calc(num1 int, num2 int) int {
// 	// var count int = num1 + num2
// 	count := num1 + num2
// 	return count
// }
// func keliling(r float32, pi float32) float32 {
// 	// var count int = num1 + num2
// 	kels := r *2 *pi
// 	return kels
// }

// func luas(r float32, pi float32) float32 {
// 	// var count int = num1 + num2
// 	luas := r * r *pi
// 	return luas
// }

// var num1 = 0
// var num2 = 0

func out(resultnum []string) {
	// for _, v := range resultnum {
	// strconv.Atoi(v[0])
	// num1 := v[0] + v[1]
	// num2 := v[2] + v[3]
	// resultnum = num1
	// resultnum = num2
	// 	fmt.Println(num1)
	// }
	if len(resultnum)%2 != 0 {
		panic("Error : slice should be even")
	}

	left := 0
	right := 0

	for _, v := range resultnum[:len(resultnum)/2] {
		isNum, _ := strconv.Atoi(v)
		left += isNum
	}
	for _, v := range resultnum[len(resultnum)/2:] {
		isNum, _ := strconv.Atoi(v)
		right += isNum
	}

	// fmt.Println(left)
	// fmt.Println(right)

	fmt.Println(left)
	fmt.Println(right)
	// length := len(resultnum)

	// if length == 4 {
	// 	left1, _ := strconv.Atoi(resultnum[0])
	// 	left2, _ := strconv.Atoi(resultnum[1])
	// 	right1, _ := strconv.Atoi(resultnum[3])
	// 	right2, _ := strconv.Atoi(resultnum[4])

	// 	fmt.Println(left1 + left2)
	// 	fmt.Println(right1 + right2)
	// } else {
	// 	left1, _ := strconv.Atoi(resultnum[0])
	// 	left2, _ := strconv.Atoi(resultnum[1])
	// 	left3, _ := strconv.Atoi(resultnum[2])
	// 	right1, _ := strconv.Atoi(resultnum[3])
	// 	right2, _ := strconv.Atoi(resultnum[4])
	// 	right3, _ := strconv.Atoi(resultnum[5])

	// 	fmt.Println(left1 + left2 + left3)
	// 	fmt.Println(right1 + right2 + right3)
	// }

	// left1, _ := strconv.Atoi(resultnum[0])
	// left2, _ := strconv.Atoi(resultnum[1])
	// left3, _ := strconv.Atoi(resultnum[2])
	// right1, _ := strconv.Atoi(resultnum[3])
	// right2, _ := strconv.Atoi(resultnum[4])
	// right3, _ := strconv.Atoi(resultnum[5])

	// fmt.Println(left1 + left2 + left3)
	// fmt.Println(right1 + right2 + right3)
}

func checkError() {
	if e := recover(); e != nil {
		if strings.Contains(e.(string), "should be even") {
			fmt.Println("angka harus genap")
		}
	}
}

func main() {
	// const num int = 123
	// const name string = "Fazztrack"
	// const isMarried bool = false
	// result := "hello"
	// fmt.Printf("Hello World %d\n%s\nIs iam married? %t \n%s", num, name, isMarried,result)

	// var num1 float32 = keliling(20, 3.14)
	// var num2 float32 = luas(20, 3.14)
	defer checkError()
	s := []string{"2", "3", "4", "5", "6", "7"}
	out(s)

	// fmt.Println("Hasil kalkulsi")
	// fmt.Printf("Hasil keliling lingkaran adalah %g\nHasil luas lingkarang adalah %g", num1, num2)
}
