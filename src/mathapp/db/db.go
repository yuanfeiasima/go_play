package db


import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func SelectNews() (s string){
	db, err := sql.Open("mysql", "root:123456@/proginn_bitcoin?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select * from news;")
	checkErr(err)

	//普通方法
	//for rows.Next(){
	//	var id int64
	//	var title sql.NullString
	//	var content sql.NullString
	//	var status sql.NullString
	//	var creationDate sql.NullString
	//	var createdId sql.NullString
	//	var createdBy sql.NullString
	//	var changeDate sql.NullString
	//	var changedId sql.NullString
	//	var changedBy sql.NullString
	//	var _type sql.NullString
	//
	//	err = rows.Scan(&id, &title, &content, &status,
	//		&creationDate, &createdId, &createdBy,
	//		&changeDate, &changedId, &changedBy, &_type)
	//	checkErr(err)
	//	if title.Valid {
	//		s += title.String + "_"
	//	} else{
	//		s += "_"
	//
	//	}
	//}

	//返回字典
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next()  {
		err = rows.Scan(scanArgs)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil{
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	return "ok"

}

func insert(title string, content string)(s string){
	db, err := sql.Open("mysql", "root:123456@/proginn_bitcoin?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare("insert into news(title, content) values (?,?)")
	checkErr(err)
	rs, err := stmt.Exec(title, content)
	checkErr(err)
	//得到插入的id
	id, err := rs.LastInsertId()
	//可以获得影响行数
	affect, err := rs.RowsAffected()
	fmt.Println("id : %s, rowsAffected : %s", id, affect)
	return affect
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
