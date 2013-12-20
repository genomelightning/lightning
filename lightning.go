// Package lighting is for processing data of genome.
package lightning

import (
	//"fmt"
	//"errors"

	"github.com/genomelightning/lightning/bits"
	"github.com/genomelightning/lightning/genome"
)

// ComputeDiffSeq compares two processed genome squences and computes bit sequence of differences.
func ComputeDiffSeq(gs1, gs2 *genome.Sequence) (*bits.Sequence, error) {
	// TODO: Concern both of two sequences have midxed tags.

	// NOTE: Current implementatin can only handle mixed tags appear in gs2.
	bs := bits.New(uint32(gs1.Length()))

	index := 0
	skipNum := 0
	isHasMix := false

	for i := 0; i < gs1.Length(); i++ {
		// Complex.
		if skipNum > 0 {
			skipNum--
			//fmt.Println("DT_COMPLEX")
			bs.Set(uint64(i), bits.DT_COMPLEX)
			continue
		}

		if !isHasMix {
			skipNum = gs2.Blocks[index].NumMixedTag
		}

		if skipNum > 0 {
			isHasMix = true
			bs.Set(uint64(i), bits.DT_COMPLEX)
			//fmt.Println("DT_COMPLEX2")
			continue
		}

		// Invalid.
		// NOTE: Here we assume invalid does not mix with complex.
		if !gs1.Blocks[i].Valid || !gs2.Blocks[index].Valid {
			//fmt.Println("DT_UNKNOWN")
			bs.Set(uint64(i), bits.DT_UNKNOWN)
			continue
		}

		isHasMix = false

		// Non-complex.
		str1, str2 := string(gs1.Blocks[i].Data), string(gs2.Blocks[index].Data)
		if str1 == str2 {
			//fmt.Println("DT_DEFAULT")
			bs.Set(uint64(i), bits.DT_DEFAULT)
		} else {
			//fmt.Println("DT_SIMPLE")
			bs.Set(uint64(i), bits.DT_SIMPLE)
		}
		index++
	}

	return bs, nil
}
