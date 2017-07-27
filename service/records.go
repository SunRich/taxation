package service

import (
	"taxation/records"
	"fmt"
)

type Recordshandler struct {
}

func (a *Recordshandler) GetRecordsByTestsId(id int32) (r *records.RecordsStruct, err error) {
	return
}

func (a *Recordshandler) AddRecords(data *records.RecordsStruct) (r int32, err error) {
	fmt.Println("records")
	return
}
