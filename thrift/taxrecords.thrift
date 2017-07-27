namespace go taxation.records

include "taxtests.thrift"

struct recordsStruct {
  1: i32 id,
  2: i32 userId (go.tag = "db:\"user_id\""),
  3: i32 testId (go.tag = "db:\"test_id\""),
  4: string content,
  5: string time
}

service taxation extends taxtests.taxtests {

  //通过练习编号获取练习记录
  recordsStruct getRecordsByTestsId(1:i32 id)

  //添加练习记录
  i32 addRecords(1:recordsStruct data)

}