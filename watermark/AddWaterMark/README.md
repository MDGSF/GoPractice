# 添加水印使用说明

## 查看帮助文档

```
$ AddWaterMark.exe -h          
Usage:
  AddWaterMark [flags]
  AddWaterMark [command]

Available Commands:
  help        Help about any command
  version     Show AddWaterMark version.

Flags:
  -a, --Alpha uint8         text color (default 20)
  -b, --Blue uint8          text color (default 255)
  -g, --Green uint8         text color (default 255)
  -r, --Red uint8           text color (default 255)
  -w, --WaterMarkType int   watermark type
  -f, --font float          font size (default 42)
  -h, --help                help for AddWaterMark
  -o, --output string       output directory
  -s, --source string       image file or directory
  -e, --suffix string       new image suffix (default "_marked")
  -t, --text string         watermark text (default "minieye")
  -v, --version             Show AddWaterMark version.
  -k, --workers int         worker thread number

Use "AddWaterMark [command] --help" for more information about a command.
```

## 参数解释

```
-s 设置要添加水印的图片目录

-t 设置水印文字，默认是 "minieye"

-e 设置添加水印之后的生成新的图片的后缀。
比如原始图片是 test.jpg，添加水印之后生成新的图片名字为 test_marked.jpg
如果把这个值设置为空字符串，如 -e "" 这样，那会覆盖掉原始图片。

-o 设置新生成的图片保存的位置，没有设置的话，则默认保存在和原始图片同一个目录下。

-a 设置水印透明度，数值越小越透明

-r, -g, -b 这个 3 个参数是设置水印的颜色，也就是 RGB 的值，默认值是 (255, 255, 255)，也就是白色。
常用颜色:
白色： (255, 255, 255)
黑色： (0, 0, 0)
红色： (255, 0, 0)
绿色： (0, 255, 0)
蓝色： (0, 0, 255)
黄色： (255, 255, 0)

-f 设置水印的字体大小，默认值是 42，数值越大字体越大

-w 设置水印类型
-w 0 默认类型，全图片，斜着的水印
-w 1 左上角
-w 2 右上角
-w 3 右下角
-w 4 左下角

-k 设置工作线程的数量，这个默认值是 cpu 的数量。
取值范围从 0 ~ cpu数量。取值至少为 1，也就是一个线程。
一般设置为 3、4 个工作线程时，速度较快。
```

## 例子

```
查看帮助文档
AddWaterMark.exe -h 

给目录 imageDir 下面的所有图片添加水印，
水印默认内容是 "minieye"，水印默认白色，默认透明度 20，默认字体大小 42
AddWaterMark.exe -s imageDir

给目录 imageDir 下面的所有图片添加水印，水印添加在图片的右下角
水印内容是 "China"，水印默认白色，默认透明度 20，默认字体大小 42
AddWaterMark.exe -s imageDir -w 3 -t "China"

修改水印透明度为 100
AddWaterMark.exe -s imageDir -w 3 -t "China" -a 100
```
