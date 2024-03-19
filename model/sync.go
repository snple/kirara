package model

import "time"

type Sync struct {
	Key     string    `bun:"type:TEXT,pk" json:"key"`
	Updated time.Time `bun:"updated" json:"updated"`
}

const (
	SYNC_PREFIX    = "sync_"
	SYNC_DEVICE    = "sync_device"
	SYNC_SLOT      = "sync_slot"
	SYNC_OPTION    = "sync_option"
	SYNC_SOURCE    = "sync_source"
	SYNC_TAG       = "sync_tag"
	SYNC_CONST     = "sync_const"
	SYNC_CLASS     = "sync_class"
	SYNC_ATTR      = "sync_attr"
	SYNC_LOGIC     = "sync_logic"
	SYNC_FN        = "sync_fn"
	SYNC_TAG_VALUE = "sync_tag_value"
)
