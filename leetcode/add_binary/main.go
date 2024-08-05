package main

import (
	"math/big"
	"strconv"
)

func main() {

}

func addBinary(a string, b string) string {
	num1, _ := new(big.Int).SetString(a, 2)
	num2, _ := new(big.Int).SetString(b, 2)
	num1.Add(num1, num2)
	return num1.Text(2)
}

func binaryConverter(a string) int {
	decimal, err := strconv.ParseInt(a, 2, 124)
	if err != nil {
		return 0
	}
	return int(decimal)
}

func decimalToBinary(decimal int) string {
	return strconv.FormatInt(int64(decimal), 2)
}

// func sqrx(a, b int) int {
// 	res := 1
// 	for i := 0; i < b; i++ {
// 		res *= a
// 	}
// 	return res
// }
