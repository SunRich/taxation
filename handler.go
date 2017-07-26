package main

import (
	"fmt"
	"taxation/records"
	"taxation/tests"
)

type recordshandler struct {
}

func (a *recordshandler) AddRecords(data *records.TaxRecords) (r int32, err error) {
	fmt.Println("records")
	return
}

func (a *recordshandler) AddTests(data *tests.TaxTests) (r int32, err error) {
	fmt.Println("tests")
	return
}
