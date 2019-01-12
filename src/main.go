package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Name string `orm:"hello" validate:"world"`
    Age  int
}

func (tmp User) MyPrint() {
    fmt.Println(tmp.Name)
}

func main() {

    // map 类型检查
    mp:=map[int]string{
        1:"golang",
        2:"C++",
    }
    mpt:=reflect.TypeOf(mp)
    fmt.Println(mpt,mpt.Kind(),mpt.Name())

    mpt1 := reflect.ValueOf(mp)
    key := reflect.ValueOf(3)
    value := reflect.ValueOf("java")
    mpt1.SetMapIndex(key,value)

    fmt.Println(mpt1.MapIndex(key).Interface())
    fmt.Println(mp)

    // striing 类型检查
    str := "hello"
    st := reflect.TypeOf(str)
    fmt.Println(st)

    st1 := reflect.ValueOf(&str)
    fmt.Println(st1.Elem().CanSet()) // 判断string值是否可以修改
    // 修改值需要通过xxx.Elem.set方法调用
    st1.Elem().Set(reflect.ValueOf("world"))
    fmt.Println("=====",st1.Elem().Interface())

    fmt.Println(str)
}

