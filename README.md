# MySQL 数据库命令行操做小工具

1. 支持传入SQL文件参数的形式，例：-f=./test.sql
2. 支持命令行直接传SQL语句的形式，例：-sql="select * from test"
3. 支持不同的操作类型，例：-op=create
   1. create: 建表
   2. insert: 插入数据
   3. update: 更新数据
   4. delete: 删除数据
   5. query: 查询数据
4. 查询语句支持输出到文件，例：-output=true -path=./output.txt
5. 支持导出数据指定分隔符，例：-split=, 
6. 支持 SQL 语句中传入变量：
   1. SQL 中写法：`select * from test_table2 where created_at <= '${mysqlvar:today}' and created_at >= '${mysqlvar:yesterday}';`
   2. 命令行参数写法：-mysqlvar=yesterday:2021-04-26 -mysqlvar=today:2021-04-27
7. 相关使用帮助：-h 参数查看

