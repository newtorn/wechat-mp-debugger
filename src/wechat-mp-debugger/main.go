package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
	"unicode"
)

var (
	times, sleep int
	sc           = make(chan struct{}, 1)
	tc           = make(chan struct{}, 1)
	inArgs       = make([]string, 24)
)

func getArg(idx uint) string {
	sc <- struct{}{}
	a := inArgs[idx]
	<-sc
	return a
}

func getArgs(s, e uint) []string {
	sc <- struct{}{}
	a := inArgs[s : e+1]
	<-sc
	return a
}

func isSpace(str string) bool {
	if str == "" {
		return true
	}
	for _, ch := range str {
		if unicode.IsSpace(ch) {
			return true
		}
	}
	return false
}

func emptyOne(args ...string) bool {
	if len(args) == 0 {
		return true
	}
	for _, arg := range args {
		if isSpace(arg) {
			return true
		}
	}
	return false
}

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.String("help", "", "This will show command line")
	flag.IntVar(&times, "times", 1, "Coroutine test times for send requests")
	flag.IntVar(&sleep, "sleep", 0, "Sleep time(:ms:) before sending each coroutine request")
	flag.StringVar(&inArgs[0], "url", "", "Wechat mp config interface url address")
	flag.StringVar(&inArgs[1], "token", "", "Wechat mp config token")
	flag.StringVar(&inArgs[2], "from", "", "receieve from user openid")
	flag.StringVar(&inArgs[3], "to", "", "send to user openid")
	flag.StringVar(&inArgs[4], "mt", "", "Message type with text/image/voice/video/shortvideo/location/link")
	flag.StringVar(&inArgs[5], "cont", "", "Text message content")
	flag.StringVar(&inArgs[6], "title", "", "Link message title")
	flag.StringVar(&inArgs[7], "desc", "", "Link message description")
	flag.StringVar(&inArgs[8], "format", "", "voice file type like amr, speex...")
	flag.StringVar(&inArgs[9], "recog", "", "Wechat server voice recognition result")
	flag.StringVar(&inArgs[10], "event", "", "Event type with subscribe/unsubscribe/SCAN/LOCATION/VIEW")
	flag.StringVar(&inArgs[11], "ekey", "", "Event key value if possible")
	flag.StringVar(&inArgs[12], "ticket", "", "QR ticket is used to fetch QR image")
	flag.StringVar(&inArgs[13], "lx", "", "Location X with location message")
	flag.StringVar(&inArgs[14], "ly", "", "Location Y with location message")
	flag.StringVar(&inArgs[15], "scale", "", "Map scale size")
	flag.StringVar(&inArgs[16], "label", "", "Location information")
	flag.StringVar(&inArgs[17], "mid", "", "Message media ID")
	flag.StringVar(&inArgs[18], "tmid", "", "Message thumb media ID")
	flag.StringVar(&inArgs[19], "purl", "", "Picture url with image message")
	flag.StringVar(&inArgs[20], "lurl", "", "Link url with Link message")
	flag.StringVar(&inArgs[21], "latd", "", "Location latitude with report location message")
	flag.StringVar(&inArgs[22], "lgtd", "", "Location longitude with report location message")
	flag.StringVar(&inArgs[23], "prec", "", "Location precision with report location message")

	argNum := len(os.Args)
	if argNum >= 2 {
		if m, _ := regexp.MatchString(`(^(-help)($|\s(.*)))|(^(--help)($|=(.*)))`, os.Args[1]); m {
			flag.Usage()
			os.Exit(1)
		}
	} else if argNum < 6*2+1 {
		fmt.Println("No Enough Args")
		os.Exit(1)
	}

	flag.Parse()
	if emptyOne(inArgs[0:5]...) {
		fmt.Println(`Wechat debugger tool expected arguments "url, token, from, to, mt"`)
		os.Exit(1)
	} else if times < 1 {
		fmt.Println(`Value of argument "times" expected more than zero`)
		os.Exit(1)
	} else if times < 0 {
		fmt.Println(`Value of argument "times" expected negative int number`)
		os.Exit(1)
	}
}

