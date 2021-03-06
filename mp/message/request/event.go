// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/gogap/wechat for the canonical source repository
// @license     https://github.com/gogap/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"fmt"
	"strings"

	"github.com/gogap/wechat/mp"
)

const (
	// 微信服务器推送过来的事件类型
	EventTypeSubscribe   = "subscribe"   // 订阅, 包括点击订阅和扫描二维码
	EventTypeUnsubscribe = "unsubscribe" // 取消订阅
	EventTypeScan        = "SCAN"        // 已经订阅的用户扫描二维码事件
	EventTypeLocation    = "LOCATION"    // 上报地理位置事件
)

// 关注事件(普通关注)
type SubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型，subscribe(订阅)
}

func GetSubscribeEvent(msg *mp.MixedMessage) *SubscribeEvent {
	return &SubscribeEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
	}
}

// 取消关注
type UnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型，unsubscribe(取消订阅)
}

func GetUnsubscribeEvent(msg *mp.MixedMessage) *UnsubscribeEvent {
	return &UnsubscribeEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
	}
}

// 用户未关注时，扫描带参数二维码进行关注后的事件推送
type SubscribeByScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，subscribe
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

// 获取二维码参数
func (event *SubscribeByScanEvent) Scene() (scene string, err error) {
	const prefix = "qrscene_"
	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %q 为前缀: %q", prefix, event.EventKey)
		return
	}
	scene = event.EventKey[len(prefix):]
	return
}

func GetSubscribeByScanEvent(msg *mp.MixedMessage) *SubscribeByScanEvent {
	return &SubscribeByScanEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		EventKey:            msg.EventKey,
		Ticket:              msg.Ticket,
	}
}

// 用户已关注时，扫描带参数二维码的事件推送
type ScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，SCAN
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

func GetScanEvent(msg *mp.MixedMessage) *ScanEvent {
	return &ScanEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		EventKey:            msg.EventKey,
		Ticket:              msg.Ticket,
	}
}

// 上报地理位置事件
type LocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event     string  `xml:"Event"     json:"Event"`     // 事件类型，LOCATION
	Latitude  float64 `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64 `xml:"Precision" json:"Precision"` // 地理位置精度
}

func GetLocationEvent(msg *mp.MixedMessage) *LocationEvent {
	return &LocationEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		Latitude:            msg.Latitude,
		Longitude:           msg.Longitude,
		Precision:           msg.Precision,
	}
}
