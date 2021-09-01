package patterns

import "fmt"

type Library struct {
	NeedChange
}

type NeedChange interface {
	Setup2()
	Setup4()
}

type LibraryWrap struct {

}

func (l *Library) Setup1()  {
	fmt.Println("setup1")
}


func (l *Library) Setup3()  {
	fmt.Println("setup3")
}



func (l *Library) Setup5()  {
	fmt.Println("setup5")
}

func (l *Library)Run()  {
	l.Setup1()
	l.Setup2()
	l.Setup3()
	l.Setup4()
	l.Setup5()
}

