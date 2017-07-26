namespace php Thriftrpc.Taxation.Records

include "tests.thrift"

struct taxRecords {
  1: i32 id,
  2: i32 userId (go.tag = "db:\"user_id\""),
  3: i32 testId (go.tag = "db:\"test_id\""),
  4: string content,
  5: string time
}

service taxation extends tests.testService{
 i32 addRecords(1:taxRecords data)
}