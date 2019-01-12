package pck

import "fmt"

type Student struct {
    Name string
    Age int
    weight int  // 小写
}

func (tmp Student)MPrint()  {
    fmt.Println(tmp.Name)
}

func (this *Student)SetW(num int)  {
    this.weight = num
}

func (this Student)show()  {
    fmt.Println(this.weight)
}