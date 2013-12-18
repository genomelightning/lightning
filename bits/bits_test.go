package bits

import (
	//"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_wordsNeeded(t *testing.T) {
	Convey("Compute number of word needed according to length of units", t, func() {
		Convey("Length that just fits full size of words", func() {
			So(wordsNeeded(32), ShouldEqual, 1)
			So(wordsNeeded(64), ShouldEqual, 2)
			So(wordsNeeded(96), ShouldEqual, 3)
			So(wordsNeeded(math.MaxUint32), ShouldEqual, 134217728)
			So(wordsNeeded(math.MaxUint32-31), ShouldEqual, 134217727)
			So(wordsNeeded(math.MaxUint32-63), ShouldEqual, 134217726)
		})

		Convey("Length that a little less to fit full size of words", func() {
			So(wordsNeeded(0), ShouldEqual, 1)
			So(wordsNeeded(31), ShouldEqual, 1)
			So(wordsNeeded(63), ShouldEqual, 2)
			So(wordsNeeded(95), ShouldEqual, 3)
			So(wordsNeeded(math.MaxUint32-32), ShouldEqual, 134217727)
			So(wordsNeeded(math.MaxUint32-64), ShouldEqual, 134217726)
		})

		Convey("Length that a little more to fit full size of words", func() {
			So(wordsNeeded(33), ShouldEqual, 2)
			So(wordsNeeded(65), ShouldEqual, 3)
			So(wordsNeeded(97), ShouldEqual, 4)
		})
	})
}

func TestNew(t *testing.T) {
	Convey("Create new bit sequence with given length", t, func() {
		Convey("Create a bit sequence that has 0 unit", func() {
			bs := New(0)
			So(bs.length, ShouldEqual, 0)
			So(len(bs.words), ShouldEqual, 1)
		})

		Convey("Create a bit sequence that has 1 unit", func() {
			bs := New(1)
			So(bs.length, ShouldEqual, 1)
			So(len(bs.words), ShouldEqual, 1)
		})

		Convey("Create a bit sequence that has 32 units", func() {
			bs := New(32)
			So(bs.length, ShouldEqual, 32)
			So(len(bs.words), ShouldEqual, 1)
		})

		Convey("Create a bit sequence that has 33 units", func() {
			bs := New(33)
			So(bs.length, ShouldEqual, 33)
			So(len(bs.words), ShouldEqual, 2)
		})
	})
}

func TestSet(t *testing.T) {
	Convey("Set unit value of given index according to DiffType", t, func() {
		Convey("Ordered set for 4 different types", func() {
			bs := New(1)
			bs.Set(0, DT_DEFAULT)
			bs.Set(1, DT_SIMPLE)
			bs.Set(2, DT_COMPLEX)
			bs.Set(3, DT_UNKNOWN)
			So(bs.DumpAsBits(), ShouldEqual,
				"0001101100000000000000000000000000000000000000000000000000000000\n")
		})

		Convey("Overwrite set for 4 different types", func() {
			bs := New(1)
			bs.Set(0, DT_DEFAULT)
			bs.Set(0, DT_SIMPLE)
			So(bs.DumpAsBits(), ShouldEqual,
				"0100000000000000000000000000000000000000000000000000000000000000\n")
			bs.Set(0, DT_COMPLEX)
			So(bs.DumpAsBits(), ShouldEqual,
				"1000000000000000000000000000000000000000000000000000000000000000\n")
			bs.Set(0, DT_UNKNOWN)
			So(bs.DumpAsBits(), ShouldEqual,
				"1100000000000000000000000000000000000000000000000000000000000000\n")
			bs.Set(0, DT_DEFAULT)
			So(bs.DumpAsBits(), ShouldEqual,
				"0000000000000000000000000000000000000000000000000000000000000000\n")
		})
	})
}

func TestGet(t *testing.T) {
	Convey("Get unit value by given index", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE)
		So(bs.Get(5), ShouldEqual, DT_SIMPLE)
	})
}

func TestDumpAsBits(t *testing.T) {
	Convey("Convert bit sequence to string format(bits form)", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE)
		bs.Set(8, DT_UNKNOWN)
		bs.Set(10, DT_COMPLEX)
		So(bs.DumpAsBits(), ShouldEqual,
			"0000000000010000110010000000000000000000000000000000000000000000\n")
	})
}

func TestDumpAsType(t *testing.T) {
	Convey("Convert bit sequence to string format(DiffType form)", t, func() {
		bs := New(10)
		bs.Set(5, DT_SIMPLE)
		bs.Set(8, DT_UNKNOWN)
		bs.Set(10, DT_COMPLEX)
		So(bs.DumpAsType(), ShouldEqual,
			"00000100302000000000000000000000\n")
	})
}
