# 获取 c6 上传到 s7 的数据

```
getC6Media 获取视频图片
getCloudC6EventV2 废弃
getCloudC6EventV2_curl 通过 Cloud 平台获取事件列表
```

## getC6Media

需要放到服务器上面运行

`datetime.md` 是需要获取事件的日期，一行一个日期

`devices.md` 是要获取的设备，一行一个

## getCloudC6EventV2_curl

需要在 chrome 浏览器里面，执行下面这个接口，要按 F12 先打开调试面板

```
https://cloud.minieye.cc/download/devices/00270d94822e70cb/dates/2019-09-05/events_v2.csv
```

然后在网络那一栏右键选择 "Copy -> Copy as cURL"，然后粘贴到 `curltemplate.sh`
这个文件中，`#!/bin/bash` 要自己手动添加。

