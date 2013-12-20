// Package genome is for operating genome data sequences.
package genome

// Block represents a piece of genome data in a long sequence.
type Block struct {
	Valid       bool
	NumMixedTag int // Number of mixed tag(for complex DiffType).
	Data        []byte
}

// Sequence represents processed genome sequence.
type Sequence struct {
	Blocks []*Block
}

// Length returns the number of blocks in the sequence.
func (s *Sequence) Length() int {
	return len(s.Blocks)
}
