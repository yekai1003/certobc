package routes

import (
	"certobc/dbs"
	"certobc/utils"
	"fmt"
	"net/http"
	"time"

	"certobc/fisco"

	"github.com/labstack/echo"
)

func PingHandler(c echo.Context) error {

	return c.String(http.StatusOK, "pong")
}

//注册
func Register(c echo.Context) error {
	//1. 响应数据结构初始化

	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	//2. 解析数据
	user := &dbs.User{}
	if err := c.Bind(user); err != nil {
		fmt.Println(user)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//3. 添加用户
	if err := user.AddUser(); err != nil {
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	return nil
}

func Login(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	//2. 解析数据
	admin := &dbs.AdminLogin{}
	if err := c.Bind(admin); err != nil {
		fmt.Println(admin)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	fmt.Println(admin)
	//3. 测试能否登陆成功
	ok, err := admin.Login()
	if !ok {
		fmt.Println("Failed to login", err)
		resp.Errno = utils.RECODE_LOGINERR
		return err
	}
	return nil
}

//证书发布
func Issue(c echo.Context) error {
	//1. 响应数据结构初始化

	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	student := struct {
		Student_id   string    `json:"student_id"`
		Student_name string    `json:"student_name"`
		Course_name  string    `json:"course_name"`
		Start_date   time.Time `json:"start_date"`
		End_date     time.Time `json:"end_date"`
		Password     string    `json:"password"`
	}{}
	//2. 解析数据
	if err := c.Bind(&student); err != nil {
		fmt.Println(student)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	fmt.Println(student)
	course := dbs.Course{}
	//查询课程
	course.CourseName = student.Course_name
	err := course.QueryCourse(course.CourseName)
	if err != nil {
		fmt.Println("Failed to QueryCourse", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//查询用户
	user := dbs.User{}
	user.UserID = student.Student_id
	err = user.QueryUser()
	if err != nil {
		fmt.Println("Failed to QueryUser", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//3. 添加用户学习记录
	history := dbs.StudyHistory{}
	history.UserID = student.Student_id
	history.CourseID = course.CourseID
	history.StartDate = student.Start_date.Local()
	history.EndDate = student.End_date.Local()
	if err := history.AddUserSH(); err != nil {
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//4. 调用区块链证书

	err = bcos.CertIssue(student.Password, fmt.Sprintf("%s", user.UUID), bcos.GetHash(course.CourseName))
	if err != nil {
		fmt.Println("Failed to CertIssue", err)
		resp.Errno = utils.RECODE_HASHERR
		return err
	}
	return nil
}

//课程添加

func AddCourse(c echo.Context) error {
	//1. 响应数据结构初始化

	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	//2. 解析数据
	course := &dbs.Course{}
	if err := c.Bind(course); err != nil {
		fmt.Println(course)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//3. 添加课程信息
	if err := course.AddCourse(); err != nil {
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	return nil
}

//证书查询 /query {}
//student_id: "1", student_name: "yekai", course_name: "blockchain"}

func Query(c echo.Context) error {
	//1. 响应数据结构初始化

	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	student := struct {
		Student_id   string `json:'student_id'`
		Student_name string `json:'student_name'`
		Course_name  string `json:'course_name'`
	}{}
	//2. 解析数据

	if err := c.Bind(&student); err != nil {
		fmt.Println(student)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	fmt.Println(student)
	//3. 获取用户uuid，课程hash
	cert := dbs.CertInfo{}
	err := cert.QueryUserStudyHistory(student.Course_name, student.Student_id)
	if err != nil {
		fmt.Println(cert)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	fmt.Println(cert)
	//4. fisco验证
	hash, err := bcos.Verify(fmt.Sprintf("%s", cert.UUID), bcos.GetHash(student.Course_name))
	if err != nil {
		fmt.Println(cert.UUID, bcos.GetHash(student.Course_name))
		resp.Errno = utils.RECODE_BCOSERR
		return err
	}
	//5. 响应数据
	cert.Hash = fmt.Sprintf("%x", hash)
	resp.Data = cert

	return nil
}
