package service

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"../models"
	"taxation/tests"
	"taxation/records"
)

func init() {
	models.InitMysql()
}

func TestTests(t *testing.T) {
	var data, getData *tests.TestsStruct
	Convey("测试添加练习", t, func() {
		Testshandler := &Testshandler{}
		data = &tests.TestsStruct{UserId: 2, Content: "ssss"}
		id, _ := Testshandler.AddTests(data)
		So(id, ShouldBeGreaterThan, 0)
		Convey("测试获取练习", func() {
			getData, _ = Testshandler.GetTestById(id)
			So(getData.Content, ShouldEqual, data.Content)
		})

		Convey("测试编辑练习", func() {
			data = &tests.TestsStruct{UserId: 3, Content: "ssss"}
			result, _ := Testshandler.EditTests(id, data)
			So(result, ShouldBeTrue)
			getData, _ = Testshandler.GetTestById(id)
			So(getData.UserId, ShouldEqual, data.UserId)

		})
	})

}

func TestRecords(t *testing.T) {
	Convey("测试添加练习记录", t, func() {
		var data, getData *records.RecordsStruct
		Recordshandler := &Recordshandler{}
		data = &records.RecordsStruct{UserId: 2, TestId:10,Content: "ssss"}
		id,_:=Recordshandler.AddRecords(data)
		So(id, ShouldBeGreaterThan, 0)
		Convey("测试获取练习记录", func() {
			getData, _ = Recordshandler.GetRecordsByTestsId(data.TestId)
			So(getData.UserId, ShouldEqual, data.UserId)
		})
	})
}
