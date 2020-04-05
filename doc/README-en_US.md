<p align="center">En | <a href="../README.md">中文</a>


# Wechat Mp Debugger

wechat public accounts message debugger

## Build
```shell
git clone https://github.com/newtorn/wechat-mp-debugger.git
cd wechat-mp-debugger/src/wechat-mp-debugger
go build

// run help
./wechat-mp-debugger -help
```

## Usage
  - cont string
    	Text message content
  - desc string
    	Link message description
  - ekey string
    	Event key value if possible
  - event string
    	Event type with subscribe/unsubscribe/SCAN/LOCATION/VIEW
  - format string
    	voice file type like amr, speex...
  - from string
    	receive from user openid
  - help string
    	This will show command line
  - label string
    	Location information
  - latd string
    	Location latitude with report location message
  - lgtd string
    	Location longitude with report location message
  - lurl string
    	Link url with Link message
  - lx string
    	Location X with location message
  - ly string
    	Location Y with location message
  - mid string
    	Message media ID
    -mt string
        	Message type with text/image/voice/video/shortvideo/location/link
  - prec string
    	Location precision with report location message
  - purl string
    	Picture url with image message
  - recog string
    	Wechat server voice recognition result
  - scale string
    	Map scale size
  - sleep int
    	Sleep time(:ms:) before sending each coroutine request
  - ticket string
    	QR ticket is used to fetch QR image
  - times int
    	Coroutine test times for send requests (default 1)
  - title string
    	Link message title
  - tmid string
    	Message thumb media ID
  - to string
    	send to user openid
  - token string
    	Wechat mp config token
  - url string
    	Wechat mp config interface url address

## Samples
- Text type message:
  ```shell
  ./wechat-mp-debugger -url http://127.0.0.1/wx -token YourMpToken -from fromOpenID -to toOpenID -mt text -cont HelloWorld -times 200 -sleep 2
  ```

- Image type message:
  ```shell
  ./wechat-mp-debugger -url http://127.0.0.1/wx -times 200 -sleep 2 -token YourMpToken -from FromOpenID -to ToOpenID -mt image -mid your_media_id -purl https://xxx.com/xxx.jpg
  ```