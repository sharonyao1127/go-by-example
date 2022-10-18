package main

import (
	"fmt"
	"strings"
)

//Go 语言的字符有以下两种：
//uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
//rune类型，代表一个 UTF-8字符。

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // true
	fmt.Println(strings.Count(a, "l"))                    // 2
	fmt.Println(strings.HasPrefix(a, "he"))               // true
	fmt.Println(strings.HasSuffix(a, "llo"))              // true
	fmt.Println(strings.Index(a, "ll"))                   // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	fmt.Println(strings.Repeat(a, 2))                     // hellohello
	fmt.Println(strings.Replace(a, "e", "E", -1))         // hEllo
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]
	fmt.Println(strings.ToLower(a))                       // hello
	fmt.Println(strings.ToUpper(a))                       // HELLO
	fmt.Println(len(a))                                   // 5
	b := "你好"
	fmt.Println(len(b)) // 6

	// string底层就是一个byte的数组，因此，也可以进行切片操作。
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)

	//修改字符串
	//需要先将其转换成[]rune或[]byte，完成后再转换为string
	str1 := "Hello world"
	ss := []byte(str1)
	ss[6] = 'G'
	ss = ss[:8]
	ss = append(ss, '!')
	str = string(ss)
	fmt.Println(str1)

	//含中文
	str2 := "你好，世界！hello world！"
	s := []rune(str2) //当需要处理中文、日文或者其他复合字符时，则需要用到rune类型
	s[3] = '够'
	s[4] = '浪'
	s[12] = 'g'
	s = s[:14]
	str = string(s)
	fmt.Println(str2)
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

	for _, r := range s { //rune
		//当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}
