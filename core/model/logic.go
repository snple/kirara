package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Fn struct {
	bun.BaseModel `bun:"fn"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	DeviceID      string    `bun:"device_id,type:TEXT" json:"device_id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Exec          string    `bun:"exec,type:TEXT" json:"exec"`
	Main          string    `bun:"main,type:TEXT" json:"main"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Debug         int32     `bun:"debug" json:"debug"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

type Logic struct {
	bun.BaseModel `bun:"logic"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	DeviceID      string    `bun:"device_id,type:TEXT" json:"device_id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Exec          string    `bun:"exec,type:TEXT" json:"exec"`
	Main          string    `bun:"main,type:TEXT" json:"main"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}
