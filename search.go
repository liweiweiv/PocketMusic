package main

import (
	//"time"
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/gin-gonic/gin"
	"strconv"
	_ "strconv"
	"time"
)

var index string

type result struct {
	id       int
	mid      int
	mname    string
	singer   string
	lrc      string
	url      string
	createAt time.Time
	updateAt time.Time
}

var songResult = result{}

func HandleSearch(c *gin.Context) {

	db, err := sql.Open("mysql", "admin:testdb123456@tcp(119.29.111.64:3306)/testdb?parseTime=true")
	checkErr(err)
	index = c.Param("Mname_Or_Singer")
	//查询数据
	rows, err := db.Query("SELECT * FROM songs WHERE mname = '" + index + "' OR singer = '" + index + "' ")
	checkErr(err)
	printSearch(rows, c)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printSearch(rows *sql.Rows, c *gin.Context) {

	for rows.Next() {
		var (
			id       int
			mid      int
			mname    string
			singer   string
			lrc      string
			url      string
			createAt time.Time
			updateAt time.Time
		)
		err := rows.Scan(&id, &mid, &mname, &singer, &lrc, &url, &createAt, &updateAt)
		checkErr(err)
		fmt.Println(id)
		songResult.id = id
		fmt.Println(mid)
		songResult.mid = mid
		fmt.Println(mname)
		songResult.mname = mname
		fmt.Println(singer)
		songResult.singer = singer
		fmt.Println(lrc)
		songResult.lrc = lrc
		fmt.Println(url)
		songResult.url = url
		fmt.Println(createAt)
		songResult.createAt = createAt
		fmt.Println(updateAt)
		songResult.updateAt = updateAt
		c.String(200, strconv.Itoa(songResult.id)+"%"+strconv.Itoa(songResult.mid)+"%"+songResult.mname+"%"+songResult.singer+"%"+songResult.url+"%"+songResult.lrc+"\n")
	}
}
