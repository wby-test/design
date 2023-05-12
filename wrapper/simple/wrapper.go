package simple

import "fmt"

// Base 零件接口
type Base interface {
	printStructName() string
}

// VeggeMania 具体零件
type VeggeMania struct {
}

func (v *VeggeMania) printStructName() string {
	fmt.Println("VeggeMania")
	return "VeggeMania"
}

// CheeseTopping 具体装饰
type CheeseTopping struct {
	base Base
}

func (c *CheeseTopping) printStructName() string {
	str := c.base.printStructName()
	fmt.Println("c *CheeseTopping  ", str)
	return "CheeseTopping " + str
}
