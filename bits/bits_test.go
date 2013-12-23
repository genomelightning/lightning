package bits

import (
	//"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_wordsNeeded(t *testing.T) {
	Convey("Compute number of word needed according to length of tails", t, func() {
		Convey("Length that just fits full size of words", func() {
			So(wordsNeeded(32, 2), ShouldEqual, 1)
			So(wordsNeeded(64, 2), ShouldEqual, 2)
			So(wordsNeeded(96, 2), ShouldEqual, 3)
			So(wordsNeeded(math.MaxUint32, 2), ShouldEqual, 134217728)
			So(wordsNeeded(math.MaxUint32-31, 2), ShouldEqual, 134217727)
			So(wordsNeeded(math.MaxUint32-63, 2), ShouldEqual, 134217726)
		})

		Convey("Length that a little less to fit full size of words", func() {
			So(wordsNeeded(0, 2), ShouldEqual, 1)
			So(wordsNeeded(31, 2), ShouldEqual, 1)
			So(wordsNeeded(63, 2), ShouldEqual, 2)
			So(wordsNeeded(95, 2), ShouldEqual, 3)
			So(wordsNeeded(math.MaxUint32-32, 2), ShouldEqual, 134217727)
			So(wordsNeeded(math.MaxUint32-64, 2), ShouldEqual, 134217726)
		})

		Convey("Length that a little more to fit full size of words", func() {
			So(wordsNeeded(33, 2), ShouldEqual, 2)
			So(wordsNeeded(65, 2), ShouldEqual, 3)
			So(wordsNeeded(97, 2), ShouldEqual, 4)
		})
	})

	Convey("Compute number of word needed according to length of combines", t, func() {
		Convey("Length that just fits full size of words", func() {
			So(wordsNeeded(16, 4), ShouldEqual, 1)
			So(wordsNeeded(32, 4), ShouldEqual, 2)
			So(wordsNeeded(48, 4), ShouldEqual, 3)
			So(wordsNeeded(math.MaxUint32, 4), ShouldEqual, 268435456)
			So(wordsNeeded(math.MaxUint32-15, 4), ShouldEqual, 268435455)
			So(wordsNeeded(math.MaxUint32-31, 4), ShouldEqual, 268435454)
		})

		Convey("Length that a little less to fit full size of words", func() {
			So(wordsNeeded(0, 4), ShouldEqual, 1)
			So(wordsNeeded(15, 4), ShouldEqual, 1)
			So(wordsNeeded(31, 4), ShouldEqual, 2)
			So(wordsNeeded(47, 4), ShouldEqual, 3)
			So(wordsNeeded(math.MaxUint32-16, 4), ShouldEqual, 268435455)
			So(wordsNeeded(math.MaxUint32-32, 4), ShouldEqual, 268435454)
		})

		Convey("Length that a little more to fit full size of words", func() {
			So(wordsNeeded(17, 4), ShouldEqual, 2)
			So(wordsNeeded(33, 4), ShouldEqual, 3)
			So(wordsNeeded(49, 4), ShouldEqual, 4)
		})
	})
}

func TestNew(t *testing.T) {
	Convey("Create new bit sequence with given length", t, func() {
		Convey("Create a bit sequence that has 0 tail", func() {
			bs := New(0)
			So(bs.length, ShouldEqual, 0)
			So(len(bs.words), ShouldEqual, 1)
			So(len(bs.combines), ShouldEqual, 1)
		})

		Convey("Create a bit sequence that has 1 tail", func() {
			bs := New(1)
			So(bs.length, ShouldEqual, 1)
			So(len(bs.words), ShouldEqual, 1)
			So(len(bs.combines), ShouldEqual, 1)
		})

		Convey("Create a bit sequence that has 32 tails", func() {
			bs := New(32)
			So(bs.length, ShouldEqual, 32)
			So(len(bs.words), ShouldEqual, 1)
			So(len(bs.combines), ShouldEqual, 2)
		})

		Convey("Create a bit sequence that has 33 tails", func() {
			bs := New(33)
			So(bs.length, ShouldEqual, 33)
			So(len(bs.words), ShouldEqual, 2)
			So(len(bs.combines), ShouldEqual, 3)
		})
	})
}

