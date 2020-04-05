package main

import "encoding/xml"

// CDATAText is origin string in xml from wechat interface xml data
type CDATAText struct {
	Text string `xml:",innerxml"`
}

// MsgBase is base struct of all message data
type MsgBase struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	MsgType      CDATAText
	CreateTime   CDATAText
	MsgID        CDATAText `xml:",omitempty"`
}

// Text is a text message struct
type Text struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Content CDATAText
}

// Image is an image message struct
type Image struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	PicURL  CDATAText
	MediaID CDATAText
}

// Voice is a voice message struct
type Voice struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	MediaID     CDATAText
	Format      CDATAText
	Recognition CDATAText `xml:",omitempty"`
}

// Video is a video message struct
type Video struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	MediaID      CDATAText
	ThumbMediaID CDATAText
}

// ShortVideo is a short video message struct
type ShortVideo struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	MediaID      CDATAText
	ThumbMediaID CDATAText
}

// Location is a location message struct
type Location struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	LocationX CDATAText
	LocationY CDATAText
	Scale     CDATAText
	Label     CDATAText
}

// Link is a link message struct
type Link struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Title       CDATAText
	Description CDATAText
	URL         CDATAText
}

// Subscribe is a subscribe message struct
type Subscribe struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event CDATAText
}

// QRScene is a QR scan scene message struct
type QRScene struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event    CDATAText
	EventKey CDATAText
	Ticket   CDATAText
}

// ReportLocation is a report location message struct
type ReportLocation struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event     CDATAText
	Latitude  CDATAText
	Longitude CDATAText
	Precision CDATAText
}

// Click is a menu button clicked event
type Click struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event    CDATAText
	EventKey CDATAText
}

// View is a menu url clicked event
type View struct {
	XMLName xml.Name `xml:"xml"`
	MsgBase
	Event    CDATAText
	EventKey CDATAText
}
