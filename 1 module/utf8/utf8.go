package main

import "fmt"

/*
used sources:
1)https://bmstu-iu9.github.io/dm/
2)https://gist.github.com/Miouyouyou/864130e8734afe3f806512b14022226f
3)https://ru.wikipedia.org/wiki/UTF-8
*/

func encode(utf32 []rune) []byte {
	utf8 := make([]byte, 0, len(utf32)/2)
	for i := 0; i < len(utf32); i++ {
		switch {
		case utf32[i] <= 0x7F:
			utf8 = append(utf8, byte(utf32[i]))
		case utf32[i] <= 0x7FF:
			utf8 = append(utf8,
				byte(0b11000000|(utf32[i]>>6)),
				byte(0b10000000|(utf32[i]&0b00111111)))
		case utf32[i] <= 0xFFFF:
			utf8 = append(utf8,
				byte(0b11100000|(utf32[i]>>12)),
				byte(0b10000000|((utf32[i]>>6)&0b00111111)),
				byte(0b10000000|(utf32[i]&0b00111111)))
		case utf32[i] <= 0x10FFFF:
			utf8 = append(utf8,
				byte(0b11110000|(utf32[i]>>18)),
				byte(0b10000000|((utf32[i]>>12)&0b00111111)),
				byte(0b10000000|((utf32[i]>>6)&0b00111111)),
				byte(0b10000000|(utf32[i]&0b00111111)))
		default:
			fmt.Printf("%d is not an Unicode character", utf32[i])
		}
	}
	return utf8
}

func decode(utf8 []byte) []rune {
	utf32 := make([]rune, 0, len(utf8)*2)
	for i := 0; i < len(utf8); {
		switch {
		case (utf8[i] >> 7) == 0: // 1 byte
			utf32 = append(utf32, rune(utf8[i]))
			i++
		case (utf8[i] >> 5) == 6: // 2 bytes
			utf32 = append(utf32,
				(rune(utf8[i]&0b00011111)<<6)|
					(rune(utf8[i+1]&0b00111111)))
			i += 2
		case (utf8[i] >> 4) == 14: // 3 bytes
			utf32 = append(utf32,
				(rune(utf8[i]&0b00001111)<<12)|
					(rune(utf8[i+1]&0b00111111)<<6)|
					(rune(utf8[i+2]&0b00111111)))
			i += 3
		case (utf8[i] >> 3) == 30: // 4 bytes
			utf32 = append(utf32,
				(rune(utf8[i]&0b00000111)<<18)|
					(rune(utf8[i+1]&0b00111111)<<12)|
					(rune(utf8[i+2]&0b00111111)<<6)|
					(rune(utf8[i+3]&0b00111111)))
			i += 4
		}
	}
	return utf32
}

func test() {
	str := "hey раз ꡀꡁꡂꡃꡄꡅ 🛧🛧🛧" //1, 2, 3, 4 byte characters
	testRune := ([]rune)(str)
	testByte := ([]byte)(str)

	fmt.Print("encode test: ")
	if string(testRune) == string(encode(testRune)) {
		fmt.Print("Passed \n")
	} else {
		fmt.Print("Failed \n")
	}
	fmt.Print("decode test: ")
	if string(testByte) == string(decode(testByte)) {
		fmt.Print("Passed \n")
	} else {
		fmt.Print("Failed \n")
	}
}

func main() {
	test()
}
