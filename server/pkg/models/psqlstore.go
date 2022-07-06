package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ReportStore struct {
	db *gorm.DB
}

func (rs *ReportStore) Init() error {
	dsn := `user=yohan password=yohan1234 database=mydb`
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}
	rs.db = db
	err = rs.db.AutoMigrate(&Report{})
	if err != nil {
		return err
	}
	return nil
}

func (rs *ReportStore) Create(r Report) error {
	rs.db.Create(&r)
	if err := rs.db.Error; err != nil {
		return err
	}
	return nil
}
func (rs *ReportStore) GetAll() ([]Report, error) {
	var res []Report
	rs.db.Find(&res)
	if err := rs.db.Error; err != nil {
		return nil, err
	}
	return res, nil
}
