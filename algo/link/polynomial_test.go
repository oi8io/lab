package link

import (
	"fmt"
	"testing"
)

var poly *Polynomial

func init() {
	d := map[int]int{
		10: 1000,
		5:  14,
		1:  0,
	}
	fmt.Println(d)
	//poly = NewPolynomial(d)
	//fmt.Println(poly)
}
func TestNewPolynomial(t *testing.T) {
	d := map[int]int{
		1: 1,
		2: 2,
		5: 5,
		3: 3,
		4: 4,
		0: 10,
	}

	p := NewPolynomial(d)

	d2 := map[int]int{
		1: 1,
		2: 2,
		5: 5,
		3: 3,
		4: 4,
		0: 11,
	}
	p2 := NewPolynomial(d2)
	p.Print()
	p = p.Merge(p2)
	p2.Print()
	p.Print()
}
