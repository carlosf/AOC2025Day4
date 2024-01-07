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
	//fmt.Printf("hash: %x\n", hash)
	//p(fmt.Sprintf("%x", hash))
	//p(fmt.Sprintf("%x", hash[:3])[:5])
	var hashString = fmt.Sprintf("%x", hash[:3])[:6]

	//for _, v := range hash {
	//	p(fmt.Sprintf("%x", v))
	//	hashString += fmt.Sprintf("%x", v)
	//	//p(hashString)
	//}
	return hashString
}

func main() {
	var input = "ckczppom"
	var postfixNum int64 = 0

	for {
		postfixNum++
		hash := CalcMDFive(append([]byte(input), []byte(strconv.FormatInt(postfixNum, 10))...))

		if hash == "000000" {
			fmt.Printf("found solution at %d\n", postfixNum)
			p(hash)
			break
		}
	}
}
