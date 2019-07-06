# goGuide
go语言入门学习


## 编程基础

<details>
<summary>内置关键字</summary>

```go
break      default       func      interface    select
case       defer         go        map          struct
chan       else          goto      package      switch
const      fallthrough   if        range        type
continue   for           import    return       var
```

</details>

<details>
<summary>预定义标识符</summary>

```go
append  bool    byte    cap     close  complex complex64 complex128 uint16
copy    false   float32 float64 imag   int     int8      int16      uint32
int32   int64   iota    len     make   new     nil       panic      uint64
print   println real    recover string true    uint      uint8      uintptr
```

</details>

<details>
<summary>行分隔符</summary>

- 在 Go 程序中，一行代表一个语句结束，不需要分隔符。
- 打算将多个语句写在同一行，它们则必须使用 `;` 人为区分，并不鼓励这种做法。

</details>

<details>
<summary>注释方法</summary>

```go
// 单行注释

/*
  多行注释
*/
```

</details>

<details>
<summary>标识符</summary>

- 标识符用来命名变量、类型等程序实体。
- 第一个字符必须是字母或下划线而不能是数字

有效标识符

```
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal
```

无效标识符

```bash
1ab  #（以数字开头）
case #（Go 语言的关键字）
a+b  #（运算符是不允许的）
```

</details>

<details>
<summary>包引用 import</summary>

```go
import "fmt"
import "os"
import "io"
```

简写方式如下

```go
import (
  "fmt"
  "os"
  "io"
)
```

**包引用介绍**

```bash
.
├── cal
│   ├── add.go
│   ├── multi
│   │   └── multiply.go
│   └── subtract.go
└── main.go
```

注意：[package-demo](example/package-demo) 文件夹复制到 `$GOPATH/src/` 目录下，不然运行报错哦

```bash
go run $GOPATH/src/package-demo/main.go
```

main.go中如何调用add.go、subtract.go或者是multiply.go文件中的函数。

> `add.go`和`subtract.go`文件中，包名为cal `package cal`  
> `multiply.go`在 multi 文件夹下，所以程序的包名为multi `package multi`  
> 如果 mian 函数要调用`add.go`或者`subtract.go`中的函数，必须要引入包"cal" `import "package-demo/cal"`  
> 要调用`multiply.go`中的函数，必须要引入包"multi"，`import "package-demo/cal/multi"`  
> Go中如果函数名的首字母大写，表示该函数是公有的，可以被其他程序调用，如果首字母小写，该函数就是是私有的

</details>

<details>
<summary>包别名</summary>

```go
import(
  ff "fmt"
)

// 或者
import ff "fmt"

// 别名包调用
ff.Println('Hello World!')
```

</details>


<details>
<summary>省略调用</summary>

```go
import(
  . "fmt"
)
func main() {
  // 省略调用
  Println('Hello World!')
}
```

</details>


<details>
<summary>可见性规则</summary>

Go语言中约定使用 **大小写** 来决定常量、变量、类型、接口、结构或函数是否可以被外部包所调用

- 函数名字首字母 **小写** 即为 `private` 私有的
- 函数名字首字母 **大写** 即为 `public` 公有

</details>

## 基本类型

<details>
<summary>布尔型</summary>

```go
var b bool
b  = true
fmt.Printf("b is of type %t\n", b)
e := bool(true)
fmt.Printf("e is of type %t\n", e)
```

- 长度：1字节
- 取值范围：true/false
- 只能使用true/false值，不能使用数字代替布尔值

</details>

<details>
<summary>整形 int/uint</summary>

```go
package main
import "fmt"
func main() {
  // n 是一个长度为 10 的数组
  var n [10]int 
  var i,j int

  /* 为数组 n 初始化元素 */         
  for i = 0; i < 10; i++ {
    n[i] = i + 100 /* 设置元素为 i + 100 */
  }

  /* 输出每个数组元素的值 */
  for j = 0; j < 10; j++ {
    fmt.Printf("Element[%d] = %d\n", j, n[j] )
  }
}
```

- int/uint
- 根据平台可能为32/64位

</details>

<details>
<summary>8位整型 int8/uint8</summary>

```go
u8 := []uint8{98, 99}

a := byte(255)  //11111111 这是byte的极限， 因为 a := byte(256)//越界报错， 0~255正好256个数，不能再高了
b := uint8(255) //11111111 这是uint8的极限，因为 c := uint8(256)//越界报错，0~255正好256个数，不能再高了
c := int8(127)  //01111111 这是int8的极限， 因为 b := int8(128)//越界报错， 0~127正好128个数，所以int8的极限只是256的一半
d := int8(a)    //11111111 打印出来则是-0000001，int8(128)、int8(255)、int8(byte(255))都报错越界，因为int极限是127，但是却可以写：int8(a)，第一位拿来当符号了
e := int8(c)    //01111111 打印出来还是01111111
fmt.Printf("%08b %d \n", a, a)
fmt.Printf("%08b %d \n", b, b)
fmt.Printf("%08b %d \n", c, c)
fmt.Printf("%08b %d \n", d, d)
fmt.Printf("%08b %d \n", e, e)
```

- int8/uint8
- 长度：1字节
- 取值范围：-128~127/0~255

</details>

<details>
<summary>字节型 byte</summary>

```go
// 这里不能写成 b := []byte{"Golang"}，这里是利用类型转换。
b := []byte("Golang")
c := []byte("go")
d := []byte("Go")
println(b,c,d)
```

- byte(uint8别名)

**基本处理函数**

