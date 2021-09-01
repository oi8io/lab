package link

import "fmt"

//Polynomial 多项式
type Polynomial struct {
	Coefficient int //系数
	Exponent    int // 最高指数
	Next        *Polynomial
}

func NewPolynomial(data map[int]int) *Polynomial {
	var p *Polynomial
	for k, v := range data {
		p = p.Insert(k, v)
		fmt.Println(p)
	}
	return p
}

//Insert
// 指数
// 系数
func (p *Polynomial) Insert(exponent, coefficient int) *Polynomial {
	var i = new(Polynomial)
	i.Exponent = exponent
	i.Coefficient = coefficient
	if p == nil {
		p = i
		return p
	}
	if i.Exponent > p.Exponent {
		i.Next = p
		p = i
		return p
	}
	if p.Exponent == i.Exponent {
		p.Coefficient = p.Coefficient + i.Coefficient
		return p
	}

	p.Next = p.Next.Insert(exponent, coefficient)
	return p
}
func (p *Polynomial) Merge(p1 *Polynomial) *Polynomial {
	for {
		p.Insert(p1.Exponent, p1.Coefficient)
		if p1.Next == nil {
			break
		} else {
			p1 = p1.Next
		}
	}
	return p
}

func (p *Polynomial) Print() (str []string) {
	for {
		str = append(str, fmt.Sprintf("%dX^%d", p.Coefficient, p.Exponent))
		if p.Next == nil {
			break
		} else {
			p = p.Next
		}
	}
	fmt.Println(str)
	return str
}
