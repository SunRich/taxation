namespace go taxation.tests

struct testsStruct {
  1: i32 id,
  2: i32 userId (go.tag = "db:\"user_id\""),
  3: string type,
  4: string content,
  5: string answerContent (go.tag = "db:\"answer_content\""),
  6: string createTime  (go.tag = "db:\"create_time\""),
  7: string updateTime  (go.tag = "db:\"update_time\""),
}

//筛选
struct Search {
  1: string keywords,
  3: i8 p
}

//分页文章列表
struct pagedata {
  1: i32 count,
  2: list<testsStruct> testslist
}

service taxtests{

 //获取报税练习列表
 pagedata getTestsList(1:Search search)

 //通过练习编号获取报税练习信息
 testsStruct getTestById(1:i32 id)

 //添加报税练习
 i32 addTests(1:testsStruct data)

 //编辑报税练习
 bool editTests(1:i32 id, 2:testsStruct data)

 //删除报税练习
 bool delTests(1:i32 id),
}

