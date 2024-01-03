package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func p(s string) {
	fmt.Println(s)
}

func main() {
	var input string = "ckczppom"
	var postfixNum int64 = 0

	for {
		postfixNum++
		hash := md5.Sum(append([]byte(input), []byte(strconv.FormatInt(postfixNum, 10))...))
		p(string(postfixNum))

		var sumFirstFiveChars = hash[:5]
		fmt.Printf("sumFirstFiveChars is %x\n", sumFirstFiveChars[0])

		if false {
			fmt.Printf("hash 0 is %x\n", hash[0])
			fmt.Printf("hash 1 is %x\n", hash[1])
			fmt.Printf("hash 2 is %x\n", hash[2])
			fmt.Printf("md5 is %x", hash)
			break
		}
	}
}
