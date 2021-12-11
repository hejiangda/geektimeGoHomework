1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

回答：

应该`Wrap`这个`error`，因为dao层本身没法判断遇到`sql.ErrNoRows`后要怎么做（是结束程序，是再发一次查询请求，还是什么也不做）。所以按照 “If the error is not going to be handled, wrap and return up the call stack.” 的原则，要把额外的上下文如输入参数或失败的查询语句`warp`，为之后的处理排错提供依据。

数据库中内容：
![alt 数据库中内容](./databaseInfo.png)

程序输出结果：

```
[{sam 1}]
2021/12/11 20:03:53 Query Id:2 not found!: sql: no rows in result set
exit status 1
```
