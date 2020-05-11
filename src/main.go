package main

func findEnd(data []rune, s rune, igrnoe rune, f int, e int) (bool, int) {
	pos := 0
	for i := e; i > f; i-- {
		if data[i] == s {
			return true, i
		}

		if data[i] == igrnoe {
			continue
		}

		return false, i
	}

	return false, pos
}

func isValid(s string) bool {

	d := "(){}[] "
	src := []rune(d)

	data := []rune(s)

	for i := 0; i < len(data); i++ {
		is := false
		for j := 0; j < len(src); j++ {
			if src[j] == data[i] {
				is = true
			}
		}

		if !is {
			return false
		}
	}

	f := 0
	e := len(data) - 1
	for f < e {

		if data[f] == src[6] {
			f++
			continue
		}

		for i := 0; i < 3; i++ {
			if data[f] == src[i*2] {
				ret, pos := findEnd(data, src[i*2+1], src[6], f, e)
				if !ret {
					return false
				}
				e = pos
			}

			if data[f] == src[i*2+1] {
				return false
			}
		}

		f++
	}

	return true
}

func main() {

	println(isValid("{[]}"))
}
