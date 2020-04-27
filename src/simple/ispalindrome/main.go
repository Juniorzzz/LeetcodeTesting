package main

import "container/list"

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

func main() {

	val := 11111
	println(val, isPalindrome(val))

	val = 12121
	println(val, isPalindrome(val))

	val = -12121
	println(val, isPalindrome(val))
	21
	val = 121
	println(val, isPalindrome(val))
}
