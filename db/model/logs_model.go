package model

import (
	"time"
)

// Logs 日志表
type Logs struct {
	ReqId     string         `gorm:"comment:请求id;type:uuid;primary_key;<-:create" json:"req_id"`
	IP        string         `gorm:"comment:ip地址;type:varchar(64);NOT NULL;<-:create" json:"ip"`
	Url       string         `gorm:"comment:请求的url;type:varchar(256);NOT NULL;<-:create" json:"url"`
	Method    string         `gorm:"comment:请求的方法;type:varchar(64);NOT NULL;<-:create" json:"method"`
	Params    string         `gorm:"comment:请求的参数;type:varchar(1024);<-:create" json:"params"`
	Header    string         `gorm:"comment:请求的header;type:varchar(1024);<-:create" json:"header"`
	Body      map[string]any `gorm:"comment:请求的body;type:json;<-:create" json:"body"`
	Resp      map[string]any `gorm:"comment:响应的body;type:json;<-:create" json:"resp"`
	CreatedAt time.Time      `gorm:"comment:创建时间;type:timestamp(0);NOT NULL;<-:create" json:"created_at,omitempty"`
}

// Create 创建一条日志信息
func (l *Logs) Create() *Logs {
	db.Create(&l)
	return l
}