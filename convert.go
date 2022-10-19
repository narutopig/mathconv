package mathconv

import (
	"strings"
)

// Convert converts an int64 to a string
func Convert(num int64) string {
	if num == 0 {
		return "zero"
	}

	res := []int{}
	n := num

	for n > 0 {
		res = append(res, (int)(n%1000))
		n /= 1000
	}

	res = reverse(res)

	stuff := []string{}

	for i, val := range res {
		if val != 0 {
			tensPower := 3 * (len(res) - i - 1)

			thing := toString(val)

			if tensPower > 0 {
				thing += " " + TensPower[tensPower]
			}

			stuff = append(stuff, thing)
		}
	}

	return strings.Join(stuff, ", ")
}

func toString(num int) string {
	segs := []string{}

	if num/100 > 0 {
		segs = append(segs, Ones[num/100], TensPower[2])
	}

	remainder := (int)(num % 100)

	teensVal, exists := Teens[remainder]

	if exists {
		segs = append(segs, teensVal)
	} else {
		temp, exists2 := Tens[remainder/10]

		if remainder%10 > 0 {
			if exists2 {
				temp += "-"
			}
			temp += Ones[remainder%10]
		}

		segs = append(segs, temp)

	}

	return strings.Join(segs, " ")
}

func reverse(a []int) []int {
	l := len(a)

	res := make([]int, l)

	for i := 0; i < l; i++ {
		res[i] = a[l-i-1]
	}

	return res
}
