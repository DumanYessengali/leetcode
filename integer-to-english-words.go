package main

var (
	lessTwenty = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens       = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousand   = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	var i int

	var str string

	for num > 0 {
		if num%1000 != 0 {
			str = toConvertString(num%1000) + thousand[i] + " " + str
		}
		num /= 1000
		i++
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' {
			return str[:i]
		}
	}
	return str
}

func toConvertString(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return lessTwenty[num] + " "
	} else if num < 100 {
		return tens[num/10] + " " + toConvertString(num%10)
	} else {
		return lessTwenty[num/100] + " Hundred " + toConvertString(num%100)
	}
}
