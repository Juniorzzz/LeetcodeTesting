package main

func maximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	if len(matrix[0]) == 0 {
		return 0
	}

	h, v, max := 0, 0, 0
	isSquare := false

	for {
		if matrix[h][v] == '1' {
			index := 0
			isSquare = true

			if max == 0 {
				max = 1
			}

			for {
				index++

				if h+index >= len(matrix) || v+index >= len(matrix[0]) {
					index--
					break
				}

				for i := 0; i < index; i++ {
					if matrix[h+index][v+i] == '0' || matrix[h+i][v+index] == '0' {
						isSquare = false
						break
					}
				}

				if !isSquare {
					break
				}

				if matrix[h+index][v+index] == '0' || matrix[h][v+index] == '0' || matrix[h+index][v] == '0' {
					index--
					break
				}

				if max < (index+1)*(index+1) && isSquare {
					max = (index + 1) * (index + 1)
				}
			}
		}

		if v != len(matrix[h])-1 {
			v++
			continue
		}

		if h != len(matrix)-1 {
			h++
			v = 0
			continue
		}

		break
	}

	return max
}

func main() {
	//matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}

	matrix := [][]byte{{'1'}}
	println(maximalSquare(matrix))
}
