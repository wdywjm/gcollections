package gcollections

import (
	"fmt"
	"testing"
)

func TestIterTool(t *testing.T) {
	it := new(IterTool)

	fmt.Println(it.PermuteInt([]int{1,2,3}))
	fmt.Println(it.PermuteString([]string{"yes", "no"}))

	fmt.Println(it.PermuteIntWithDup([]int{1,1,2,3}))
	fmt.Println(it.PermuteStringWithDup([]string{"yes", "no", "yes"}))

	fmt.Println(it.SubsetsInt([]int{1,2,3}))
	fmt.Println(it.SubsetsString([]string{"yes", "no"}))

	fmt.Println(it.SubsetsIntWithDup([]int{1,1,2,3}))
	fmt.Println(it.SubsetsStingWithDup([]string{"yes", "no", "yes"}))
}
