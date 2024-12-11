## 使用 make 命令创建新模型

* 模型文件
```bash
go run main.go make model model_name
```
* 迁移文件
```bash
go run main.go make migration model_name
```
* 请求验证文件
```bash
go run main.go make request model_name
```
* 控制器文件
```bash
go run main.go make apicontroller v1/model_name
```
* 数据工厂文件
```bash
go run main.go make factory model_name
```
* 批量生成数据文件
```bash
go run main.go make seeder model_name
```
* 授权策略
```bash
go run main.go make policy project
```

* 数据库迁移：
```bash
// 生成数据表
go run main.go migrate up
// 填充数据
go run main.go seed SeedCategoriesTable
```