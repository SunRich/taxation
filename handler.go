package main

import (
	"fmt"
	"records"
	"taxtests"
)

type Recordshandler struct {
}



func (a *Recordshandler)  AddRecords(data *records.RecordsStruct) (r int32, err error) {
	fmt.Println("records")
	return
}

func (a *Recordshandler) AddTests(data *taxtests.TestsStruct) (r int32, err error) {
	fmt.Println("tests")
	return
}
