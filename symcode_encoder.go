package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"bufio"
)

const Cipher = "symantec"
const PadValue = "$"
const OutputFile = "output.txt"

func main() {
	inputFile := "basic.txt"
	read, err := os.Open(inputFile)
	check(err)
	defer read.Close()

	write, err := os.Create(OutputFile)
	check(err)
	defer write.Close()

	output := bufio.NewWriter(write)
	scanner := bufio.NewScanner(read)

	// read one line of input and write one line of output
	for scanner.Scan() {
		text := scanner.Text()
		encoded := EncodeString(text)
		fmt.Fprintln(output, encoded)
	}
	output.Flush() // need to flush output after using it
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func EncodeString(s string) string {
	// first need to strip leading and trailing whitespace
	s = strings.TrimSpace(s)

	// create the binary representation for the ASCII string
	// should store as individual bits, but this is easier to convert
	// the whole string to
	bin := ""
	for _, c := range s {
		bin = fmt.Sprintf("%s%.8b", bin, c)
	}

	// need to apply padding if necessary
	pad := GetPad(bin)

	bin = PadBinaryString(bin)
	
	// need to take the binary string and break it up into 3 bit chunks
	// each group of 3 bits is used to determine the cipher character
	shift := 2 // because we're reading left to right, shift is backwards
	result := ""
	bitValue := 0

	for i := 0; i <= len(bin); i++ {
		if i % 3 == 0 && i > 0 {
			// if we are at the point of padding, don't want to encode characters anymore
			if i > pad {
				result += PadValue
			} else {
				result += EncodeChar(bitValue)
				bitValue = 0
				shift = 2
			}
		}
		if i < len(bin) {
			v, err := strconv.ParseInt(string(bin[i]), 10, 32)
			check(err)
			value := int(v << uint(shift)) // the shift operation results in a uint value
			bitValue += value
			shift--
		}
	}
	return result
}

/*
	given a string, determine the index at which padding starts
	this would be when len(s) % 3 == 0 next
*/
func GetPad(s string) int {
	count := len(s)
	for count % 3 != 0 {
		count++
	}
	return count
}

/*
	for a given binary string, pad it until it can be 
	evenly divided by 24 -> divisble by 8 and 3
*/
func PadBinaryString(s string) string {
	for len(s) % 24 != 0 {
		s += "0"
	}
	return s
}

/*
	takes an integer value (the integer value for a 3-bit number)
	and returns the encoded character for it
	s - 0
	y - 1
	m - 2
	a - 3
	n - 4
	t - 5
	e - 6
	c - 7

	note: value is expected to conform to
	0 <= n <= 7
*/
func EncodeChar(value int) string{
	return string(Cipher[value])
}
