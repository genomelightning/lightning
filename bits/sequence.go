// Package bits if for operating bit sequences with 2-bit as a tile.
package bits

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

const (
	wordSize     = uint64(64) // Word size of a bit set.
	log2WordSize = uint(6)    // for laster arith.
)

// wordsNeeded computes how many words needed according to the length of tiles.
func wordsNeeded(i uint32, size int) uint64 {
	if i == 0 {
		return 1
	}
	return (uint64(i)*uint64(size) + (wordSize - 1)) >> log2WordSize
}

type DiffType int

const (
	DT_DEFAULT DiffType = iota
	DT_SIMPLE
	DT_COMPLEX
	DT_UNKNOWN
)

// Sequence represents bit sequence that use 2 bits as a tile.
type Sequence struct {
	length   uint32
	words    []uint64
	combines []uint64
}

// New initializes a new 2-bit tile sequence.
func New(length uint32) *Sequence {
	return &Sequence{
		length:   length,
		words:    make([]uint64, wordsNeeded(length, 2)),
		combines: make([]uint64, wordsNeeded(length, 4)),
	}
}

// Set sets tile value and combination of given index according to DiffType.
//
// 	Default - 00
// 	Simple  - 01
// 	Complex - 10
// 	Unknown - 11
func (s *Sequence) Set(i uint64, dt DiffType, num1, num2 int) *Sequence {
	// TODO: extend size as needed.

	i *= 2
	index := i >> log2WordSize
	baseSize := i & (wordSize - 1)

	switch dt {
	case DT_DEFAULT:
		s.words[index] &^= 1 << baseSize
		s.words[index] &^= 1 << (baseSize + 1)
	case DT_SIMPLE:
		s.words[index] &^= 1 << baseSize
		s.words[index] |= 1 << (baseSize + 1)
	case DT_COMPLEX:
		s.words[index] |= 1 << baseSize
		s.words[index] &^= 1 << (baseSize + 1)
	case DT_UNKNOWN:
		s.words[index] |= 1 << baseSize
		s.words[index] |= 1 << (baseSize + 1)
	}
	//fmt.Printf("length in bits: %d, real size of sets: %d, bits: %d, index: %d\n", s.length, len(s.words), i, index)

	i *= 2
	index = i >> log2WordSize
	baseSize = i & (wordSize - 1)
	combineIndex := fmt.Sprintf("%04s", strconv.FormatInt(int64(GetCombineTableIndex(num1, num2)), 2))
	//fmt.Printf("combine index: %s, number: %d, index: %d\n", combineIndex, i, index)
	var j uint64 = 0
	for ; j < 4; j++ {
		if combineIndex[3-j] > 48 {
			s.combines[index] |= 1 << (baseSize + j)
		} else {
			s.combines[index] &^= 1 << (baseSize + j)
		}
	}
	//fmt.Println(reverse(fmt.Sprintf("%064b", s.combines[index])))
	return s
}

// Get returns tile value by given index.
func (s *Sequence) Get(i uint64) DiffType {
	i *= 2
	index := i >> log2WordSize
	baseSize := i & (wordSize - 1)
	low := s.words[index] & (1 << baseSize)
	high := s.words[index] & (1 << (baseSize + 1))
	switch {
	case low == 0 && high != 0:
		return DT_SIMPLE
	case low != 0 && high == 0:
		return DT_COMPLEX
	case low != 0 && high != 0:
		return DT_UNKNOWN
	}
	return DT_DEFAULT
}

// GetCombine returns index in CombinationTable by given index of tile.
func (s *Sequence) GetCombine(i uint64) int {
	i *= 4
	index := i >> log2WordSize
	baseSize := i & (wordSize - 1)
	num := 0

	var j uint64 = 0
	for ; j < 4; j++ {
		if s.combines[index]&(1<<(baseSize+j)) != 0 {
			num += int(math.Pow(2, float64(j)))
		}
		//fmt.Println(num, s.combines[index]&(1<<(baseSize+j)), int(math.Pow(2, float64(j))))
	}
	return num
}

func reverse(str string) string {
	byt1 := []byte(str)
	l := len(byt1) - 1
	byt2 := make([]byte, l+1)
	for i := range byt1 {
		byt2[l-i] = byt1[i]
	}
	return string(byt2)
}

// DumpWordsAsBits converts tile values to string format(bits form):
func (s *Sequence) DumpWordsAsBits() string {
	buf := bytes.NewBufferString("")
	l := wordsNeeded(s.length, 2)
	var i uint64 = 0
	for ; i < l; i++ {
		buf.WriteString(reverse(fmt.Sprintf("%064b", s.words[i])))
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}

// DumpCombinesAsBits converts combine indexes to string format(bits form):
func (s *Sequence) DumpCombinesAsBits() string {
	buf := bytes.NewBufferString("")
	l := wordsNeeded(s.length, 4)
	var i uint64 = 0
	for ; i < l; i++ {
		buf.WriteString(reverse(fmt.Sprintf("%064b", s.combines[i])))
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}

// DumpWordsAsType converts tile values to string format(DiffType form):
//
// 	0 - Default
// 	1 - Simple
// 	2 - Complex
// 	3 - Unknown
func (s *Sequence) DumpWordsAsType() string {
	buf := bytes.NewBufferString("")
	l := wordsNeeded(s.length, 2)
	var i uint64 = 0
	for ; i < l; i++ {
		var j uint64 = 0
		for ; j < 32; j += 1 {
			buf.WriteString(string(s.Get(i*31+j) + 48))
		}
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}

// DumpCombinesAsType converts combine indexes to string format(decimal form):
//
// 	0 - Default
// 	1 - Simple
// 	2 - Complex
// 	3 - Unknown
func (s *Sequence) DumpCombinesAsType() string {
	buf := bytes.NewBufferString("")
	l := wordsNeeded(s.length, 4)
	var i uint64 = 0
	for ; i < l; i++ {
		var j uint64 = 0
		for ; j < 16; j += 1 {
			buf.WriteString(string(s.GetCombine(i*15+j) + 48))
		}
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}
