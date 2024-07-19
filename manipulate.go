package main

func Manipulate(text []string) []string {
	duplicate := make(map[string]struct{})
	for _, data := range text {

		strData := []rune(data)

		var Shuffle func(int, int)

		Shuffle = func(l, r int) {
			if l == r {
				duplicate[string(strData)] = struct{}{}
			} else {
				for i := l; i <= r; i++ {
					strData[l], strData[i] = strData[i], strData[l]
					Shuffle(l+1, r)
					strData[l], strData[i] = strData[i], strData[l]
				}
			}
		}
		Shuffle(0, len(strData)-1)
	}

	var result []string
	for dataRes := range duplicate {
		result = append(result, dataRes)
	}

	return result

}
