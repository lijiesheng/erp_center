package lib

import (
	"fmt"
	"testing"
)

func TestQueryMultiRowDemo(t *testing.T) {
	getNavT(Nav)
	for _, v := range NavT {
		fmt.Println(v)
	}
	fmt.Println(len(NavT))
}