- `Contains()` 返回是否包含子切片
- `Count()` 子切片非重叠实例的数量
- `Map()` 函数，将byte 转化为Unicode，然后进行替换
- `Repeat()` 将切片复制count此，返回这个新的切片
- `Replace()` 将切片中的一部分 替换为另外的一部分
- `Runes()` 将 S 转化为对应的 UTF-8 编码的字节序列，并且返回对应的Unicode 切片
- `Join()` 函数，将子字节切片连接到一起。

可以参考下面列子来理解上面7个方法，例子 [byte.go](./example/byte/byte.go)

```go
package main
import (
  "bytes"
  "fmt"
)
func main() {
  // 这里不能写成 b := []byte{"Golang"}，这里是利用类型转换。
  b := []byte("Golang")
  subslice1 := []byte("go")
  subslice2 := []byte("Go")
  // func Contains(b, subslice [] byte) bool
  // 检查字节切片b ，是否包含子字节切片 subslice
  fmt.Println(bytes.Contains(b, subslice1))
  fmt.Println(bytes.Contains(b, subslice2))


  s2 := []byte("同学们，上午好")
  m := func(r rune) rune {
    if r == '上' {
      r = '下'
    }
    return r
  }
  fmt.Println(string(s2))
  // func Map(mapping func(r rune) rune, s []byte) []byte
  // Map函数: 首先将 s 转化为 UTF-8编码的字符序列，
  // 然后使用 mapping 将每个Unicode字符映射为对应的字符，
  // 最后将结果保存在一个新的字节切片中。
  fmt.Println(string(bytes.Map(m, s2)))


  s3 := []byte("google")
  old := []byte("o")
  //这里 new 是一个字节切片，不是关键字了
  new := []byte("oo")
  n := 1
  // func Replace(s, old, new []byte, n int) []byte
  //返回字节切片 S 的一个副本， 并且将前n个不重叠的子切片 old 替换为 new，如果n < 0 那么不限制替换的数量
  fmt.Println(string(bytes.Replace(s3, old, new, n)))
  fmt.Println(string(bytes.Replace(s3, old, new, -1)))


  // 将字节切片 转化为对应的 UTF-8编码的字节序列，并且返回对应的 Unicode 切片。
  s4 := []byte("中华人民共和国")
  r1 := bytes.Runes(s4)
  // func Runes(b []byte) []rune
  fmt.Println(string(s4), len(s4))  // 字节切片的长度
  fmt.Println(string(r1), len(r1))  // rune 切片的长度


  // 字节切片 的每个元素，依旧是字节切片。
  s5 := [][]byte{
    []byte("你好"),
    []byte("世界"),  //这里的逗号，必不可少
  }
  sep := []byte(",")
  // func Join(s [][]byte, sep []byte) []byte
  // 用字节切片 sep 吧 s中的每个字节切片连接成一个，并且返回.
  fmt.Println(string(bytes.Join(s5, sep)))
}
```

</details>

<details>
<summary>16位整型 int16/uint16</summary>

- int16/uint16
- 长度：2字节
- 取值范围：-32768~32767/0~65535

</details>

<details>
<summary>32位整型 int32(别名rune)/uint32</summary>

- int32(别名rune)/uint32
- 长度：4字节
- 取值范围：-2^32/2~2^32/2-1/0~2^32-1

</details>

<details>
<summary>64位整型 int64/uint64</summary>

- int64/uint64
- 长度：8字节
- 取值范围：-2^64/2~2^64/2-1/0~2^64-1

</details>

<details>
<summary>浮点型 float32/float64</summary>

```go
package main
import "fmt"

func main() {
  var x float64
  x = 20.0
  fmt.Println(x)
  fmt.Printf("x is of type %T\n", x)

  a := float64(20.0)
  b := 42 
  fmt.Println(a)
  fmt.Println(b)
  fmt.Printf("a is of type %T\n", a)
  fmt.Printf("b is of type %T\n", b)
}
```

实例：[float.go](./example/float/float.go)

- float32/float64
- 长度：4/8字节
- 小数位：精确到 7/15 位小数

</details>

<details>
<summary>复数 complex64/complex128</summary>

- complex64/complex128
- 长度：8/16

</details>

<details>
<summary>指针整数型 uintptr</summary>

用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。

- uintptr
- 保存指正的 32 位或者 64 位整数型

```go
// 示例：通过指针修改结构体字段
package main
import (
  "fmt"
  "unsafe"
)

func main() {
  s := struct {
    a byte
    b byte
    c byte
    d int64
  }{0, 0, 0, 0}

  // 将结构体指针转换为通用指针
  p := unsafe.Pointer(&s)
  // 保存结构体的地址备用（偏移量为 0）
  up0 := uintptr(p)
  // 将通用指针转换为 byte 型指针
  pb := (*byte)(p)
  // 给转换后的指针赋值
  *pb = 10
  // 结构体内容跟着改变
  fmt.Println(s)

  // 偏移到第 2 个字段
  up := up0 + unsafe.Offsetof(s.b)
  // 将偏移后的地址转换为通用指针
  p = unsafe.Pointer(up)
  // 将通用指针转换为 byte 型指针
  pb = (*byte)(p)
  // 给转换后的指针赋值
  *pb = 20
  // 结构体内容跟着改变
  fmt.Println(s)

  // 偏移到第 3 个字段
  up = up0 + unsafe.Offsetof(s.c)
  // 将偏移后的地址转换为通用指针
  p = unsafe.Pointer(up)
  // 将通用指针转换为 byte 型指针
  pb = (*byte)(p)
  // 给转换后的指针赋值
  *pb = 30
  // 结构体内容跟着改变
  fmt.Println(s)

  // 偏移到第 4 个字段
  up = up0 + unsafe.Offsetof(s.d)
  // 将偏移后的地址转换为通用指针
  p = unsafe.Pointer(up)
  // 将通用指针转换为 int64 型指针
  pi := (*int64)(p)
  // 给转换后的指针赋值
  *pi = 40
  // 结构体内容跟着改变
  fmt.Println(s)
}
```

