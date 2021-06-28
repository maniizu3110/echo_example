package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OriginSchedule struct {
	gorm.Model
	Summary string
	Location string
	Start Start 
	End End 
}

type Start struct{
	Date time.Time
	DateTime time.Time
	TimeZone string
}
type End struct{
	Date time.Time
	DateTime time.Time
	TimeZone string
}


