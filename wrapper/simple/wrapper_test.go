package simple

import (
	"fmt"
	"testing"
)

func TestSimpleWrapper(t *testing.T) {
	v := &VeggeMania{}
	c := &CheeseTopping{
		base: v,
	}

	fmt.Println(c.printStructName())
}
