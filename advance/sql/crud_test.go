package sql

import (
	"context"
	"time"
	"fmt"
	"testing"
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql" 
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type sqlTestSuite struct {
	suite.Suite

	// 配置字段
	driver string
	dsn    string

	// 初始化字段
	db *sql.DB
}

func (s *sqlTestSuite) TearDownTest() {
	_, err := s.db.Exec("DELETE FROM test_model;")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *sqlTestSuite) SetupSuite() {
	db, err := sql.Open(s.driver, s.dsn)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = s.db.ExecContext(ctx, `
CREATE TABLE IF NOT EXISTS test_model(
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    age INTEGER,
    last_name TEXT NOT NULL
)
`)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *sqlTestSuite) TestCRUD() {
	t := s.T()
	
	db, err := sql.Open(s.driver, s.dsn)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()


	{   // 连接池配置
		db.SetConnMaxLifetime(time.Second * 10)  
		db.SetMaxOpenConns(200)  
		db.SetMaxIdleConns(10) 

	}

	{   // 测试
		if err := db.Ping(); err != nil {
			t.Fatal(err)
		}
	}

	{
		// 或者 Exec(xxx)
		res, err := db.ExecContext(ctx, 
			"INSERT INTO `test_model`(`id`, `first_name`, `age`, `last_name`) VALUES (1, 'Tom', 18, 'Jerry')")  
		if err != nil {
			t.Fatal(err)
		}
		affected, err := res.RowsAffected()  
		if err != nil {
			t.Fatal(err)
		}
		if affected != 1 {
			t.Fatal(err)
		}
	}
	
	{
		rows, err := db.QueryContext(context.Background(),
			"SELECT `id`, `first_name`,`age`, `last_name` FROM `test_model` LIMIT ?", 1)
		if err != nil {
			t.Fatal()
		}
		for rows.Next() {
			tm := &TestModel{}
			err = rows.Scan(&tm.Id, &tm.FirstName, &tm.Age, &tm.LastName)
			// 常见错误，缺了指针
			// err = rows.Scan(tm.Id, tm.FirstName, tm.Age, tm.LastName)
			if err != nil {
				rows.Close()
				t.Fatal(err)
			}
			assert.Equal(t, "Tom", tm.FirstName)
		}
		rows.Close()
	}

	{
		// 或者 Exec(xxx)
		res, err := db.ExecContext(ctx, 
			"UPDATE `test_model` SET `first_name` = 'changed' WHERE `id` = ?", 1)
		if err != nil {
			t.Fatal(err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}
		if affected != 1 {
			t.Fatal(err)
		}
	}

	{
		row := db.QueryRowContext(context.Background(), 
			"SELECT `id`, `first_name`,`age`, `last_name` FROM `test_model` LIMIT 1")
		if row.Err() != nil {
			t.Fatal(row.Err())
		}
		tm := &TestModel{}

		err = row.Scan(&tm.Id, &tm.FirstName, &tm.Age, &tm.LastName)  // 结果集 - 列顺序要摆对，类型兼容或者冲突
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "changed", tm.FirstName)
	}
}

func TestMySQL(t *testing.T) {
	suite.Run(t, &sqlTestSuite{
		driver: "mysql",
		dsn:    "gopher:123456@tcp(172.23.0.1:13306)/crud?charset=utf8mb4&parseTime=True&loc=Local",
	})
}


type TestModel struct {
	Id        int64 `eorm:"auto_increment,primary_key"`
	FirstName string
	Age       int8
	LastName  *sql.NullString
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(0)
	fmt.Println(timer.Stop())
	<- timer.C
}