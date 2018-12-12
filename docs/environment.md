### How to set environment variables into ubuntu18

1. 临时设置

```
export PATH=/home/fly/go:$PATH
```

2. 当前用户的全局设置

打开～/.bashrc (如果使用的是bash[, 或者~/.zshrc使用的是zsh]),添加行：

```
export PATH=/home/fly/go:$PATH
```

使生效：

```
source .bashrc
```

3. 所有用户的全局设置

```
$vim /etc/profile
```

里面加入：

```
export PATH=/home/fly/go:$PATH
```

使生效

```
source profile
```

### Test current environment variables

```
echo $PATH
or
env
```

用户登录后加载profile和bashr的顺序：

```
1. /etc/profile
    ->/etc/profile.d/*.sh

2. $HOME/.profile
    ->$HOME/.bashrc
        ->/etc/bashrc
```

说明：

bash首先执行/etc/profile脚本,/etc/profile脚本先依次执行/etc/profile.d/*.sh

随后bash会执行用户主目录下的.profile脚本，.profile脚本会执行用户主目录下的.bashrc脚本

而.bashrc脚本会执行/etc/bashrc脚本

至此，所有的环境变量和初始化设定都已经加载完成。