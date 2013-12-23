package bits

// Combination represents a combination of any two tails.
type Combination struct {
	Nums [2]int
	Same bool
}

// CombinationTable contains common possible combinations of two tails,
// and uses the index 0 for marking exceptions.
var CombinationTable = [16]Combination{
	1:  Combination{[2]int{1, 1}, true},
	2:  Combination{[2]int{1, 2}, false},
	3:  Combination{[2]int{2, 1}, false},
	4:  Combination{[2]int{2, 2}, true},
	5:  Combination{[2]int{1, 3}, false},
	6:  Combination{[2]int{2, 3}, false},
	7:  Combination{[2]int{3, 1}, false},
	8:  Combination{[2]int{3, 2}, false},
	9:  Combination{[2]int{3, 3}, true},
	10: Combination{[2]int{1, 4}, false},
	11: Combination{[2]int{1, 5}, false},
	12: Combination{[2]int{1, 6}, false},
	13: Combination{[2]int{2, 5}, false},
	14: Combination{[2]int{3, 5}, false},
	15: Combination{[2]int{4, 5}, false},
}

// Length to be 16 is because it can be represented by 4 bits.
const CombineTableLength = 16

// GetCombineTableIndex returns index of combinations in CombinationTable,
// and it returns 0 when it does not match any.
func GetCombineTableIndex(num1, num2 int) int {
	for i := 1; i < CombineTableLength; i++ {
		if num1 == CombinationTable[i].Nums[0] &&
			num2 == CombinationTable[i].Nums[1] {
			return i
		}
	}
	return 0
}
