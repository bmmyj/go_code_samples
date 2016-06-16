
package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:p@ssw0rdYJ@tcp(localhost:3306)/test?charset=utf8")
	db.SetMaxOpenConns(40)
	db.SetMaxIdleConns(20)
	db.Ping()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//插入数据
func insert() {
	stmt, err := db.Prepare(`INSERT picture_table (pic_id,pic_title,pic_url) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(22, "hello", "http://www.hkjlsd.com/hello.png")
	//res, err := stmt.Exec(24, "small pic2", "http://www.bmm.com/small2.png")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

// 删除数据
func delete() {
	stmt, err := db.Prepare(`DELETE FROM picture_table WHERE pic_id=?`)
	checkErr(err)
	res, err := stmt.Exec(22)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//查询数据
func query() {
	rows, err := db.Query("SELECT * FROM picture_table")
	checkErr(err)

	//普通demo
	//for rows.Next() {
	//    var userId int
	//    var userName string
	//    var userAge int
	//    var userSex int

	//    rows.Columns()
	//    err = rows.Scan(&userId, &userName, &userAge, &userSex)
	//    checkErr(err)

	//    fmt.Println(userId)
	//    fmt.Println(userName)
	//    fmt.Println(userAge)
	//    fmt.Println(userSex)
	//}

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//更新数据
func update() {

	stmt, err := db.Prepare(`UPDATE picture_table SET pic_title=?,pic_url=? WHERE pic_id=?`)
	checkErr(err)
	res, err := stmt.Exec("BigPic", "http://www.baidu.com/2.png", 22)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func main(){
	insert()
	//query()
	//update()
	//delete()
}