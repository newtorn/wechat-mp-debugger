<p align="center">中文 | <a href="./doc/README-en_US.md">En</a>

# Wechat Mp Debugger

微信公众号并发调试助手

## 编译构建
```shell
git clone https://github.com/newtorn/wechat-mp-debugger.git
cd src/wechat-mp-debugger
go build

# run help
./wechat-mp-debugger -help
```



## 参数说明

  - cont string
    	Text消息类型
  - desc string
    	Link消息类型
  - ekey string
    	Event消息类型，可能是键-值
  - event string
    	Event消息类型，包含subscribe/unsubscribe/SCAN/LOCATION/VIEW
  - format string
    	语音文件类型，像amr, speex...
  - from string
    	发送者openid
  - help string
    	命令行帮助显示
  - label string
    	位置信息
  - latd string
    	带有报告位置消息类型的位置纬度
  - lgtd string
    	带有报告位置消息类型的位置经度
  - lurl string
    	带有Link消息类型的链接url
  - lx string
    	带位置信息的位置X
  - ly string
    	带位置信息的位置Y
  - mid string
    	消息的media ID
  - mt string
    	消息类型，可选范围：text/image/voice/video/shortvideo/location/link
  - prec string
    	带有报告位置消息的位置精度
  - purl string
    	图片消息的图片url
  - recog string
    	微信服务器的语音识别结果
  - scale string
    	地图缩放大小
  - sleep int
    	每个测试协程的时间间隔，单位毫秒
  - ticket string
    	用于获取二维码图像的ticket
  - times int
    	需要进行测试次数
  - title string
    	Link消息类型的标题
  - tmid string
    	缩略图media ID
  - to string
    	接收者openid
  - token string
    	微信公众号后台配置的token
  - url string
    	微信公众号后台配置的接口地址url



## 样例

- 文本类型消息:
  ```shell
  ./wechat-mp-debugger -url http://127.0.0.1/wx -token YourMpToken -from fromOpenID -to toOpenID -mt text -cont HelloWorld -times 200 -sleep 2
  ```

- 图片类型消息:
  ```shell
  ./wechat-mp-debugger -url http://127.0.0.1/wx -times 200 -sleep 2 -token YourMpToken -from FromOpenID -to ToOpenID -mt image -mid your_media_id -purl https://xxx.com/xxx.jpg
  ```

