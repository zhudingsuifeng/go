## go Method

一个方法只是一个函数，它有一个特殊的接收者(receiver)类型，该接收者放在func 关键字和函数名之间。接收者可以是结构体类型或者非结构体类型。可以在方法内部访问接收者。

通过下面的语法创建一个方法：

```
func (t Type) methodName(parameter list){

}
```

上面的代码片段创建一个名为methodName的方法，该方法有一个类型为Type的接收者。

编写一个简单的程序，创建一个结构体类型的方法并调用它。

```
package main

import (
    "fmt"
)

type Employee struct{
    name string
    salary int
    currency string
}

/*
displaySalary() method has Employee as the receiver type
*/

func (e Employee) displaySalary(){
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main(){
    emp1 := Employee{
        name: "bob",
        salary:5000,
        currency:"$"
    }
    emp1.displaySalary() // Calling displaySalary() method of Employee type
}
```

我们为Employee创建一个名为displaySalary的方法。在displaySalary()方法内部可以访问它的接收者e (类型为Employee).

### 为什么使用方法而不是函数

上面的程序可以使用函数而不是方法重写如下：

```
package main

import (
    "fmt"
)

type Employee struct{
    name string
    salary int
    currency string
}

/*
displaySalary() method converted to function with Employee as parameter
*/
func displaySalary(e Employee){
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main(){
    emp1 := Employee{
        name: "bob",
        salary:5000,
        currency:"$"
    }
    displaySalary(emp1)
}
```

使用displaySalary函数替换方法，并将Employee 结构体作为参数传给它。输出结果一样。

为什么我们可以使用函数完成的工作，还要使用方法？

1. Go 不是一个春面向对象的编程语言，它不支持class 类型。因此通过在一个类型上建立方法来实现与class 相似的行为。

2. 同名方法可以定义在不同的类型上，但是Go 不允许同名函数。

```
package main

import (
    "fmt"
    "math"
)

type Rectangle struct{
    length int
    width int
}

type Circle struct{
    radius float64
}

func (r Rectangle) Area() int {
    return r.length * r.width
}

func (c Circle) Area() float64{
    return math.Pi * c.radius * c.radius
}

func main(){
    r := Rectangle{
        length:10,
        width:5,
    }
    fmt.Printf("Area of rectangle %d\n", r.Area())

    c := Circle{
        radius:12,
    }
    fmt.Printf("Area of circle %f", c.Area())
}
```

### 指针接收者 VS 值接收者

值和指针都可以作为接收者。两者的区别在于，以指针作为接收者时，方法内部对其的修改对于调用者时可见的，但是作为值作为接收者却不是。

```
package main

import (
    "fmt"
)

type Employee struct{
    name string
    age int
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string){
    e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int){
    e.age = newAge
}

func main(){
    e := Employee{
        name: "fly",
        age: 30,
    }
    fmt.Printf("Employee before change name is %s and age is %d", e.name, e.age)
    e.changeName("bob")
    (&e).changeAge(24)
    fmt.Printf("Employee after change name is %s and age is %d", e.name, e.age)
}
```

程序中，changeName 方法有一个值接收者(e Employee),而 changeAge方法有一个指针接收者(e *Employee).在changeName中改变Employee的字段name的值对调用者而言时不可见的，因此程序在调用e.changeName("bob") 方法之前和之后，打印name是一致的。而因为changeAge的接收者是一个指针(e *Employee),因此通过调用方法(&e).changeAge(24)来修改age对于调用者时可见的。

changeAge有一个指针类型的接收者需要使用(&e)来调用。Go允许我们省略&符号，因此可以直接写e.changeAge(24).Go将e.changeAge(24)解析为(&e).changeAge(24).

### 何时使用指针接收者，何时使用值接收者？

一般来讲，指针接收者可用于对接收者的修改应该对调用者可见的场合。

指针接收者也可用于拷贝结构体代价较大的场合。考虑一个包含较多字段的结构体，若使用值作为接收者则必须拷贝整个结构体，这样的代价很大。这种情况下使用指针接收者将避免结构体的拷贝，而仅仅是指向结构体指针的拷贝。

其他情况下可以使用值接收者。

### 匿名字段函数

匿名字段的方法可以被包含该匿名字段的结构体的变量调用，就好像该匿名字段的方法属于包含该字段的结构体一样。

```
package main

import (
    "fmt"
)

type address struct{
    city string
    state string
}

func (a address) fulladdress(){
    fmt.Printf("Full address: %s, %s", a.city, .state)
}

type person struct{
    firstName string
    lastName string
    address
}

func main(){
    p := person{
        firstName:"Elon",
        lastName:"Musk",
        address:address{
            city:"Los Angeles",
            state:"California",
        },
    }
    p.fullAddress() //accessing fullAddress method of address struct
}
```

p.fullAddress() 调用address 的方法fullAddress().

### 方法的值接收者VS函数的值接收者

#### 函数值参数只能接受一个值参数，方法值接收者可以接受值或者指针。

```
package main

import (
    "fmt"
)

type rectangle struct{
    length int
    width int
}

func area(r rectangle){
    fmt.Printf("Area Fnction result: %d\n", (r.length * r.width))
}

func (r rectangle) area(){
    fmt.Printf("Area Method result: %d \n", (r.length * r.width))
}

func main(){
    r := rectangle{
        length:10,
        width:5,
    }
    area(r)
    r.area()

    p := &r
    /*
        compilation error, cannot use p(type *rectangle) as type rectangle in argument to area
    */
    //area(p)
    p.area() //calling value receiver with a pointer
}
```

### 方法的指针接收者VS函数的指针接收者

#### 函数指针参数只能接受指针，而方法指针接收者可以接受值接收者或者指针接收者。

```
package main

import (
    "fmt"
)

type rectangle struct{
    length int
    width int
}

func perimeter(r *rectangle){
    fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *rectangle) perimeter(){
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main(){
    r := rectangle{
        length:10,
        width:5,
    }
    p := &r     // pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type retangle) as type *rectange in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()   // calling pointer receiver with a value
}
```