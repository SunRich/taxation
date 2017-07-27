package service

import (
	"taxation/tests"
	"fmt"
	"../models"
	"time"
)

type Testshandler struct {
}

//获取所有练习列表
func (a *Testshandler) GetTestsList(search *tests.Search) (r *tests.Pagedata, err error) {
	var count int32
	pageData := tests.Pagedata{}
	result := []*tests.TestsStruct{}
	models.DB.Table("tax_tests").Select("id,user_id,type,content,answer_content,create_time").Find(&result)
	models.DB.Table("articles").Count(&count)
	pageData.Testslist = result
	pageData.Count = count
	r = &pageData
	return
}

//通过练习编号获取练习详情
func (a *Testshandler) GetTestById(id int32) (r *tests.TestsStruct, err error) {
	models.DB.Table("tax_tests").Where("id = ?", id).Select("id,user_id,type,content,answer_content,create_time").Find(r)
	return
}

//添加练习
func (a *Testshandler) AddTests(data *tests.TestsStruct) (r int32, err error) {
	data.CreateTime = time.Now().Local().Format("2006-01-02 15:04:05")
	err = models.DB.Table("tax_tests").Create(&data).Error
	r = data.ID
	return
}

//编辑练习
func (a *Testshandler) EditTests(id int32, data *tests.TestsStruct) (r bool, err error) {
	err = models.DB.Table("tax_tests").Where("id = ?", id).Updates(data).Error
	if err != nil {
		r = false
		fmt.Println(err)
	} else {
		r = true
	}
	return
}

//删除练习
func (a *Testshandler) DelTests(id int32) (r bool, err error) {
	fmt.Println("tests")
	return
}
