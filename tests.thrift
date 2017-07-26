namespace php Thriftrpc.Taxation.Tests

struct taxTests {
  1: i32 id,
  2: i32 userId (go.tag = "db:\"user_id\""),
  3: string type,
  4: string content,
  5: string answerContent (go.tag = "db:\"answer_content\""),
  6: string createTime  (go.tag = "db:\"create_time\""),
  7: string updateTime  (go.tag = "db:\"update_time\""),
}

service testService{
 i32 addTests(1:taxTests data)
}

