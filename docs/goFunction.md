## 2 - Function

上一节： [第一篇 变量](https://github.com/zhudingsuifeng/go/blob/master/docs/goStart.md)
下一节： [第三篇 包](https://github.com/zhudingsuifeng/go/blob/master/docs/goPackage.md)

### 函数声明

```
func functionname(parametername type) returntype {
    //function body
}
```

函数中，参数和返回值是可选的。

```
func functionname(){

}
```

如果连续的参数具有相同的类型，只需在结束的时候写一次就可以了。

```
func calculate(price , no int ) int {
    var total = price * no
    return total
}
```

### 多个返回值

一个函数可以返回多个值。如果一个函数有多个返回值，返回值应该用小括号()括起来。

```
func  function( a, b flaot64) (float64, float64){

}
```

### 命名返回值

可以给返回值指定名字。

```
func rectProps(length,width int) (area, perimeter int){
    area = length * width
    perimeter = (length + width) * 2
    return // no explicit return value 
}
```

return 没有指定任何返回值。因为在函数声明时已经指定，遇到return 语句自动从函数中返回。

### 空指示符

下划线_表示空指示符blank identifier.可以用来代替任何类型的任何值。

```
area, _ = rectProps(10, 5)
```

我们仅获取了area, 而使用空指示符_来忽略第二个返回值perimeter。


上一节： [第一篇 变量](https://github.com/zhudingsuifeng/go/blob/master/docs/goStart.md)
下一节： [第三篇 包](https://github.com/zhudingsuifeng/go/blob/master/docs/goPackage.md)