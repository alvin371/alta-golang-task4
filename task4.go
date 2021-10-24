package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) []string {

	// your code here
	return strings.Split(a, b)
}

func swap(px, py *int) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncryptPlaintext(key int, plaintext string) string {
	return rotateText(plaintext, key)
}

func rotateText(inputText string, rot int) string {
	rot %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			rotatedText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))+rot)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			rotatedText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))+rot)%26]
		}
	}
	return string(rotatedText)
}

// Min Max
func getMinMax(numbers ...*int) (min int, max int) {
	nilaiMinimum := numbers[0]
	nilaiMax := numbers[0]

	for _, i := range numbers {
		if nilaiMinimum < i {
			nilaiMinimum = i
		}
	}

	for _, i := range numbers {
		if nilaiMax > i {
			nilaiMax = i
		}
	}

	return *nilaiMinimum, *nilaiMax

}

// Student Score
type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	avg := 0
	for _, v := range s.score {
		avg += v
	}
	return float64(avg) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {

	nilaiMinimum := s.score[0]
	nama := s.name[0]

	// O(N) complexity!
	for _, v := range s.score {
		if v < nilaiMinimum {
			nilaiMinimum = v
		}
	}

	return nilaiMinimum, nama
}

func (s Student) Max() (max int, name string) {
	nilaiMax := s.score[0]
	nama := s.name[0]

	// O(N) complexity!
	for _, v := range s.score {
		if v > nilaiMax {
			nilaiMax = v
		}
	}

	return nilaiMax, nama
}

// Chipper

type student struct {
	name string

	nameEncode string

	score int
}

type Chiper interface {
	Encode() string

	Decode() string
}

func (s *student) Encode() string {

	var nameEncode = ""

	return cipher(nameEncode, +1)

}

func (s *student) Decode() string {

	var nameDecode = ""

	return cipher(nameDecode, -1)

}
func cipher(text string, direction int) string {
	shift, offset := rune(3), rune(26)

	runes := []rune(text)

	for index, char := range runes {
		switch direction {
		case -1: // encoding
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1: // decoding
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}
		runes[index] = char
	}

	return string(runes)
}

func main() {
	// Problem 1 – Compare String
	fmt.Println(Compare("AKASHI", "AKA"))

	// Problem 2 – Caesar Cipher
	fmt.Println(EncryptPlaintext(3, "abc")) // def

	fmt.Println(EncryptPlaintext(2, "alta")) // cnvc

	fmt.Println(EncryptPlaintext(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(EncryptPlaintext(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(EncryptPlaintext(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

	// Problem 3 – Swap Two Number Using Pointer
	x := int(1)
	y := int(2)
	fmt.Println("x was", x)
	fmt.Println("y was", y)
	swap(&x, &y)

	fmt.Println("x is now", x)
	fmt.Println("y is now", y)

	// Problem 4 – Min and Max Using Pointer

	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
	// Problem 5 – Students Score
	var a = Student{}
	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("input " + string(i) + " Student’s Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+"(", scoreMin, ")")
	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+"(", scoreMax, ")")

	// Problem 6 – Substitution Cipher
	var menu int

	var chip = student{}

	var c Chiper = &chip

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&chip.name)

		fmt.Print("\nEncode of Student’s Name " + chip.name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&chip.nameEncode)

		fmt.Print("\nDecode of Student’s Name " + chip.nameEncode + " is : " + c.Decode())

	} else {

		fmt.Println("Wrong input name menu")

	}
}
