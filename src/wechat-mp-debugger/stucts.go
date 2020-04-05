package main

import "encoding/xml"

// CDATAText This is origin string in xml from wechat interface xml data
type CDATAText struct {
	Text string `xml:",innerxml"`
}

// MsgBase This is base struct of all message data
type MsgBase struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	MsgType      CDATAText
	CreateTime   CDATAText
	MsgId        CDATAText `xml:",omitempty"`
}

// Text This is a text message struct
type Text struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Content CDATAText
}

// Image This is an image message struct
type Image struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	PicUrl  CDATAText
	MediaId CDATAText
}

// Voice This is a voice message struct
type Voice struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	MediaId     CDATAText
	Format      CDATAText
	Recognition CDATAText `xml:",omitempty"`
}

// Video This is a video message struct
type Video struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	MediaId      CDATAText
	ThumbMediaId CDATAText
}

// ShortVideo This is a short video message struct
type ShortVideo struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	MediaId      CDATAText
	ThumbMediaId CDATAText
}

// Location This is a location message struct
type Location struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Location_X CDATAText
	Location_Y CDATAText
	Scale      CDATAText
	Label      CDATAText
}

// Link This is a link message struct
type Link struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Title       CDATAText
	Description CDATAText
	Url         CDATAText
}

// Subscribe This is a subscribe message struct
type Subscribe struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event CDATAText
}

// QRScene This is a QR scan scene message struct
type QRScene struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event    CDATAText
	EventKey CDATAText
	Ticket   CDATAText
}

// ReportLocation This is a report location message struct
type ReportLocation struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event     CDATAText
	Latitude  CDATAText
	Longitude CDATAText
	Precision CDATAText
}

// Click This is a menu button clicked event
type Click struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event    CDATAText
	EventKey CDATAText
}

// View This is a menu url clicked event
type View struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event    CDATAText
	EventKey CDATAText
}
