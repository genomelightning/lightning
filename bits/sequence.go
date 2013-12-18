// Package bits if for operating bit sequences with 2-bit as a unit.
package bits

import (
	"bytes"
	"fmt"
)

const (
	wordSize     = uint64(64) // Word size of a bit set.
	log2WordSize = uint(6)    // for laster arith.
)

// wordsNeeded computes how many words needed according to the length of units.
func wordsNeeded(i uint32) uint64 {
	if i == 0 {
		return 1
	}
	return (uint64(i)*2 + (wordSize - 1)) >> log2WordSize
}

type DiffType int

const (
	DT_DEFAULT DiffType = iota
	DT_SIMPLE
	DT_COMPLEX
	DT_UNKNOWN
)

// Sequence represents bit sequence that use 2 bits as a unit.
type Sequence struct {
	length uint32
	words  []uint64
}

// New initializes a new 2-bit unit sequence.
func New(length uint32) *Sequence {
	return &Sequence{
		length: length,
		words:  make([]uint64, wordsNeeded(length)),
	}
}

// Set sets unit value of given index according to DiffType.
//
// 	Default - 00
// 	Simple  - 01
// 	Complex - 10
// 	Unknown - 11
func (s *Sequence) Set(i uint64, dt DiffType) *Sequence {
	// TODO: extend size as needed.

	i *= 2
	index := i >> log2WordSize
	baseSize := i & (wordSize - 1)

	switch dt {
	case DT_DEFAULT:
		s.words[index] &^= 1 << (baseSize)
		s.words[index] &^= 1 << (baseSize + 1)
	case DT_SIMPLE:
		s.words[index] &^= 1 << baseSize
		s.words[index] |= 1 << (baseSize + 1)
	case DT_COMPLEX:
		s.words[index] |= 1 << (baseSize)
		s.words[index] &^= 1 << (baseSize + 1)
	case DT_UNKNOWN:
		s.words[index] |= 1 << baseSize
		s.words[index] |= 1 << (baseSize + 1)
	}
	//fmt.Printf("length in bits: %d, real size of sets: %d, bits: %d, index: %d\n", s.length, len(s.words), i, index)
	return s
}

// Get returns unit value by given index.
func (s *Sequence) Get(i uint32) DiffType {
	i *= 2
	index := i >> log2WordSize
	baseSize := uint64(i) & (wordSize - 1)
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

func reverse(str string) string {
	byt1 := []byte(str)
	l := len(byt1) - 1
	byt2 := make([]byte, l+1)
	for i := range byt1 {
		byt2[l-i] = byt1[i]
	}
	return string(byt2)
}

// DumpAsBits converts unit values to string format(bits form):
func (s *Sequence) DumpAsBits() string {
	buf := bytes.NewBufferString("")
	l := int(wordsNeeded(s.length))
	for i := 0; i < l; i++ {
		buf.WriteString(reverse(fmt.Sprintf("%064b", s.words[i])))
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}

// DumpAsType converts unit values to string format(DiffType form):
//
// 	0 - Default
// 	1 - Simple
// 	2 - Complex
// 	3 - Unknown
func (s *Sequence) DumpAsType() string {
	buf := bytes.NewBufferString("")
	l := int(wordsNeeded(s.length))
	for i := 0; i < l; i++ {
		for j := 0; j < 32; j += 1 {
			buf.WriteString(string(s.Get(uint32(i*31+j)) + 48))
		}
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}
