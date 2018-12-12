
## How to use the vscode

### use the git in the vscode

在本地安装git，并像之前一样能够正常提交，说明个git可以使用。

在vscode中安装git支持，点击添加插件按钮，或者使用ctrl + p,后输入ext install 在列表中选择想要安装的插件。

git插件安装好之后，在vscode中添加文件 go/src/hello/hello.go

```
package main
import "fmt"
func main(){
    fmt.Println("hello, world!")
}
```

这时候点击左边栏的分支按钮，可以在最下面的CHANGES看到对文件的修改，把鼠标放到里面的文件里，点击+号，stage changes 相当于在命令行执行：

```
$git add hello.go
$git status
```
接下来，把鼠标放到根目录go上，点击状态栏最右边的更多，里面有commit 和 push选项。相当于命令行的：

```
$git commit -m 'commit words'
$git push -u origin master
$git status
```

简单的使用暂时就这些，一些细致的使用后续再研究。

### use the vim in the vscode

对于之前一直使用vim来进行开发的用户来说，vscode还是很友好的，安装vim插件之后，vscode就能支持一些简单的vim快捷键和编程习惯。

### use golang in the vscode

首先在vscode 中安装go 插件，之后编写go程序，最后debug go program。

在vscode中debug go 程序，报错：

```
cd $GOPATH
go get github.com/derekparker/delve/cmd/dlv
```

如果没有报错，就可以在vscode中直接debug go 程序了。

debug 快捷键F5.

ctrl+` 显示终端