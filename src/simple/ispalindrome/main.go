package main

import (
	"container/list"
	"strconv"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	l := list.New()

	for true {
		if x/10 != 0 {
			l.PushBack(x % 10)
			x = x / 10
			continue
		}
		l.PushBack(x)
		break
	}

	//if l.Len() %2 == 0 {
	//	return false
	//}

	front := l.Front()
	back := l.Back()

	for true {
		if front == back {
			return true
		}

		if front.Value != back.Value {
			return false
		}

		front = front.Next()
		back = back.Prev()
	}

	return false
}

func isPalindromeEx(x int) bool {

	if x < 0 {
		return false
	}

	src := x
	result := 0
	for true {
		if x/10 != 0 {
			result = result*10 + x%10
			x = x / 10
			continue
		}
		result = result*10 + x
		break
	}

	if src == result {
		return true
	}
	return false
}

func isPalindromeStr(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	for i := range str {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func main() {

	val := 11111
	println(val, isPalindrome(val))

	val = 12121
	println(val, isPalindrome(val))

	val = -12121
	println(val, isPalindrome(val))
	println(val, isPalindromeStr(val))

	val = 121121
	println(val, isPalindromeEx(val))
	println(val, isPalindromeStr(val))

}
