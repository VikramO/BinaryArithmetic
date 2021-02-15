package main

import (
	"strings"
)

func Add(num1 string, num2 string) string {

	len_1 := len([]rune(num1))
	len_2 := len([]rune(num2))

	max_length := Max(len_1, len_2)
	min_length := Min(len_1, len_2)

	diff := max_length - min_length

	if len_1 > len_2 {
		str_2_rev := Reverse(num2)
		append := strings.Repeat("0", diff)
		str_2_rev = str_2_rev + append
		num2 = Reverse(str_2_rev)
	} else if len_1 < len_2 {
		str_1_rev := Reverse(num1)
		append := strings.Repeat("0", diff)
		str_1_rev = str_1_rev + append
		num1 = Reverse(str_1_rev)
	} else {

	}

	var sub_sums []string
	var carries []string

	for i := max_length - 1; i > -1; i-- {
		integ_1 := string([]rune(num1)[i])
		integ_2 := string([]rune(num2)[i])
		group := integ_1 + integ_2

		if strings.Contains(group, "1") == true {
			if strings.Count(group, "1") == 2 {
				sub_sums = append(sub_sums, "0")
				carries = append(carries, "1")
			} else {
				sub_sums = append(sub_sums, "1")
				carries = append(carries, "0")
			}
		} else {
			sub_sums = append(sub_sums, "0")
			carries = append(carries, "0")
		}
	}

	var total_sum []string

	total_sum = append(total_sum, sub_sums[0])

	for j := 0; j < max_length-1; j++ {
		sum := sub_sums[j+1]
		carry := carries[j]
		sub_group := string(sum + carry)

		if strings.Contains(sub_group, "1") == true {
			if strings.Count(sub_group, "1") == 2 {
				total_sum = append(total_sum, "0", "1")
			} else {
				total_sum = append(total_sum, "1")
			}
		} else {
			total_sum = append(total_sum, "0")
		}

	}

	string_sum := ""

	for k := len(total_sum) - 1; k > -1; k-- {
		string_sum = string_sum + total_sum[k]
	}

	return string_sum
}
