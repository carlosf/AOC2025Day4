package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func p(s string) {
	fmt.Println(s)
}

func CalcMDFive(s []byte) string {

	hash := md5.Sum(s)
	fmt.Printf("hash: %x\n", hash)
	var hashString = ""

	//for _, v := range hash {
	//	p(fmt.Sprintf("%x", v))
	//	hashString += fmt.Sprintf("%x", v)
	//	//p(hashString)
	//}
	return hashString
}

func main() {
	var input = "abcdef" //"ckczppom"
	var postfixNum int64 = 0

	postfixNum = 609043
	hash := CalcMDFive(append([]byte(input), []byte(strconv.FormatInt(postfixNum, 10))...))
	p(hash)

	for {
		postfixNum++
		hash := md5.Sum(append([]byte(input), []byte(strconv.FormatInt(postfixNum, 10))...))
		var firstFiveChars = ""

		for _, v := range hash {
			firstFiveChars += fmt.Sprintf("%x", v)
			//p(firstFiveChars)
		}
		if firstFiveChars[:5] == "00000" {
			fmt.Printf("found solution at %d\n", postfixNum)
			p(firstFiveChars)
			break
		}
		//fmt.Printf("tried %d\n", postfixNum)
	}
}
