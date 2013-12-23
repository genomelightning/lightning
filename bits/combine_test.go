package bits

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_init(t *testing.T) {

}

func TestGetCombineTableIndex(t *testing.T) {
	Convey("Get index of combination in CombinationTable", t, func() {
		Convey("Combination that does exist", func() {
			So(GetCombineTableIndex(3, 1), ShouldEqual, 7)
		})

		Convey("Combination that does not exist", func() {
			So(GetCombineTableIndex(3, 4), ShouldEqual, 0)
		})
	})
}
