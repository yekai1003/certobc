# certobc

## describe
Online education certificate issuing platform based on blockchain Technology

## 接口测试

### 1. 注册


```sh
curl  -H "Content-type: application/json" -X POST -d '{"user_name":"yekai","age":30, "sex":"man", "remark":"神人"}' "http://localhost:8080/register"
```

### 2. 管理员登陆 

```sh
curl  -H "Content-type: application/json" -X POST -d '{"name":"yekai","password":"123"}' "http://localhost:8080/login"
```
返回结果：
```sh
{"errno":"0","errmsg":"成功","data":null}
```



### 3. 证书发放

```sh

```

```sh
curl  -H "Content-type: application/json" -X POST -d '{"course_id":1,"user_id":1,"start_date":"2020-02-20","end_date":"2020-03-31"}' "http://localhost:8080/issue"
```



### 4. 增加课程



```
Course struct {
		CourseID   int    `json:"course_id"`
		CourseName string `json:"Course_name"`
		Remark     string `json:"remark"`
		Hash       string `json:"hash"`
	}
```

```sh
curl  -H "Content-type: application/json" -X POST -d '{"Course_name":"区块链应用工程师","remark":"oh my god，买它！"}' "http://localhost:8080/course"
```

##  部署

```sh
export MYSQL_ENV="admin:123456@tcp(10.211.55.3:3306)/cert?charset=utf8"
export FISCO_NETWORK="http://localhost:8545"
```
