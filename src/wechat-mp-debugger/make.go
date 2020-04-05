package main

import (
	"fmt"
)

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
