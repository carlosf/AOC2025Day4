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
		fmt.Printf("current number is %v\n", postfixNum)
		fmt.Printf("length of hash is %d\n", len(hash))
		// for statement for length of hash.

		var sumFirstFiveChars string = ""

		for _, v := range hash {
			sumFirstFiveChars += fmt.Sprintf("%x", v)
			p(sumFirstFiveChars)
		}
		p(sumFirstFiveChars[:5])

		fmt.Printf("sumFirstFiveChars is %d\n", sumFirstFiveChars[0])

		if true {
			fmt.Printf("hash 0 is %x\n", hash[0])
			fmt.Printf("hash 1 is %x\n", hash[1])
			fmt.Printf("hash 2 is %x\n", hash[2])
			fmt.Printf("md5 is %x", hash)
			break
		}
	}
}
