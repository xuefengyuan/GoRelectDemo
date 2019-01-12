package main

import (
    "reflect"
    "fmt"
)

// 切片逆转
func InvertIntSlice(tmp []int) []int {
    return []int{}
}

// 使用反射进行泛型编程,实现切片逆置
func InvertSlice(args []reflect.Value) (results []reflect.Value) {
    // 获取切片value，以及切片的长度
    inslice, n := args[0], args[0].Len()

    // 创建新的切片,用于保留逆置的结果
    fmt.Println("type = ", inslice.Type())
    outslice := reflect.MakeSlice(inslice.Type(), 0, n)
    // 逆置切片
    for i := n - 1; i >= 0; i-- {
        // 获取切片元素值
        element := inslice.Index(i)
        // 把值追加到新的切片中
        outslice = reflect.Append(outslice, element)
    }
    // 返回结果
    return []reflect.Value{outslice}
}

func Bind(p interface{}, f func(args []reflect.Value) (results []reflect.Value)) {
    // 获取函数变量的reflect.value
    invert := reflect.ValueOf(p).Elem() // 获取函数变量value，目的修改

    // 修改invert的值
    invert.Set(reflect.MakeFunc(invert.Type(), f)) // makefunc 生成方法
}

func main() {
    var invertInt func([]int) []int
    Bind(&invertInt, InvertSlice) // 方法绑定
    fmt.Println(invertInt([]int{1, 2, 3, 4, 5, 6}))

    var invertStr func([]string) []string
    Bind(&invertStr, InvertSlice) // 方法绑定
    str := invertStr([]string{"c", "go", "java", "php"})
    fmt.Println(str)
}