</details>

<details>
<summary>数组类型 array</summary>

数组声明语法

```go
var variable_name [SIZE]variable_type
```

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。下面是一个简单的对数组操作的例子[array.go](./example/array/array.go)

```go
package main
import "fmt"
func main() {
  // 声明一个长度为5的整数数组
  // 一旦数组被声明了，那么它的数据类型跟长度都不能再被改变。
  var array1 [5]int
  
  fmt.Printf("array1: %d\n\n", array1)

  // 声明一个长度为5的整数数组
  // 初始化每个元素
  array2 := [5]int{12, 123, 1234, 12345, 123456}
  array2[1] = 5000
  fmt.Printf("array2: %d\n\n", array2[1])
  
  // n 是一个长度为 10 的数组
  var n [10]int 
  var i,j int

  /* 为数组 n 初始化元素 */         
  for i = 0; i < 10; i++ {
    n[i] = i + 100 /* 设置元素为 i + 100 */
  }

  /* 输出每个数组元素的值 */
  for j = 0; j < 10; j++ {
    fmt.Printf("Element[%d] = %d\n", j, n[j] )
  }

  /* 数组 - 5 行 2 列*/
  var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
  var e, f int

  /* 输出数组元素 */
  for  e = 0; e < 5; e++ {
    for f = 0; f < 2; f++ {
        fmt.Printf("a[%d][%d] = %d\n", e,f, a[e][f] )
    }
  }
}
```

初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

```go
var array1 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

```go
float32 salary = array1[9]
```

以上实例读取了数组`array1`第`10`个元素的值。

多维数组，下面例子

```go
// 三行四列
a = [3][4]int{  
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11}   /*  第三行索引为 2 */
}
```

访问多维数组

```go
// 访问第2行第3列
int val = a[2][3]
```

</details>

<details>
<summary>结构类型 struct</summary>

```go
package main
import "fmt"
type Vertex struct {
	X int
	Y int
}
func main() {
  fmt.Println(Vertex{1, 2})
  
  // 结构体字段使用点号来访问。
	v := Vertex{1, 2}
	v.X = 4
  fmt.Println(v.X)
  
  // 结构体字段可以通过结构体指针来访问。
	e := Vertex{1, 2}
	p := &e
	p.X = 1e9
  fmt.Println(e)
  
  
  var (
    v1 = Vertex{1, 2}  // 类型为 Vertex
    v2 = Vertex{X: 1}  // Y:0 被省略
    v3 = Vertex{}      // X:0 和 Y:0
    p  = &Vertex{1, 2} // 类型为 *Vertex , 特殊的前缀 & 返回一个指向结构体的指针
  )
  fmt.Println(v1, p, v2, v3)
}
```

简单的结构体

```go
type T struct {a, b int}
```

结构体里的字段都有 名字，像 `field1`、`field2` 等，如果字段在代码中从来也不会被用到，那么可以命名它为 `_`。上面简单的结构体定义，下面调用方法：

```go
var s T
s.a = 5
s.b = 8
```

数组可以看作是一种结构体类型，不过它使用下标而不是具名的字段。

```go
var t *T
t = new(T)
```

上面简单的管用语句方法`t := new(T)`，变量 `t` 是一个指向 `T` 的指针，此时结构体字段的值是它们所属类型的零值。

声明 `var t T` 也会给 `t` 分配内存，并零值化内存，但是这个时候 `t` 是类型`T`。在这两种方式中，`t` 通常被称做类型 `T` 的一个实例（instance）或对象（object）。

一个非常简单的例子[structs_fields.go](./example/structs/structs_fields.go)运行例子查看结果：

```bash
→ go run test/structs_fields.go

The int is: 10
The float is: 15.500000
The string is: Chris
&{10 15.5 Chris}
```

**使用 new**

</details>

<details>
<summary>字符串类型 string</summary>

```go
var str string //声明一个字符串
str = "Go lang"  //赋值
ch :=str[0]    //获取第一个字符
len :=len(str) //字符串的长度,len是内置函数 ,len=5
```

len函数是Go中内置函数，不引入strings包即可使用。len(string)返回的是字符串的字节数。len函数所支持的入参类型如下：

- len(Array) 数组的元素个数
- len(*Array) 数组指针中的元素个数,如果入参为nil则返回0
- len(Slice) 数组切片中元素个数,如果入参为nil则返回0
- len(map) 字典中元素个数,如果入参为nil则返回0
- len(Channel) Channel buffer队列中元素个数

</details>

<details>
<summary>接口类型 inteface</summary>

```go
package main
import (
   "fmt"
   "math"
)

/* 定义一个 interface */
type shape interface {
   area() float64
}

/* 定义一个 circle */
type circle struct {
   x,y,radius float64
}

/* 定义一个 rectangle */
type rectangle struct {
   width, height float64
}

