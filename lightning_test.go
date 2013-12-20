package lightning

import (
	//"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/genomelightning/lightning/genome"
)

func TestComputeDiffSeq(t *testing.T) {
	Convey("Compute DiffType sequence of two processed genome", t, func() {
		Convey("Genome sequences only contains 'Default'", func() {
			gs1 := &genome.Sequence{}
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})

			gs2 := &genome.Sequence{}
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			bs, err := ComputeDiffSeq(gs1, gs2)
			So(err, ShouldBeNil)
			So(bs.DumpAsType(), ShouldEqual,
				"00000000000000000000000000000000\n")
		})

		Convey("Genome sequences only contains 'Simple'", func() {
			gs1 := &genome.Sequence{}
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCACCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGCGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAGAAAACCCCCCCCC"),
			})

			gs2 := &genome.Sequence{}
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			bs, err := ComputeDiffSeq(gs1, gs2)
			So(err, ShouldBeNil)
			So(bs.DumpAsType(), ShouldEqual,
				"11100000000000000000000000000000\n")
		})

		Convey("Genome sequences contains 'Simple' and 'Invalid'", func() {
			gs1 := &genome.Sequence{}
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: false,
				Data:  []byte("GGGGGGGGAAAANAAACCACCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGCGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: false,
				Data:  []byte("GGGGNNGGAAAGAAAACCCCCCCCC"),
			})

			gs2 := &genome.Sequence{}
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: false,
				Data:  []byte("GGGGGGGGAAAAAAAACCNNCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: false,
				Data:  []byte("GGGGGGGNNAAAAAAACCCCCCCCC"),
			})
			bs, err := ComputeDiffSeq(gs1, gs2)
			So(err, ShouldBeNil)
			So(bs.DumpAsType(), ShouldEqual,
				"31300000000000000000000000000000\n")
		})

		Convey("Genome sequences contains 'Simple' and 'Unknown'", func() {
			gs1 := &genome.Sequence{}
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCACCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGCGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAGAAAACCCCCCCCC"),
			})
			gs1.Blocks = append(gs1.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAGAAAACCCCCCCCC"),
			})

			gs2 := &genome.Sequence{}
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid: true,
				Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid:       true,
				NumMixedTag: 1,
				Data:        []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			gs2.Blocks = append(gs2.Blocks, &genome.Block{
				Valid:       true,
				NumMixedTag: 0,
				Data:        []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
			})
			bs, err := ComputeDiffSeq(gs1, gs2)
			So(err, ShouldBeNil)
			So(bs.DumpAsType(), ShouldEqual,
				"12210000000000000000000000000000\n")
		})
	})
}
