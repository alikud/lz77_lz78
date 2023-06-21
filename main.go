package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Offset int
	Length int
	Next   string
}

func lastIndex(s, substring string) int {
	for i := len(s) - len(substring); i >= 0; i-- {
		if s[i:i+len(substring)] == substring {
			return i
		}
	}
	return -1
}

func compressLZ77(input string) []Node {
	var compressed []Node
	windowSize := 10
	lookaheadSize := 10

	pos := 0

	for pos < len(input) {
		var match Node

		startIndex := pos - windowSize
		if startIndex < 0 {
			startIndex = 0
		}
		endIndex := pos + lookaheadSize
		if endIndex > len(input) {
			endIndex = len(input)
		}

		searchArea := input[startIndex:pos]
		searchString := input[pos:endIndex]

		// search most common string
		for i := len(searchString); i > 0; i-- {
			substring := searchString[:i]
			matchIndex := lastIndex(searchArea, substring)

			if matchIndex != -1 {
				match.Offset = pos - (startIndex + matchIndex)
				match.Length = len(substring)
				if pos+match.Length < len(input) {
					match.Next = string(input[pos+match.Length])
				}

				break
			}
		}

		//fmt.Printf("Match string %v: ", match)

		if match.Length > 0 {
			compressed = append(compressed, match)
			pos += match.Length + 1
		} else {
			compressed = append(compressed, Node{0, 0, string(input[pos])})
			pos++
		}
	}

	return compressed
}

func decompressLZ77(compressed []Node) string {
	var decompressed strings.Builder

	for _, tuple := range compressed {
		if tuple.Length == 0 {
			decompressed.WriteString(tuple.Next)
		} else {
			startIndex := decompressed.Len() - tuple.Offset
			for i := 0; i < tuple.Length; i++ {
				ch := decompressed.String()[startIndex+i]
				decompressed.WriteByte(ch)
			}
			decompressed.WriteString(tuple.Next)
		}
	}
	return decompressed.String()
}

func compressLZ78(input string) string {
	res := ""
	codes_map := make(map[string]string)
	entry := ""
	index := 1

	for _, char := range input {
		entry += string(char)

		if _, ok := codes_map[entry]; !ok {

			codes_map[entry] = strconv.Itoa(index)

			if len(entry) == 1 {
				res += "0" + entry
			} else {
				entry_index := codes_map[entry[:len(entry)-1]]
				res += entry_index + string(entry[len(entry)-1])
			}

			index++
			entry = ""
		}
	}

	return res
}

func decompresslz78(input string) string {
	mapCodes := make(map[string]string)
	mapCodes["0"] = ""
	res := ""
	entry := ""
	index := 1

	for _, char := range input {
		if strings.Contains("0123456789", string(char)) {
			entry += string(char)
		} else {
			mapCodes[strconv.Itoa(index)] = mapCodes[entry] + string(char)

			res += mapCodes[entry] + string(char)
			entry = ""
			index++
		}
	}

	return res
}

func main() {
	//fmt.Println(lastIndex("pabcdea", "a")) //6
	compressed := compressLZ77("pabcdeqabcde")
	fmt.Println("Compressed lz77:", compressed)
	fmt.Println("Decompressed lz77:", decompressLZ77(compressed))

	compressed_lz78 := compressLZ78("abbcbcababcaabcaab")
	fmt.Println("Compressed lz78", compressed_lz78)
	fmt.Println("Decompressed lz78", decompresslz78(compressed_lz78))
}
