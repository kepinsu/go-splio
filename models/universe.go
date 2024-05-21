package models

import "time"

type UniverseConsumption struct {
	SmsCredit       uint            `json:"sms_credit,omitempty"`
	Usage30D        *uint           `json:"usage_30_d,omitempty"`
	MaxSmsVolume30D *uint           `json:"max_sms_volume_30_d,omitempty"`
	SmsPostPaid     bool            `json:"sms_post_paid,omitempty"`
	Stats           []UniverseStats `json:"stats,omitempty"`
}

type UniverseStats struct {
	Individuals uint      `json:"individuals,omitempty"`
	EmailsSent  uint      `json:"emails_sent,omitempty"`
	EmailsTried uint      `json:"emails_tried,omitempty"`
	SmsTried    uint      `json:"sms_tried,omitempty"`
	SmsSent     uint      `json:"sms_sent,omitempty"`
	FwdSms      uint      `json:"fwd_sms,omitempty"`
	Fwd_email   uint      `json:"fwd__email,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

type UniverseUniqueKey struct {
	UniqueKey string `json:"unique_key,omitempty"`
}
