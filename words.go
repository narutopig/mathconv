package mathconv

// TensPower lists the powers of ten
var TensPower = map[int]string{
	0: "",
	2: "hundred",
	3: "thousand",
	6: "million",
	9: "billion",
	// int32 limit after about 2 billion
	12: "trillion",
	15: "quadrillion",
	18: "quintillion",
	// int64 limit after about 9 quintillion
	21: "sextillion",
	24: "septillion",
	27: "octillion",
	30: "nonillion",
	33: "decillion",
	36: "undecillion",
	39: "duodecillion",
	42: "tredecillion",
	45: "quattuordecillion",
	48: "quindecillion",
	51: "sexdecillion",
	54: "septendecillion",
	57: "octodecillion",
	60: "novemdecillion",
	63: "vigintillion",
}

// Tens lists 10 * n, where n is between 1 and 9, inclusive
var Tens = map[int]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

// Teens is nums from 11 to 19, inclusive
var Teens = map[int]string{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

// Ones is from 1 to 9, inclusive
var Ones = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}
