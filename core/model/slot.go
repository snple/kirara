package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Slot struct {
	bun.BaseModel `bun:"slot"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	DeviceID      string    `bun:"device_id,type:TEXT" json:"device_id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Secret        string    `bun:"secret,type:TEXT" json:"secret"`
	Location      string    `bun:"location,type:TEXT" json:"location"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}
