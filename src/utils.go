package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func randStringRunes(n int) string {
	lrs := []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bt := make([]rune, n)
	for idx := range bt {
		bt[idx] = lrs[rand.Intn(len(lrs))]
	}
	return string(bt)
}

func prettyXML(data []byte) ([]byte, error) {
	bt := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(bt)
	encoder.Indent("", "	")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return bt.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}

func strToCDATA(str string) CDATAText {
	return CDATAText{"<![CDATA[" + str + "]]>"}
}

func signature(timestamp string, nonce string, token string) string {
	strls := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strls)
	str := ""
	for _, s := range strls {
		str += s
	}
	hs := sha1.New()
	hs.Write([]byte(str))
	return fmt.Sprintf("%x", hs.Sum(nil))
}

func (msg *MsgBase) initMessage(from, to, msgtype, timestamp string) {
	msg.FromUserName = strToCDATA(from)
	msg.ToUserName = strToCDATA(to)
	msg.MsgType = strToCDATA(msgtype)
	msg.MsgId = strToCDATA(strconv.FormatInt(rand.Int63(), 10))
	msg.CreateTime = strToCDATA(timestamp)
}

func makeTextMessage(from, to, msgtype, content, timestamp string) Text {
	var message Text
	message.initMessage(from, to, msgtype, timestamp)
	message.Content = strToCDATA(content)
	return message
}

func makeImageMessage(from, to, msgtype, mid, purl, timestamp string) Image {
	var message Image
	message.initMessage(from, to, msgtype, timestamp)
	message.MediaId = strToCDATA(mid)
	message.PicUrl = strToCDATA(purl)
	return message
}

func makeVoiceMessage(from, to, msgtype, mid, ft, rg, timestamp string) Voice {
	var message Voice
	message.initMessage(from, to, msgtype, timestamp)
	message.MediaId = strToCDATA(mid)
	message.Format = strToCDATA(ft)
	message.Recognition = strToCDATA(rg)
	return message
}

func makeVideoMessage(from, to, msgtype, mid, tmid, timestamp string) Video {
	var message Video
	message.initMessage(from, to, msgtype, timestamp)
	message.MediaId = strToCDATA(mid)
	message.ThumbMediaId = strToCDATA(tmid)
	return message
}

func makeShortVideoMessage(from, to, msgtype, mid, tmid, timestamp string) ShortVideo {
	var message ShortVideo
	message.initMessage(from, to, msgtype, timestamp)
	message.MediaId = strToCDATA(mid)
	message.ThumbMediaId = strToCDATA(tmid)
	return message
}

func makeLocationMessage(from, to, msgtype, lx, ly, sl, ll, timestamp string) Location {
	var message Location
	message.initMessage(from, to, msgtype, timestamp)
	message.Location_X = strToCDATA(lx)
	message.Location_Y = strToCDATA(ly)
	message.Scale = strToCDATA(sl)
	message.Label = strToCDATA(ll)
	return message
}

func makeLinkMessage(from, to, msgtype, tt, dc, lurl, timestamp string) Link {
	var message Link
	message.initMessage(from, to, msgtype, timestamp)
	message.Title = strToCDATA(tt)
	message.Description = strToCDATA(dc)
	message.Url = strToCDATA(lurl)
	return message
}

func makeSubscribeMessage(from, to, msgtype, et, timestamp string) Subscribe {
	var message Subscribe
	message.initMessage(from, to, msgtype, timestamp)
	message.Event = strToCDATA(et)
	return message
}

func makeQRSceneMessage(from, to, msgtype, et, ek, tkt, timestamp string) QRScene {
	var message QRScene
	message.initMessage(from, to, msgtype, timestamp)
	message.Event = strToCDATA(et)
	message.EventKey = strToCDATA(ek)
	message.Ticket = strToCDATA(tkt)
	return message
}

func makeReportLocationMessage(from, to, msgtype, et, latd, lgtd, prec, timestamp string) ReportLocation {
	var message ReportLocation
	message.initMessage(from, to, msgtype, timestamp)
	message.Event = strToCDATA(et)
	message.Latitude = strToCDATA(latd)
	message.Longitude = strToCDATA(lgtd)
	message.Precision = strToCDATA(prec)
	return message
}

func makeClickMessage(from, to, msgtype, et, ek, timestamp string) Click {
	var message Click
	message.initMessage(from, to, msgtype, timestamp)
	message.Event = strToCDATA(et)
	message.EventKey = strToCDATA(ek)
	return message
}

func makeViewMessage(from, to, msgtype, et, ek, timestamp string) View {
	var message View
	message.initMessage(from, to, msgtype, timestamp)
	message.Event = strToCDATA(et)
	message.EventKey = strToCDATA(ek)
	return message
}

func send(url string, message string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(message))
	if err != nil {
		return nil, err
	}
	defer func() { req.Close = true }()

	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