/* 定义一个circle方法 (实现 shape.area())*/
func(circle circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

/* 定义一个rectangle方法 (实现 shape.area())*/
func(rect rectangle) area() float64 {
   return rect.width * rect.height
}

/* 定义一个shape的方法*/
func getArea(shape shape) float64 {
   return shape.area()
}

func main() {
   circle := circle{x:0,y:0,radius:5}
   rectangle := rectangle {width:10, height:5}

   fmt.Printf("circle area: %f\n",getArea(circle))
   fmt.Printf("rectangle area: %f\n",getArea(rectangle))
}
```

实例：[inteface.go](./example/inteface/inteface.go)

</details>

<details>
<summary>函数类型 func</summary>

```go
package main
import "fmt"
type functinTyoe func(int, int) // 声明了一个函数类型
func (f functinTyoe)Serve() {
  fmt.Println("serve2")
}
func serve(int,int) {
  fmt.Println("serve1")
}
func main() {
  a := functinTyoe(serve)
  a(1,2)
  a.Serve()
}
```

实例：[func.go](./example/func/func.go)

</details>

<details>
<summary>引用类型 func</summary>


**切片**

> 是一种可以动态数组，可以按我们的希望增长和收缩。

- slice

**Map**

> 是一种无序的键值对的集合。是一种集合，所以我们可以像迭代数组和 slice 那样迭代它。

- map

```go
// 通过 make 来创建
dict := make(map[string]int)
// 通过字面值创建
dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

// 给 map 赋值就是指定合法类型的键，然后把值赋给键
colors := map[string]string{}
colors["Red"] = "#da1337"

// 不初始化 map , 就会创建一个 nil map。nil map 不能用来存放键值对，否则会报运行时错误
var colors map[string]string
colors["Red"] = "#da1337"
// Runtime Error:
// panic: runtime error: assignment to entry in nil map

//选择是只返回值，然后判断是否是零值来确定键是否存在。
value := colors["Blue"]
if value != "" {
  fmt.Println(value)
}
```

在函数间传递 map 不是传递 map 的拷贝。所以如果我们在函数中改变了 map，那么所有引用 map 的地方都会改变

```go
func main() {
  colors := map[string]string{
     "AliceBlue":   "#f0f8ff",
     "Coral":       "#ff7F50",
     "DarkGray":    "#a9a9a9",
     "ForestGreen": "#228b22",
  }
  for key, value := range colors {
      fmt.Printf("Key: %s  Value: %s\n", key, value)
  }
  removeColor(colors, "Coral")
  for key, value := range colors {
      fmt.Printf("Key: %s  Value: %s\n", key, value)
  }
}
func removeColor(colors map[string]string, key string) {
    delete(colors, key)
}
```


**通道**

- chan

</details>

<details>
<summary>类型别名</summary>

```go
type (
  byte int8
  rune init32
  文本 string
)
var b 文本
b = "别名类型，可以是中文！"
```

</details>

## 常量变量

<details>
<summary>常量 const</summary>

```go
package main
import "unsafe"
// 常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。
// 常量表达式中，函数必须是内置函数，否则编译不过：
const (
  a = "abc"
  b = len(a)
  c = unsafe.Sizeof(a)
)

func main(){
  const (
    PI     = 3.14
    const1 = "1"
  )
  const LENGTH int = 10
  const e, f, g = 1, false, "str" //多重赋值
  println(a, b, c,PI, LENGTH)
}
```

上面例子[const.go](./example/const/const.go)

**iota** 特殊常量，可以认为是一个可以被编译器修改的常量。[iota.go](./example/iota/iota.go)

```go
package main
import "fmt"
func main() {
  const (
    // 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
    // 所以 a=0, b=1, c=2 可以简写为如下形式：
    a = iota   //0
    b          //1
    c          //2
    d = "ha"   //独立值，iota += 1
    e          //"ha"   iota += 1
    f = 100    //iota +=1
    g          //100  iota +=1
    h = iota   //7,恢复计数
    i          //8
  )
  fmt.Println(a,b,c,d,e,f,g,h,i)
}
```

</details>

<details>
<summary>变量 var</summary>

```go
var (
  name  = "gopher"
  name1 = "1"
)
// 变量声明
var a int
a = 11 /* 赋值 */

// 变量声明 并赋值
var b int = 12

// 应用在函数体内的方式
var a, b, c, d int = 1, 2, 3, 4
// a =1 
// b =2 
// c =3 
// d =4 


var a, _, c, d int = 1, 2, 3, 4
// 忽略 _ 返回值忽略
```

- 全局变量名 以大写开头
- 全局变量不可以省略 var ，可以使用并行的方式
- 所有变量都可以使用类型推断
- 局部变量不可以使用`var()`简写的形式

</details>

<details>
<summary>变量的类型转换</summary>

```go
// 只能类型显式转换
var a float32 = 1.1
// 省略var, 简短形式，使用 := 赋值操作符
b := int(a)
// 不兼容的类型不能转换类型
```

</details>

<details>
<summary>多变量声明</summary>

```go
var x, y int
// 这种因式分解关键字的写法一般用于声明全局变量
var (
  a int
  b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
```

</details>

## 语言运算符

<details>
<summary>算术运算符</summary>

```go
package main
import "fmt"
func main() {
  var a int = 21
  var b int = 10
  var c int
  c = a + b
  fmt.Printf("第一行 - c 的值为 %d\n", c ) // 第一行 - c 的值为 31
  c = a - b
  fmt.Printf("第二行 - c 的值为 %d\n", c ) // 第二行 - c 的值为 11
  c = a * b
  fmt.Printf("第三行 - c 的值为 %d\n", c ) // 第三行 - c 的值为 210
  c = a / b
  fmt.Printf("第四行 - c 的值为 %d\n", c ) // 第四行 - c 的值为 2
  c = a % b
  fmt.Printf("第五行 - c 的值为 %d\n", c ) // 第五行 - c 的值为 1
  a++
  fmt.Printf("第六行 - c 的值为 %d\n", a ) // 第六行 - c 的值为 22
  a--
  fmt.Printf("第七行 - c 的值为 %d\n", a ) // 第七行 - c 的值为 21
}
```

下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| + | 相加  | A + B 输出结果 30 |
| - | 相减  | A - B 输出结果 -10 |
| * | 相乘  | A * B 输出结果 200 |
| / | 相除  | B / A 输出结果 2 |
| % | 求余  | B % A 输出结果 0 |
| ++ | 自增 | A++ 输出结果 11 |
| -- | 自减 | A-- 输出结果 9 |

</details>

<details>
<summary>关系运算符</summary>

```go
package main
import "fmt"
func main() {
   var a int = 21
   var b int = 10
   if( a == b ) {
      fmt.Printf("第一行 - a 等于 b\n" )
   } else {
      fmt.Printf("第一行 - a 不等于 b\n" )
   }
   if ( a < b ) {
      fmt.Printf("第二行 - a 小于 b\n" )
   } else {
      fmt.Printf("第二行 - a 不小于 b\n" )
   } 
   
   if ( a > b ) {
      fmt.Printf("第三行 - a 大于 b\n" )
   } else {
      fmt.Printf("第三行 - a 不大于 b\n" )
   }
   /* 让我们改变a和b的值 */
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("第四行 - a 小于等于 b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("第五行 - b 大于等于 a\n" )
   }
}
```

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| ==  | 检查两个值是否相等，如果相等返回 True 否则返回 False。 |  (A == B) 为 False |
| !=  | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。 |  (A != B) 为 True |
| > | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。 |  (A > B) 为 False |
| < | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。 |  (A < B) 为 True |
| >=  | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 |  (A >= B) 为 False |
| <=  | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | A <= B) 为 True |

</details>

<details>
<summary>逻辑运算符</summary>

```go
package main
import "fmt"
func main() {
  var a bool = true
  var b bool = false
  if ( a && b ) {
    fmt.Printf("第一行 - 条件为 true\n" )
  }
  if ( a || b ) {
    fmt.Printf("第二行 - 条件为 true\n" )
  }
  /* 修改 a 和 b 的值 */
  a = false
  b = true
  if ( a && b ) {
    fmt.Printf("第三行 - 条件为 true\n" )
  } else {
    fmt.Printf("第三行 - 条件为 false\n" )
  }
  if ( !(a && b) ) {
    fmt.Printf("第四行 - 条件为 true\n" )
  }
}
```

下表列出了所有Go语言的逻辑运算符。假定 A 值为 `True`，B 值为 `False` d。

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| && | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False |
| \|\| | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。|	(A || B) 为 True |
| ! | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 |	!(A && B) 为 True |

</details>

<details>
<summary>位运算符</summary>

```go
package main
import "fmt"
func main() {

  var a uint = 60	/* 60 = 0011 1100 */  
  var b uint = 13	/* 13 = 0000 1101 */
  var c uint = 0          

  c = a & b       /* 12 = 0000 1100 */ 
  fmt.Printf("第一行 - c 的值为 %d\n", c ) // 第一行 - c 的值为 12

  c = a | b       /* 61 = 0011 1101 */
  fmt.Printf("第二行 - c 的值为 %d\n", c )  // 第二行 - c 的值为 61

  c = a ^ b       /* 49 = 0011 0001 */
  fmt.Printf("第三行 - c 的值为 %d\n", c ) // 第三行 - c 的值为 49

  c = a << 2     /* 240 = 1111 0000 */
  fmt.Printf("第四行 - c 的值为 %d\n", c ) // 第四行 - c 的值为 240

  c = a >> 2     /* 15 = 0000 1111 */
  fmt.Printf("第五行 - c 的值为 %d\n", c )  // 第五行 - c 的值为 15
}
```

Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| & | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 |	(A & B) 结果为 12, 二进制为 0000 1100 |
| \| | 按位或运算符 \| 是双目运算符。 其功能是参与运算的两数各对应的二进位相或。 |	(A \| B) 结果为 61, 二进制为 0011 1101 |
| ^ | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 |	(A ^ B) 结果为 49, 二进制为 0011 0001 |
| << | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >> | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。| A >> 2 结果为 15 ，二进制为 0000 1111 |

</details>

<details>
<summary>赋值运算符</summary>

```go
package main
import "fmt"
func main() {
  var a int = 21
  var c int

  c =  a
  fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )
  // 第 1 行 - =  运算符实例，c 值为 = 21
  c +=  a
  fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )
  // 第 2 行 - += 运算符实例，c 值为 = 42
  c -=  a
  fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )
  // 第 3 行 - -= 运算符实例，c 值为 = 21
  c *=  a
  fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )
  // 第 4 行 - *= 运算符实例，c 值为 = 441
  c /=  a
  fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )
  // 第 5 行 - /= 运算符实例，c 值为 = 21
  c  = 200; 
  c <<=  2
  fmt.Printf("第 6 行  - <<= 运算符实例，c 值为 = %d\n", c )
  // 第 6 行  - <<= 运算符实例，c 值为 = 800
  c >>=  2
  fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )
  // 第 7 行 - >>= 运算符实例，c 值为 = 200
  c &=  2
  fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )
  // 第 8 行 - &= 运算符实例，c 值为 = 0
  c ^=  2
  fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )
  // 第 9 行 - ^= 运算符实例，c 值为 = 2
  c |=  2
  fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )
  // 第 10 行 - |= 运算符实例，c 值为 = 2
}
```

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| = | 简单的赋值运算符，将一个表达式的值赋给一个左值s | C = A + B 将 A + B 表达式结果赋值给 C |
| += | 相加后再赋值s | C += A 等于 C = C + A |
| -= | 相减后再赋值s | C -= A 等于 C = C - A |
| *= | 相乘后再赋值s | C *= A 等于 C = C * A |
| /= | 相除后再赋值s | C /= A 等于 C = C / A |
| %= | 求余后再赋值s | C %= A 等于 C = C % A |
| <<= | 左移后赋值s | C <<= 2 等于 C = C << 2 |
| >>= | 右移后赋值s | C >>= 2 等于 C = C >> 2 |
| &= | 按位与后赋值s | C &= 2 等于 C = C & 2 |
| ^= | 按位异或后赋值s | C ^= 2 等于 C = C ^ 2 |
| \|= | 按位或后赋值s | C \|= 2 等于 C = C \| 2 |

</details>

<details>
<summary>其他运算符</summary>

```go
package main
import "fmt"
func main() {
  var a int = 4
  var b int32
  var c float32
  var ptr *int

  /* 运算符实例 */
  fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a ); // 第 1 行 - a 变量类型为 = int
  fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b ); // 第 2 行 - b 变量类型为 = int32
  fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c ); // 第 3 行 - c 变量类型为 = float32

  /*  & 和 * 运算符实例 */
  ptr = &a	/* 'ptr' 包含了 'a' 变量的地址 */
  fmt.Printf("a 的值为  %d\n", a);   // a 的值为  4
  fmt.Printf("*ptr 为 %d\n", *ptr); // *ptr 为 4
}
```

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| & | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
| * | 指针变量。 | *a; 是一个指针变量 |

</details>

<details>
<summary>运算符优先级</summary>

```go
package main
import "fmt"
func main() {
  var a int = 20
  var b int = 10
  var c int = 15
  var d int = 5
  var e int;
  // 通过使用括号来临时提升某个表达式的整体运算优先级。
  e = (a + b) * c / d;      // ( 30 * 15 ) / 5
  fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );
  e = ((a + b) * c) / d;    // (30 * 15 ) / 5
  fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );
  e = (a + b) * (c / d);   // (30) * (15/5)
  fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );
  e = a + (b * c) / d;     //  20 + (150/5)
  fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );  
}
```

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| 优先级 | 运算符 |
| ---- | ---- |
| 7 | ^ ! |
| 6 | * / % << >> & &^ |
| 5 | + - \| ^ |
| 4 | == != < <= >= > |
| 3 | <- |
| 2 | && |
| 1 | \|\| |

</details>

## 流程控制语句

<details>
<summary>for 循环语句</summary>

```go
package main
import "fmt"
func main() {
  sum := 0
  // 如果条件表达式的值变为 false，那么迭代将终止。
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
  
  // 循环初始化语句和后置语句都是可选的。
  // for 是 Go 的 “while”
  // 基于此可以省略分号：C 的 while 在 Go 中叫做 for 。
  // 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
  sum2 := 1
  for ; sum2 < 1000; {
    sum2 += sum2
  }
  fmt.Println(sum2)
}
```
基本的 for 循环包含三个由分号分开的组成部分：

1. 初始化语句：在第一次循环执行前被执行
1. 循环条件表达式：每轮迭代开始前被求值
1. 后置语句：每轮迭代后被执行

</details>

<details>
<summary>if 语句</summary>

```go
package main
import (
  "fmt"
  "math"
)
func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}
func main() {
  fmt.Println(sqrt(2), sqrt(-4))
}
```

就像 for 循环一样，Go 的 if 语句也不要求用 ( ) 将条件括起来，同时， { } 还是必须有的。

**if 的便捷语句**

```go
package main
import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}
func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
```

</details>

<details>
<summary>if 和 else 语句</summary>

```go
package main
import (
  "fmt"
  "math"
)
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // 这里开始就不能使用 v 了
  return lim
}

