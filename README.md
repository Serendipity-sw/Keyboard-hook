# Keyboard-hook

## 配置文件说明

其中config.json为程序配置文件

配置文件字段说明:

```
{
    isStartBattle  是否需要程序启动战网
    BattlePath     战网路径
    baiduAPIKey    百度api key
    baiduSecretKey 百度secretkey码
}
```

百度文字识别api key 及 secret key 申请教程 [教程地址](https://jingyan.baidu.com/article/49ad8bcefe65bd5834d8fad3.html)

魔兽登陆 排队 防暂离工具,工具调用键盘及鼠标事件,防止被检测

>> 本软件需借助MinGW-w64 第三方工具包
>
>> MinGW-w64 [下载地址](http://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win32/Personal%20Builds/mingw-builds/installer/mingw-w64-install.exe/download)
>
>> 不放心者请自行搜索 MinGW-w64 安装
>
>> MinGW-w64 [安装手册](https://www.cnblogs.com/findumars/p/8289454.html) 

>> 1. 启动并进入游戏
>> 2. 双击Keyboard-hook.exe 启动程序
>> 3. 10秒内切换到游戏界面
>
> 本人比较懒. 暂离工具运行轨迹,请查阅 main.go 源代码.  内有中文注释