func makeMessage(timestamp string) interface{} {
	var message interface{}
	switch getArg(4) {
	case "text":
		if isSpace(getArg(5)) {
			fmt.Println(`Wechat debugger tool expected argument "content"`)
			return nil
		}
		message = makeTextMessage(getArg(2), getArg(3), getArg(4),
			getArg(5), timestamp)
	case "image":
		if emptyOne(getArg(17), getArg(19)) {
			fmt.Println(`Wechat debugger tool expected arguments "mid, purl"`)
			return nil
		}
		message = makeImageMessage(getArg(2), getArg(3), getArg(4),
			getArg(17), getArg(19), timestamp)
	case "voice":
		if emptyOne(getArg(17), getArg(8)) {
			fmt.Println(`Wechat debugger tool expected arguments "mid, ft"`)
			return nil
		}
		message = makeVoiceMessage(getArg(2), getArg(3), getArg(4),
			getArg(17), getArg(8), getArg(9), timestamp)
	case "video":
		if emptyOne(getArg(17), getArg(18)) {
			fmt.Println(`Wechat debugger tool expected arguments "mid, tmid"`)
			return nil
		}
		message = makeVideoMessage(getArg(2), getArg(3), getArg(4),
			getArg(17), getArg(19), timestamp)
	case "shortvideo":
		if emptyOne(getArg(17), getArg(18)) {
			fmt.Println(`Wechat debugger tool expected arguments "mid, tmid"`)
			return nil
		}
		message = makeShortVideoMessage(getArg(2), getArg(3),
			getArg(4), getArg(17), getArg(18), timestamp)
	case "location":
		if emptyOne(getArgs(13, 16)...) {
			fmt.Println(`Wechat debugger tool expected arguments "lx, ly, sl, ll"`)
			return nil
		}
		message = makeLocationMessage(getArg(2), getArg(3), getArg(4),
			getArg(13), getArg(14), getArg(15), getArg(16), timestamp)
	case "link":
		if emptyOne(getArg(6), getArg(7), getArg(20)) {
			fmt.Println(`Wechat debugger tool expected arguments "tt, dc, lurl"`)
			return nil
		}
		message = makeLinkMessage(getArg(2), getArg(3), getArg(4),
			getArg(6), getArg(7), getArg(20), timestamp)
	case "event":
		if emptyOne(getArg(10)) {
			fmt.Println(`Wechat debugger tool expected argument "et"`)
			return nil
		}

		if getArg(10) == "subscribe" || getArg(10) == "unsubscribe" {
			message = makeSubscribeMessage(getArg(2), getArg(3), getArg(4),
				getArg(10), timestamp)
			break
		}

		if getArg(10) == "SCAN" && !emptyOne(getArg(11), getArg(12)) {
			message = makeQRSceneMessage(getArg(2), getArg(3), getArg(4),
				getArg(10), getArg(11), getArg(12), timestamp)
			break
		} else {
			fmt.Println(`Wechat debugger tool expected arguments "ek, tkt"`)
		}

		if getArg(10) == "LOCATION" && !emptyOne(getArgs(21, 23)...) {
			message = makeReportLocationMessage(getArg(2), getArg(3),
				getArg(4), getArg(10), getArg(21), getArg(22), getArg(23), timestamp)
			break
		} else {
			fmt.Println(`Wechat debugger tool expected arguments "latd, lgtd, prec"`)
		}

		if getArg(10) == "CLICK" && !isSpace(getArg(11)) {
			message = makeClickMessage(getArg(2), getArg(3), getArg(4),
				getArg(10), getArg(11), timestamp)
			break
		} else {
			fmt.Println(`Wechat debugger tool expected argument "ek"`)
		}

		if getArg(10) == "View" && !isSpace(getArg(11)) {
			message = makeViewMessage(getArg(2), getArg(3), getArg(4),
				getArg(10), getArg(11), timestamp)
			break
		} else {
			fmt.Println(`Wechat debugger tool expected argument "ek"`)
		}
		return nil
	default:
		fmt.Println("Invalid message type")
		return nil
	}
	return message
}

func main() {
	var costTick int64
	var finishedCnt = 0

	wg := sync.WaitGroup{}
	wg.Add(times)

	costTick = time.Now().UnixNano() / 1e6
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()

			timestamp := strconv.FormatInt(time.Now().Unix(), 10)
			nonce := randStringRunes(8)
			sign := signature(timestamp, nonce, getArg(1))

			var message = makeMessage(timestamp)

			if message == nil {
				return
			}

			postURL := fmt.Sprintf("%s?signature=%s&timestamp=%s&nonce=%s", getArg(0), sign, timestamp, nonce)

			xml, err := xml.Marshal(message)

			if err != nil {
				fmt.Println(err)
				return
			}

			resp, err := send(postURL, string(xml))

			if err != nil {
				fmt.Println(err)
				return
			}

			xs, errs := prettyXML(xml)
			if errs == nil && resp != nil {
				xr, errr := prettyXML(resp)
				if errr == nil {
					tc <- struct{}{}
					finishedCnt++
					<-tc
					fmt.Println("URL:", getArg(0))
					fmt.Println("--------------------")
					fmt.Println("Send Message:")
					fmt.Println(string(xs))
					fmt.Println("--------------------")
					fmt.Println("Response:")
					fmt.Println(string(xr))
					fmt.Println("--------------------")
				} else {
					fmt.Println("URL:", getArg(0))
					fmt.Println(errr)
				}
			} else {
				fmt.Println("URL:", getArg(0))
				fmt.Println(errs)
			}
			fmt.Println()
			return
		}()
		if sleep != 0 {
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}
	wg.Wait()
	costTick = time.Now().UnixNano()/1e6 - costTick

	sec := costTick / 1e3
	costTick -= sec * 1e3
	fmt.Printf("Finished! successed count: %d, fialed count: %d, cost time: %ds%dms\n", finishedCnt, times-finishedCnt, sec, costTick)
}
