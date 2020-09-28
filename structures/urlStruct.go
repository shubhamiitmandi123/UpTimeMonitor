package structures

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;not null;autoIncrement:false"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// URLInfo : For Storing data of url
type URLInfo struct {
	Base
	URL              string `json:"url"`
	CrawlTimeout     int    `json:"crawl_timeout" gorm:"not null"`
	Frequency        int    `json:"frequency" gorm:"not null"`
	FailureThreshold int    `json:"failure_threshold" gorm:"not null"`
	FailureCount     int    `json:"failurecount"`
	Status           string `json:"statushii"`
	Crawling         bool   `json:"crawling"`
}