func main() {
  // 两个 pow 调用都在 main 调用 fmt.Println 前执行完毕了。
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
```

在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。

</details>

<details>
<summary>switch 语句</summary>

```go
package main
import (
  "fmt"
  "runtime"
)
func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
    case "darwin":
      fmt.Println("OS X.")
    case "linux":
      fmt.Println("Linux.")
    default:
      // freebsd, openbsd,
      // plan9, windows...
      fmt.Printf("%s.", os)
  }
}
```

在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。

**switch 的执行顺序：** 条件从上到下的执行，当匹配成功的时候停止。

```go
package main
import (
  "fmt"
  "time"
)
func main() {
  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()
  switch time.Saturday {
    case today + 0:
      fmt.Println("Today.")
    case today + 1:
      fmt.Println("Tomorrow.")
    case today + 2:
      fmt.Println("In two days.")
    default:
      fmt.Println("Too far away.")
  }
}
```

**没有条件的 switch 同 switch true 一样。**

```go
package main
import (
  "fmt"
  "time"
)
func main() {
  t := time.Now()
  switch {
    case t.Hour() < 12:
      fmt.Println("Good morning!")
    case t.Hour() < 17:
      fmt.Println("Good afternoon.")
    default:
      fmt.Println("Good evening.")
  }
}
```

</details>

<details>
<summary>defer 语句</summary>

```go
package main
import "fmt"
func main() {
  // 2. 在输出 world
  defer fmt.Println("world")
  // 1. 先输出 hello
  fmt.Println("hello")
}
```

> defer 语句会延迟函数的执行直到上层函数返回。
> 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。

**defer 栈**

延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

```go
package main
import "fmt"
func main() {
  fmt.Println("counting")
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}
```

可以运行demo [defer](example/defer/defer.go) 查看效果。

</details>

## 结构体

<details>
<summary>结构体字段</summary>

结构体字段使用点号来访问。

```go
package main
import "fmt"
type Vertex struct {
  X int
  Y int
}
func main() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)
}
```

</details>

<details>
<summary>结构体指针</summary>

结构体指针使用 & 来访问。

```go
package main
import "fmt"
type Vertex struct {
  X int
  Y int
}
func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}

