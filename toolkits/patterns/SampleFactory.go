package patterns

import "fmt"

type SampleFactory interface {
	Product()
}

//ProdCup 
type ProdCup struct {
	Name string
	Pin  string
}

func (p *ProdCup) Product() {
	fmt.Println("start make cup")
	p.Name = ""
}

//ProdBox 
type ProdBox struct {
	Size     int
	Name     string
	ColorPin string
}

func (p ProdBox) Product() {
	panic("implement me")
}


