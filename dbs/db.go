package dbs

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pborman/uuid"
	"github.com/yekai1003/gobcos/crypto"
)

var connstr = "admin:123456@tcp(10.211.55.3:3306)/cert?charset=utf8"

type (
	User struct {
		UserID   string    `json:"user_id"`
		UserName string    `json:"user_name"`
		Age      int       `json:"age"`
		Sex      string    `json:"sex"`
		Remark   string    `json:"remark"`
		UUID     uuid.UUID `json:"uuid"`
	}
	Course struct {
		CourseID   int    `json:"course_id"`
		CourseName string `json:"Course_name"`
		Remark     string `json:"remark"`
		Hash       string `json:"hash"`
	}

	StudyHistory struct {
		StudyID   int       `json:"study_id"`
		CourseID  int       `json:"course_id"`
		UserID    string    `json:"user_id"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
	}
	AdminLogin struct {
		Name string `json:"user"`
		Pwd  string `json:"pass"`
	}
	CertInfo struct {
		Student_name string    `json:"student_name"`
		Start_date   string    `json:"start_date"`
		End_date     string    `json:"end_date"`
		Hash         string    `json:"hash"`
		UUID         uuid.UUID `json:"uuid"`
	}
)

var DBConn *sql.DB

//init函数是本包被其他文件引用时自动执行，并且整个工程只会执行一次
func init() {
	mysqlconn := os.Getenv("MYSQL_ENV")
	if mysqlconn != "" {
		connstr = mysqlconn
	}
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	DBConn = db
}

//用户提交数据保存
func (u User) AddUser() error {
	u.UUID = uuid.NewRandom()
	_, err := DBConn.Exec("insert into t_user(user_name,age,sex,remark,uuid) values(?,?,?,?,?)",
		u.UserName, u.Age, u.Sex, u.Remark, u.UUID)
	if err != nil {
		fmt.Println("Failed to Exec insert  t_user", err)
		return err
	}
	return nil
}

//用户提交数据保存
func (u *User) QueryUser() error {
	rows, err := DBConn.Query("select user_id, user_name, age, sex, uuid from t_user where user_id = ?",
		u.UserID)
	if err != nil {
		fmt.Println("Failed to Exec insert  t_user", err)
		return err
	}
	if rows.Next() {
		err = rows.Scan(&u.UserID, &u.UserName, &u.Age, &u.Sex, &u.UUID)
		if err != nil {
			fmt.Println("Failed to scan t_user", err)
			return err
		}
	}
	return nil
}

func (c *CertInfo) QueryUserStudyHistory(course_name, userID string) error {
	rows, err := DBConn.Query("select u.uuid,h.start_date,h.end_date,u.user_name from t_user u,t_course c,t_study_history h where u.user_id = ? and c.course_name =? and c.course_id=h.course_id and u.user_id = h.user_id",
		userID, course_name)
	if err != nil {
		fmt.Println("Failed to query QueryUserStudyHistory", err)
		return err
	}
	if rows.Next() {
		err = rows.Scan(&c.UUID, &c.Start_date, &c.End_date, &c.Student_name)
		if err != nil {
			fmt.Println("Failed to scan QueryUserStudyHistory", err)
			return err
		}
	}
	return nil
}

//增加课程
func (c Course) AddCourse() error {
	c.Hash = crypto.Keccak256Hash([]byte(c.CourseName)).Hex()
	_, err := DBConn.Exec("insert into t_course(course_name,remark,hash) values(?,?,?)",
		c.CourseName, c.Remark, c.Hash)
	if err != nil {
		fmt.Println("Failed to Exec insert  t_course", err)
		return err
	}
	return nil
}

//查询课程
func (c *Course) QueryCourse(name string) error {
	rows, err := DBConn.Query("select course_id,hash from t_course where course_name =?", name)
	if err != nil {
		fmt.Println("Failed to Exec Query  t_course", err)
		return err
	}
	if rows.Next() {
		err = rows.Scan(&c.CourseID, &c.Hash)
		if err != nil {
			fmt.Println("Failed to Scan course", err)
			return err
		}
	}
	return nil
}

//提交用户学习记录
func (s StudyHistory) AddUserSH() error {
	_, err := DBConn.Exec("insert into t_study_history(course_id,user_id,start_date,end_date) values(?,?,?,?)",
		s.CourseID, s.UserID, s.StartDate.Local(), s.EndDate.Local())
	if err != nil {
		fmt.Println("Failed to Exec insert  t_study_history", err)
		return err
	}
	return nil
}

//登陆操作
func (l AdminLogin) Login() (bool, error) {
	rows, err := DBConn.Query("select 1 from t_admin where admin_name = ? and admin_pwd = ?",
		l.Name, l.Pwd)
	if err != nil {
		fmt.Println("Failed to Query", err)
		return false, err
	}
	return rows.Next(), err
}