```

</details>

<details>
<summary>结构体文法</summary>

> 结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。
> 使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）  
> 特殊的前缀 & 返回一个指向结构体的指针。

```go
package main
import "fmt"
type Vertex struct {
  X int
  Y int
}
func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
```

</details>

## 资源导航

<details>
<summary>官方</summary>

- [Playground](http://play.golang.org)：Go 语言代码在线运行

</details>

<details>
<summary>国内镜像</summary>

- [Go 指南国内镜像](http://tour.golangtc.com/)
- [Go 语言国内下载镜像](http://www.golangtc.com/download)
- [Go 官方网站国内镜像](http://docs.studygolang.com/)

</details>

<details>
<summary>Web 框架</summary>

- [mux](https://github.com/gorilla/mux): 一个强大的URL路由器和golang调度程序。
- [Iris](https://github.com/kataras/iris): 一个快速，简单但功能齐全且非常高效的Web框架。
- [gin](https://github.com/gin-gonic/gin):  HTTP Web框架，它具有类似Martini的API，具有更好的性能
- [Macaron](https://go-macaron.com/): 模块化 Web 框架
- [Beego](http://beego.me/): 重量级 Web 框架
- [Revel](https://github.com/revel/revel): 较早成熟的重量级 Web 框架
- [Martini](https://github.com/go-martini/martini): 一个强大为了编写模块化 Web 应用而生的 Go 语言框架
- [web](https://github.com/hoisie/web): 轻量级超简单的 web 框架。

</details>

<details>
<summary>ORM 以及数据库驱动</summary>

- [gorm](https://github.com/jinzhu/gorm): 支持 MySQL、PostgreSQL、SQLite3 以及 SQL Server
- [xorm](https://github.com/go-xorm/xorm): 支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL
- [mgo](http://labix.org/mgo): MongoDB 官方推荐驱动

</details>

<details>
<summary>辅助站点</summary>

- [Go Walker](https://gowalker.org): Go 语言在线 API 文档
- [gobuild.io](http://gobuild.io/): Go 语言在线二进制编译与下载
- [Rego](http://regoio.herokuapp.com/): Go 语言正则在线测试
- [gopm.io](https://gopm.io): 科学下载第三方包

</details>

<details>
<summary>其它文章</summary>

> 文章资源来源与Go社区的[每日新闻](https://gocn.io/topic/%E6%AF%8F%E6%97%A5%E6%96%B0%E9%97%BB)

- [使用Go和TensorFlow构建图片识别API](https://outcrawl.com/image-recognition-api-go-tensorflow/)
- [Golang C++11 中的原子操作对比](http://pp-qq.github.io/2017/10/18/go-c-atomic/)
- [Docker 镜像优化与最佳实践](http://weibo.com/ttarticle/p/show?id=2309614164586369582704&u=1889019865&m=4164582908028008&cu=2278216581)
- [微服务7篇介绍](https://www.nginx.com/blog/introduction-to-microservices/)
- [gogland eap 16 发布](https://blog.jetbrains.com/go/2017/10/18/gogland-eap-16-file-watcher-tons-of-new-inspections-smarter-navigate-to-test-and-more/)
- [在Google Go team工作体验](https://medium.com/@ljrudberg/working-on-the-go-team-at-google-917b2c8d35ff)
- [Go学习一月总结](https://blog.learngoprogramming.com/learn-go-programming-monthly-recap-a0b1e494393c)
- [devops的肮脏的秘密](https://chrisdodds.net/dirty-secrets-of-devops/)
- [从ELK到EFK](https://mp.weixin.qq.com/s/UMzq0Mt2_nm5pWn1Spba3Q)
- [我们怎么从Python切换到Go](https://getstream.io/blog/switched-python-go/)
- [Golang 微服务在腾讯游戏用户运营领域的探索及实践](http://www.infoq.com/cn/presentations/exploration-practice-of-golang-micro-service-in-the-tencent-game-users)
- [如何重构C++项目到Go](https://medium.com/@brendanleglaunec/how-refactoring-my-c-application-into-a-go-library-made-it-better-in-every-way-b99aa15fcfdf)
- [Go并发详解](https://gist.github.com/rushilgupta/228dfdf379121cb9426d5e90d34c5b96)
- [Go框架、IDE和工具集介绍](https://dzone.com/articles/golang-guide-a-list-of-top-golang-frameworks-ides)
- [Go开发大型分布式系统的好与坏](https://www.youtube.com/watch?v=8IKxf98h65Y)
- [RESTFul风格的API管理后台,基于beego和layui](https://github.com/george518/PPGo_ApiAdmin)
- [阿里云基于 Go 的微服务架构分享](https://mp.weixin.qq.com/s/Ftd8pFVCrhtppvFjBNSF5Q)
- [ElasticSearch 集群监控](http://www.54tianzhisheng.cn/2017/10/15/ElasticSearch-cluster-health-metrics/)
- [组织Go代码在一个大型repo经验之谈](https://blog.digitalocean.com/cthulhu-organizing-go-code-in-a-scalable-repo/)
- [基于openapi文档的自动化生产代码工具](https://github.com/jbowes/oag)
- [黑一下Go (实际上是作者不理解Go的interfaceO(∩_∩)O)](https://zhuanlan.zhihu.com/p/30120861)
- [基于nats的Go分布式微服务系统](https://medium.com/@shijuvar/building-distributed-systems-and-microservices-in-go-with-nats-streaming-d8b4baa633a2)
- [深入浅出 Raft - 基本概念](http://www.jianshu.com/p/138b4d267084)
- [RESTful API 设计最佳实践](http://www.zcfy.cc/article/restful-api-design-best-practices-in-a-nutshell-4388.html)
- [七个JWT最佳实践](https://dev.to/neilmadden/7-best-practices-for-json-web-tokens)
- [Go语言三驾马车](https://cloud.tencent.com/community/article/649192)
- [反射的原理与使用详解](http://www.cnblogs.com/susufufu/p/7653579.html)
- [从零构建一个神经网络](http://www.datadan.io/building-a-neural-net-from-scratch-in-go/)
- [glot图标库应用入门](https://medium.com/@Arafat./introducing-glot-the-plotting-library-for-golang-3133399948a1)
- [Go如何不适用CGO调用rust](https://speakerdeck.com/filosottile/calling-rust-from-go-without-cgo-at-gothamgo-2017)
- [GopherJS 令人惊讶的性能改进](https://medium.com/gopherjs/surprises-in-gopherjs-performance-4a0a49b04ecd)
- [Go 1.10 版本将支持编译window dll](https://go-review.googlesource.com/c/go/+/69091)
- [如何学习Go](https://dev.to/codehakase/how-i-learned-go-programming)
- [微服务组织](https://developers.redhat.com/blog/2017/08/02/organizing-microservices-modern-integration/)
- [玩转 Go Hack 之官方攻略 | 赛前指南](https://mp.weixin.qq.com/s/7mrxUxuQlOxDKxB2HTZHIw)
- [golang 标准库间依赖的可视化展示](http://blog.csdn.net/sinat_30800357/article/details/78178844)
- [golang——glide的使用手册](http://blog.csdn.net/lastsweetop/article/details/78185713)
- [hack CCTV视频库](https://github.com/EtixLabs/cameradar)
- [如何使用gRPC开发CS系统](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)
- [Go内存调优](https://blog.fmpwizard.com/2017/09/29/memory-profiling-in-go/)
- [Go奇葩点](https://i6448038.github.io/2017/10/06/GolangDetailsTwo/)
- [gRPC和RESTFul API性能对比](https://dev.to/plutov/benchmarking-grpc-and-rest-in-go-565)
- [graphql的概述](https://dzone.com/refcardz/an-overview-of-graphql)
- [2017年学习Go的视频和书籍集合](https://reactdom.com/blog/go-books)
- [《微服务：从设计到部署》](https://github.com/oopsguy/microservices-from-design-to-deployment-chinese)
- [基于openfaas的serverless服务](https://blog.alexellis.io/serverless-golang-with-openfaas/)
- [使用Go加密整个互联网](https://speakerdeck.com/filosottile/encrypting-the-internet-with-go-at-gophercon-2017)
- [go run -race的底层实现](https://speakerdeck.com/kavya719/go-run-race-under-the-hood)
- [监控和追踪Go服务](https://speakerdeck.com/chimeracoder/monitoring-and-tracing-your-go-services-gothamgo-2017)
- [如何优化高负载服务](https://blogs.dropbox.com/tech/2017/09/optimizing-web-servers-for-high-throughput-and-low-latency/)
- [Go概述](https://blog.learngoprogramming.com/about-go-language-an-overview-f0bee143597c)
- [Go之美](https://hackernoon.com/the-beauty-of-go-98057e3f0a7d)
- [k8s很酷的原因](https://jvns.ca/blog/2017/10/05/reasons-kubernetes-is-cool/)
- [Go实现的FaaS](https://github.com/fnproject/fn)
- [快速开发基于pg的API原型工具](https://github.com/dhax/go-base)
- [使用Go+gRPC打造高性能API](http://www.agiratech.com/building-high-performance-apis-go-grpc/)
- [log重负载应用的内存优化](http://agniva.me/go/2017/10/03/improving-log-write.html)
- [每个Gopher都应该值得关注的Go interface](https://www.writeingo.com/blog/important-go-interfaces/)
- [可视化调试Go程序](https://honeycomb.io/blog/2017/08/golang-observability-using-the-new-pprof-web-ui-to-debug-memory-usage/)
- [老程序员的感慨](https://mp.weixin.qq.com/s/G4KdOwoKNyTYOW3J3Wk7qA)
- [如何使用gonum进行数据统计](https://sbinet.github.io/posts/2017-10-04-intro-to-stats-with-gonum/)
- [深入理解channel笔记](https://blog.lab99.org/post/golang-2017-10-04-video-understanding-channels.html)
- [使用Go写出优雅的代码](https://scene-si.org/2017/10/04/elegant-code-and-go/)
- [Java开发者的Go指南](https://dzone.com/articles/making-the-jump-to-go-a-guide-for-java-developers)
- [Go 1.8.4 和Go 1.9.1 发布](https://groups.google.com/forum/m/#!topic/golang-nuts/sHfMg4gZNps)
- [Google开源的Go实现的配置语言](https://github.com/google/skylark)
- [老王出品的k8s完整教程](https://github.com/jolestar/kubernetes-complete-course)
- [Go实现的lua VM](https://github.com/milochristiansen/lua)
- [基于kafka和Go的微服务实现和测试](https://semaphoreci.com/community/tutorials/writing-and-testing-an-event-sourcing-microservice-with-kafka-and-go)
- [swift调用Go实例](https://medium.com/@rakyll/calling-go-from-swift-be88709942c3)
- [Go for rails](https://sphereinc.com/go-for-rails-developers/)
- [fmt包详解](https://medium.com/go-walkthrough/go-walkthrough-fmt-55a14bbbfc53)
- [基于 Linux bcc/BPF 实现 Go 程序动态追踪](http://www.jianshu.com/p/f1781fc452f6)
- [如何测试Go程序](https://getstream.io/blog/how-we-test-go-at-stream/)
- [Go如何改善性能](https://www.youtube.com/watch?v=DJ4d_PZ6Gns)

</details>