func TestSet(t *testing.T) {
	Convey("Set unit value of given index according to DiffType", t, func() {
		Convey("Ordered set for 4 different types", func() {
			bs := New(1)
			bs.Set(0, DT_DEFAULT, 1, 1)
			bs.Set(1, DT_SIMPLE, 1, 2)
			bs.Set(2, DT_COMPLEX, 1, 2)
			bs.Set(3, DT_UNKNOWN, 0, 0)
			So(bs.DumpWordsAsBits(), ShouldEqual,
				"0001101100000000000000000000000000000000000000000000000000000000\n")
		})

		Convey("Overwrite set for 4 different types", func() {
			bs := New(1)
			bs.Set(0, DT_DEFAULT, 1, 1)
			bs.Set(0, DT_SIMPLE, 1, 2)
			So(bs.DumpWordsAsBits(), ShouldEqual,
				"0100000000000000000000000000000000000000000000000000000000000000\n")
			bs.Set(0, DT_COMPLEX, 1, 2)
			So(bs.DumpWordsAsBits(), ShouldEqual,
				"1000000000000000000000000000000000000000000000000000000000000000\n")
			bs.Set(0, DT_UNKNOWN, 0, 0)
			So(bs.DumpWordsAsBits(), ShouldEqual,
				"1100000000000000000000000000000000000000000000000000000000000000\n")
			bs.Set(0, DT_DEFAULT, 1, 1)
			So(bs.DumpWordsAsBits(), ShouldEqual,
				"0000000000000000000000000000000000000000000000000000000000000000\n")
		})
	})
}

func TestGet(t *testing.T) {
	Convey("Get tail value by given index", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE, 1, 2)
		So(bs.Get(5), ShouldEqual, DT_SIMPLE)
	})
}

func TestGetCombine(t *testing.T) {
	Convey("Get index in CombinationTable by given index", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE, 1, 2)
		So(bs.GetCombine(5), ShouldEqual, 2)
	})
}

func TestDumpWordsAsBits(t *testing.T) {
	Convey("Convert bit sequence to string format(bits form)", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE, 1, 2)
		bs.Set(8, DT_UNKNOWN, 0, 0)
		bs.Set(10, DT_COMPLEX, 1, 2)
		So(bs.DumpWordsAsBits(), ShouldEqual,
			"0000000000010000110010000000000000000000000000000000000000000000\n")
	})
}

func TestDumpCombinesAsBits(t *testing.T) {
	Convey("Convert bit sequence combine indexes to string format(bits form)", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE, 1, 2)
		bs.Set(8, DT_UNKNOWN, 0, 0)
		bs.Set(10, DT_COMPLEX, 1, 2)
		So(bs.DumpCombinesAsBits(), ShouldEqual,
			"0000000000000000000001000000000000000000010000000000000000000000\n")
	})
}

func TestDumpWordsAsType(t *testing.T) {
	Convey("Convert bit sequence to string format(DiffType form)", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE, 1, 2)
		bs.Set(8, DT_UNKNOWN, 0, 0)
		bs.Set(10, DT_COMPLEX, 1, 2)
		So(bs.DumpWordsAsType(), ShouldEqual,
			"00000100302000000000000000000000\n")
	})
}

func TestDumpCombinesAsType(t *testing.T) {
	Convey("Convert bit sequence combine indexes to string format(decimal form)", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE, 1, 2)
		bs.Set(8, DT_UNKNOWN, 0, 0)
		bs.Set(10, DT_COMPLEX, 1, 2)
		So(bs.DumpCombinesAsType(), ShouldEqual,
			"0000020000200000\n")
	})
}
