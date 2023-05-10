package visitorfunc

type Info struct {
	a int
}

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func(*Info, error) error

type Test1Info struct {
	a int
}

func (T *Test1Info) accept(visitor Visitor) error {
	return visitor.Visit(func(info *Info, err error) error {
		return nil
	})
}
