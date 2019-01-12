# GoRelectDemo

Go 反射示例Demo



## Go的反射

> reflect.TypeOf()   ：获取类型
> reflect.ValueOf()  ：获取值与修改值



## 一、静态类型

通过静态类型创建reflet.Type类型，主要用于描述变量的类型

### 1、int类型

```go
var a = 10
// 通过静态类型创建reflet.Type类型，主要用于描述变量的类型
at := reflect.TypeOf(a) 
fmt.Println(at)

var ai interface{}
ai = a
typ1 := reflect.TypeOf(ai)
fmt.Println(typ1)
```

获取值与修改值

```go
// 获取 值 与修改值
av:=reflect.ValueOf(a)  // valueof 同静态类型变量获取动态类型值的
fmt.Println(av)
// 检查value是否能被修改
fmt.Println(av.CanSet())  

// 要想通过反射修改变量本身的值要传递指针，和方法修改变量一个道理
av1:=reflect.ValueOf(&a)
fmt.Println(av1)
fmt.Println(av1.Elem(),av1.Elem().Type())
fmt.Println(av1.CanSet(),av1.Elem().CanSet())  // false true

av1.Elem().Set(reflect.ValueOf(1)) // 通过反射修改值
fmt.Println(a)  // 原始值被修改了

```

### 2、string类型

```go
// string 类型检查
str := "hello"
st := reflect.TypeOf(str)
fmt.Println(st)
```

获取值与修改值

```go
st1 := reflect.ValueOf(&str)
fmt.Println(st1.Elem().CanSet()) // 判断string值是否可以修改
// 修改值需要通过xxx.Elem.set方法调用
st1.Elem().Set(reflect.ValueOf("world"))
// 获取值
fmt.Println(st1.Elem().Interface())
fmt.Println(str)

```

## 二、切片类型

```go
// 切片类型
sl := []int{1, 2, 3, 4, 5, 6}
slt := reflect.TypeOf(sl)
fmt.Println(slt)
// 通过反射获取容器类型中保存的数据的类型 Elem()
fmt.Println(slt.Elem()) // 检查切片容器中的数据类型
```

获取值与修改值

```go
slv:=reflect.ValueOf(sl)
fmt.Println(slv.CanSet())
// 修改slv切片value的值，可以修改的原因同切片传递参数修改一样的
slv.Index(1).Set(reflect.ValueOf(100)) 
fmt.Println(sl)

// 通过指针传参
slv1:=reflect.ValueOf(&sl)
fmt.Println(slv1.CanSet())
// 通过反射访问切片的值
fmt.Println("index = ",slv1.Elem().Index(3))
slv1.Elem().Index(1).Set(reflect.ValueOf(300))
fmt.Println(sl)
slv1.Elem().Set(reflect.ValueOf([]int{5,6,7,8,9,100,200}))
fmt.Println(sl)
```

## 三、Map类型

```go
// map 类型检查
mp:=map[int]string{
   1:"golang",
   2:"C++",
}
mpt:=reflect.TypeOf(mp)
fmt.Println(mpt,mpt.Kind(),mpt.Name())
```

获取与修改值

```go
// 修改值
mpt1 := reflect.ValueOf(mp)
key := reflect.ValueOf(3)
value := reflect.ValueOf("java")
mpt1.SetMapIndex(key,value)
// 获取值 
fmt.Println(mpt1.MapIndex(key).Interface())
fmt.Println(mp)

// 利用反射重新创建map
mapType := reflect.TypeOf(mpt)
// 创建新值
mapReflect := reflect.MakeMap(mapType)
// 使用新创建的变量
key2 := reflect.ValueOf(4)
value1 := reflect.ValueOf("php")

mapReflect.SetMapIndex(key2, value1)
mp2 := mapReflect.Interface().(map[int]string)
fmt.Println(mp2)

```



## 四、结构体

### 1、反射操作结构体一

```go
type User struct {
	Name string  `orm:"hello" validate:"world"`
	Age int
}

func (tmp User)MyPrint()  {
	fmt.Println(tmp.Name)
}
```

反射获取结构体信息

```go
// 查看结构体中的字段信息
fmt.Println(usert.NumField())  // 结构体字段的个数
// 检查字段名字
for i:=0;i<usert.NumField();i++{
	curF:=usert.Field(i)  // 获取当前字段的类型信息
	fmt.Println(curF.Name,curF.Type)  // 获取当前字段的名字与类型
	fmt.Println(curF.Tag)  // 获取tag信息
	tg1:=curF.Tag.Get("orm")  // 获取指定名字的标签的内容
	tg2:=curF.Tag.Get("validate")
	fmt.Println(tg1,tg2)
}
// 通过反射检查结构体的方法
fmt.Println(usert.NumMethod())  // 检查结构体有几个方法
fnt:=usert.Method(0)
fmt.Println(fnt.Name,fnt.Type) //
```

通过反射修改结构体信息和调用结构体方法

```go
userv := reflect.ValueOf(&user)
for i := 0; i < userv.Elem().NumField(); i++ {
    curFV := userv.Elem().Field(i)
    fmt.Println(curFV)
    if curFV.Type().Kind() == reflect.String {
         curFV.Set(reflect.ValueOf("MrWang"))
     } else {
         curFV.Set(reflect.ValueOf(50))
     }

}
fmt.Println(user)

// 通过反射去调用结构体方法
fn:=userv.Elem().Method(0)  //获取结构体的方法
fmt.Println(fn)
fn.Call([]reflect.Value{})  // 调用函数
```

### 2、反射操作结构体二

```go
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
```

获取结构体类型，方法、修改值，执行方法

```go
func main() {
    // 如果valueoof传递是实体参数，那么我只能获取绑定的实体的类型的可以导出的方法
    // 如果传递的是结构体指针, 你可以调用所有的可以导出的方法。
    stu:=pck.Student{Name:"MrLi",Age:20}
    stu.SetW(1000)
    fmt.Println(stu)

    stuv:=reflect.ValueOf(stu)
    fmt.Println(stuv.NumField()) // 检查结构体字段个数
    fmt.Println("=========")
    for i:=0;i<stuv.NumField();i++{
        curFV:=stuv.Field(i)
        fmt.Println(curFV)
    }

    // 获取方法的时候
    // 获取方法个数，传递是实体时，只能获取绑定的是实体类型的方而且是可以导出的
    fmt.Println(stuv.NumMethod())  
   
    stuv.Method(0).Call([]reflect.Value{})

    stupv:=reflect.ValueOf(&stu)
    // 获取方法个数，传递的是结构体指针, 可以调用所有的可以导出的方法。
    fmt.Println(stupv.NumMethod())  //输出是2
    stfn:=stupv.Method(0)
    fmt.Println(stfn)
    stfn.Call([]reflect.Value{})
    stfn1:=stupv.Method(1)

    stfn1.Call([]reflect.Value{reflect.ValueOf(100)})
    fmt.Println(stu)
}
```

## 五、示例：切片逆置

```go
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
```

