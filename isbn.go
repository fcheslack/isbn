package isbn

import (
	//"errors"
	"strconv"
	//"strings"
	"log"
	"unicode"
)

const (
	ISBN10to13Prefix = "978"
)

type ISBN string

func Normalize(in string) string {
	//log.Print("input: " + in)
	s := []rune{}
	inr := []rune(in)
	for _, c := range inr {
		if unicode.IsDigit(c) {
			//log.Print("digit, appending")
			s = append(s, c)
		}
	}

	if len(s) == 10 {
		//validate ISBN-10
		if checksum10(string(s)) {
			s = append([]rune(ISBN10to13Prefix), s...)
			check13 := checksum13digit(string(s))
			s[12] = []rune(strconv.Itoa(check13))[0]
			//log.Printf("isbn10 was valid, normalized isbn13: %s", string(s))
			if checksum13(string(s)) {
				return string(s)
			} else {
				return ""
			}
		}
	}

	if len(s) != 13 {
		return ""
	}

	if !checksum13(string(s)) {
		return ""
	}

	//log.Print("output: " + string(s))
	return string(s)
}

//verify the checksum character in a 13 digit string ISBN
//assumes that the string is pre-validated as at least 13 characters that are all digits
func checksum13(s string) bool {
	if len(s) != 13 {
		return false
	}

	for _, c := range s {
		if !unicode.IsDigit(c) {
			log.Print("non-digit rune in checksum13")
			return false
		}
	}
	checkint, err := strconv.Atoi(string(s[12]))
	if err != nil {
		return false
	}

	checkdigit := checksum13digit(s)

	if checkdigit == checkint {
		return true
	}

	return false
}

func checksum13digit(s string) int {
	sum := 0

	if len(s) < 12 {
		return -1
	}

	for i := 0; i < 12; i++ {
		toInt, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return -1
		}
		if i%2 == 0 {
			sum += toInt
		} else {
			sum += toInt * 3
		}
	}

	checkdigit := (10 - sum%10) % 10
	return checkdigit
}

func checksum10(s string) bool {
	if len(s) < 10 {
		return false
	}

	var checkint int
	var err error
	if string(s[9]) == "X" {
		checkint = 10
	} else {
		checkint, err = strconv.Atoi(string(s[9]))
		if err != nil {
			return false
		}
	}

	checkdigit := checksum10digit(s)

	if checkdigit == checkint {
		return true
	}

	return false
}

func checksum10digit(s string) int {
	sum := 0

	if len(s) < 9 {
		return -1
	}

	for i := 0; i < 9; i++ {
		toInt, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return -1
		}
		sum += (10 - i) * toInt
	}

	checkdigit := 11 - sum%11
	return checkdigit
}
