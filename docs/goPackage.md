## 3 - Package

上一节： [第二篇 函数](https://github.com/zhudingsuifeng/go/blob/master/docs/goFunction.md)

下一节： [第四篇 逻辑控制](https://github.com/zhudingsuifeng/go/blob/master/docs/goLogc.md)

包用于组织Go源代码，以获得更好的重用性和可读性。

每个可执行的go程序都必须包含一个main函数。这个函数是执行程序的入口。main函数应该包含在main包中。

指定一个特定源文件属于一个包的语法为： 

```
package packagename
```

这条语句应该放在源文件的第一行。

属于同一个包的源文件应该放在独立的文件夹中，按照go的惯例，该文件夹的名字应该与包名相同。

在同一个目录下（不包括子目录）的go文件必须属于同一个package,不然报错：

```
/*
{
	"resource": "/home/fly/go/src/hello/test.go",
	"owner": "go",
	"severity": 8,
	"message": "can't load package: package hello: found packages hello (hello.go) and main (test.go) in /home/fly/go/src/hello\n",
	"source": "go",
	"startLineNumber": 1,
	"startColumn": 1,
	"endLineNumber": 1,
	"endColumn": 13
}
*/
```

每个包都可以包含一个init函数。该函数不应该有任何参数和返回值，并且在我们的代码中不能显式调用它。



上一节： [第二篇 函数](https://github.com/zhudingsuifeng/go/blob/master/docs/goFunction.md)

下一节： [第四篇 逻辑控制](https://github.com/zhudingsuifeng/go/blob/master/docs/goLogc.md)