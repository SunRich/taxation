package service

import (
	"taxation/records"
	"../models"
	"time"
)

type Recordshandler struct {
}

func (a *Recordshandler) GetRecordsByTestsId(id int32) (r *records.RecordsStruct, err error) {
	models.DB.Table("tax_records").Where("id = ?", id).Select("id,user_id,test_id,content,time").Find(r)
	return
}

func (a *Recordshandler) AddRecords(data *records.RecordsStruct) (r int32, err error) {
	data.Time = time.Now().Local().Format("2006-01-02 15:04:05")
	err = models.DB.Table("tax_records").Create(&data).Error
	r = data.ID
	return
}
