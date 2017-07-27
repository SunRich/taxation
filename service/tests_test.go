package service

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"../models"
	"taxation/tests"
	"fmt"
)

func init() {
	models.InitMysql()
}

func TestAddTests(t *testing.T) {
	Convey("测试添加文章", t, func() {
		Testshandler := &Testshandler{}
		data := tests.TestsStruct{UserId: 2, Content: "ssss"}
		id, _ := Testshandler.AddTests(&data)
		fmt.Println(id)
	})
}
