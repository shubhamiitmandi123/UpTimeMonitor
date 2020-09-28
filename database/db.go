package database

import (
	"log"

	uuid "github.com/satori/go.uuid"

	st "up_time_monitor/structures"

	"github.com/jinzhu/gorm"
	//blank_import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database :
type Database interface {
	ConnectToDatabase() *gorm.DB
	GetURLByID(id uuid.UUID) (st.URLInfo, error)
	UpdateDatabase(st.URLInfo) error
	UpdateColumnInDatabase(uuid.UUID, string, interface{}) error
	DeleteFromDatabase(uuid.UUID) error
	Migrate(interface{}) error
	CreateDataBase(st.Request) (uuid.UUID, error)
	GetAllRecords() []st.URLInfo
}

var myDatabase Database

// GetDatabase : returns Database instance
func GetDatabase() Database {
	return myDatabase
}

// SetDatabase : sets Database instance
func SetDatabase(implDatabase Database) {
	myDatabase = implDatabase
}

// GormDatabase : To interact with mysql using gorm
type GormDatabase struct {
	Db *gorm.DB
}

// Migrate : migrate the table into database
func (r GormDatabase) Migrate(tableStruct interface{}) error {
	r.Db.AutoMigrate(tableStruct)
	return nil
}

// GetAllRecords : fetch all url records from database and return it
func (r GormDatabase) GetAllRecords() []st.URLInfo {
	var records []st.URLInfo
	r.Db.Find(&records)
	return records
}

// ConnectToDatabase : Establishes connection to database
func (r *GormDatabase) ConnectToDatabase() *gorm.DB {

	db, err := gorm.Open("mysql", dbURL(buildDBConfig()))
	if err != nil {
		log.Println("Connection Failed to Open", err.Error())
	} else {
		log.Println("Connection Established")
	}
	r.Db = db
	return db
}

// GetURLByID : for Fetching a url information from database
func (r *GormDatabase) GetURLByID(id uuid.UUID) (st.URLInfo, error) {
	var info st.URLInfo
	r.Db.First(&info, "id = ?", id)
	return info, nil
}

// UpdateDatabase : Update a recod into database
func (r *GormDatabase) UpdateDatabase(info st.URLInfo) error {
	id := info.ID
	r.Db.Model(&st.URLInfo{}).Where("id = ?", id).Updates(st.URLInfo{
		URL:      info.URL,
		Crawling: info.Crawling,
		Status:   info.Status,
	})
	r.UpdateColumnInDatabase(id, "frequency", info.Frequency)
	r.UpdateColumnInDatabase(id, "crawl_timeout", info.CrawlTimeout)
	r.UpdateColumnInDatabase(id, "failure_threshold", info.FailureThreshold)
	r.UpdateColumnInDatabase(id, "failure_count", info.FailureCount)
	return nil
}

// CreateDataBase : Creates a new record
func (r *GormDatabase) CreateDataBase(info st.Request) (uuid.UUID, error) {
	urlinfo := &st.URLInfo{
		URL:              info.URL,
		CrawlTimeout:     info.CrawlTimeout,
		Frequency:        info.Frequency,
		FailureThreshold: info.FailureThreshold,
		FailureCount:     0,
		Status:           "active",
		Crawling:         true,
	}
	if r.Db.Create(&urlinfo).Error != nil {
		log.Panic("Unable to create Record.")
	} else {
		log.Printf("Record Created")
	}
	return urlinfo.ID, nil
}

// UpdateColumnInDatabase : Update a single column into database
func (r *GormDatabase) UpdateColumnInDatabase(id uuid.UUID, columnName string, value interface{}) error {
	r.Db.Model(&st.URLInfo{}).Where("id = ?", id).UpdateColumn(columnName, value)
	return nil
}

// DeleteFromDatabase : Delete a record from database
func (r *GormDatabase) DeleteFromDatabase(id uuid.UUID) error {
	r.Db.Where("id=?", id).Delete(&st.URLInfo{})
	return nil
}
