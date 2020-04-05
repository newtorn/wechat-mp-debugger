package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
	"unicode"
)

var (
	times, sleep int
	sc           = make(chan struct{}, 1)
	tc           = make(chan struct{}, 1)
	args         = make([]string, 24)
)

func getArg(idx uint) string {
	sc <- struct{}{}
	a := args[idx]
	<-sc
	return a
}

func getArgs(s, e uint) []string {
	sc <- struct{}{}
	a := args[s : e+1]
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
	flag.StringVar(&args[0], "url", "", "Wechat mp config interface url address")
	flag.StringVar(&args[1], "token", "", "Wechat mp config token")
	flag.StringVar(&args[2], "from", "", "receieve from user openid")
	flag.StringVar(&args[3], "to", "", "send to user openid")
	flag.StringVar(&args[4], "mt", "", "Message type with text/image/voice/video/shortvideo/location/link")
	flag.StringVar(&args[5], "cont", "", "Text message content")
	flag.StringVar(&args[6], "title", "", "Link message title")
	flag.StringVar(&args[7], "desc", "", "Link message description")
	flag.StringVar(&args[8], "format", "", "voice file type like amr, speex...")
	flag.StringVar(&args[9], "recog", "", "Wechat server voice recognition result")
	flag.StringVar(&args[10], "event", "", "Event type with subscribe/unsubscribe/SCAN/LOCATION/VIEW")
	flag.StringVar(&args[11], "ekey", "", "Event key value if possible")
	flag.StringVar(&args[12], "ticket", "", "QR ticket is used to fetch QR image")
	flag.StringVar(&args[13], "lx", "", "Location X with location message")
	flag.StringVar(&args[14], "ly", "", "Location Y with location message")
	flag.StringVar(&args[15], "scale", "", "Map scale size")
	flag.StringVar(&args[16], "label", "", "Location information")
	flag.StringVar(&args[17], "mid", "", "Message media ID")
	flag.StringVar(&args[18], "tmid", "", "Message thumb media ID")
	flag.StringVar(&args[19], "purl", "", "Picture url with image message")
	flag.StringVar(&args[20], "lurl", "", "Link url with Link message")
	flag.StringVar(&args[21], "latd", "", "Location latitude with report location message")
	flag.StringVar(&args[22], "lgtd", "", "Location longitude with report location message")
	flag.StringVar(&args[23], "prec", "", "Location precision with report location message")

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
	if emptyOne(args[0:5]...) {
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

func main() {
	debug()
}
