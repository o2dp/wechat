// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/gogap/wechat for the canonical source repository
// @license     https://github.com/gogap/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"net/http"
	"net/url"
)

// 微信服务器推送过来的消息(事件)处理接口
type MessageHandler interface {
	ServeMessage(w http.ResponseWriter, r *Request)
}

type MessageHandlerFunc func(http.ResponseWriter, *Request)

func (fn MessageHandlerFunc) ServeMessage(w http.ResponseWriter, r *Request) {
	fn(w, r)
}

// 消息(事件)请求信息
type Request struct {
	HttpRequest *http.Request // 可以为 nil, 因为某些 http 框架没有提供此参数

	// 下面的字段必须提供, 如果有的话

	QueryValues url.Values // 回调请求 URL 中的查询参数集合
	Signature   string     // 回调请求 URL URL 中的签名: signature
	Timestamp   int64      // 回调请求 URL URL 中的时间戳: timestamp
	Nonce       string     // 回调请求 URL URL 中的随机数: nonce

	RawMsgXML []byte        // "明文"消息的 XML 文本
	MixedMsg  *MixedMessage // RawMsgXML 解析后的消息

	// 下面的字段是 AES 模式才有的
	MsgSignature string   // 请求 URL 中的消息体签名: msg_signature
	EncryptType  string   // 请求 URL 中的加密方式: encrypt_type
	AESKey       [32]byte // 当前消息 AES 加密的 key
	Random       []byte   // 当前消息加密时所用的 random, 16 bytes

	// 下面字段是公众号的基本信息, 回包需要
	WechatAppId string // 请求消息所属公众号的 AppId
	WechatToken string // 请求消息所属公众号的 Token
}

// 微信服务器推送过来的消息(事件)通用的消息头
type CommonMessageHeader struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string `xml:"MsgType"      json:"MsgType"`
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMessage struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonMessageHeader

	// fuck, MsgId != MsgID
	MsgId int64 `xml:"MsgId" json:"MsgId"`
	MsgID int64 `xml:"MsgID" json:"MsgID"`

	Content      string  `xml:"Content"      json:"Content"`
	MediaId      string  `xml:"MediaId"      json:"MediaId"`
	PicURL       string  `xml:"PicUrl"       json:"PicUrl"`
	Format       string  `xml:"Format"       json:"Format"`
	Recognition  string  `xml:"Recognition"  json:"Recognition"`
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"   json:"Location_X"`
	LocationY    float64 `xml:"Location_Y"   json:"Location_Y"`
	Scale        int     `xml:"Scale"        json:"Scale"`
	Label        string  `xml:"Label"        json:"Label"`
	Title        string  `xml:"Title"        json:"Title"`
	Description  string  `xml:"Description"  json:"Description"`
	URL          string  `xml:"Url"          json:"Url"`

	Event    string `xml:"Event"    json:"Event"`
	EventKey string `xml:"EventKey" json:"EventKey"`

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`
		ScanResult string `xml:"ScanResult" json:"ScanResult"`
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"`

	SendPicsInfo struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMD5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo" json:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"`
		LocationY float64 `xml:"Location_Y" json:"Location_Y"`
		Scale     int     `xml:"Scale"      json:"Scale"`
		Label     string  `xml:"Label"      json:"Label"`
		Poiname   string  `xml:"Poiname"    json:"Poiname"`
	} `xml:"SendLocationInfo" json:"SendLocationInfo"`

	Ticket      string  `xml:"Ticket"      json:"Ticket"`
	Latitude    float64 `xml:"Latitude"    json:"Latitude"`
	Longitude   float64 `xml:"Longitude"   json:"Longitude"`
	Precision   float64 `xml:"Precision"   json:"Precision"`
	Status      string  `xml:"Status"      json:"Status"`
	TotalCount  int     `xml:"TotalCount"  json:"TotalCount"`
	FilterCount int     `xml:"FilterCount" json:"FilterCount"`
	SentCount   int     `xml:"SentCount"   json:"SentCount"`
	ErrorCount  int     `xml:"ErrorCount"  json:"ErrorCount"`
	OrderId     string  `xml:"OrderId"     json:"OrderId"`
	OrderStatus int     `xml:"OrderStatus" json:"OrderStatus"`
	ProductId   string  `xml:"ProductId"   json:"ProductId"`
	SKUInfo     string  `xml:"SkuInfo"     json:"SkuInfo"`

	// card
	CardId         string `xml:"CardId"         json:"CardId"`
	IsGiveByFriend int    `xml:"IsGiveByFriend" json:"IsGiveByFriend"`
	FriendUserName string `xml:"FriendUserName" json:"FriendUserName"`
	UserCardCode   string `xml:"UserCardCode"   json:"UserCardCode"`
	OuterId        int64  `xml:"OuterId"        json:"OuterId"`

	// poi
	UniqId string `xml:"UniqId" json:"UniqId"`
	PoiId  string `xml:"PoiId"  json:"PoiId"`
	Result string `xml:"Result" json:"Result"`
	Msg    string `xml:"Msg"    json:"Msg"`
}
