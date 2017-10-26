package models

import (
	"log"
	"time"
)

type Api struct {
	Id    int        `json:"id" form:"id"`
	Name  string     `json:"name" form:"name"`
	Url   string     `json:"url" form:"url"`
}

//新增数据
func (db *DbInstance) AddApis(api Api) (insertId int64) {
	rs, err := db.Db.Exec("INSERT INTO apis(name, url, addtime) VALUES (?, ?, ?)", api.Name, api.Url, time.Now().Unix())
	if err != nil {
		log.Fatalln(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return id
}

//删除数据
func (db *DbInstance) DeleteApi(id int) (count int64) {
	rs, err := db.Db.Exec("DELETE FROM apis WHERE id=?", id)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	return ra
}

//更新数据
func (db *DbInstance) UpdateApi(id int, api Api) (count int64) {
	stmt, err := db.Db.Prepare("UPDATE apis SET name=?, url=? WHERE id=?")
	defer stmt.Close()

	if err != nil {
		log.Fatalln(err)
	}

	rs, err := stmt.Exec(api.Name, api.Url, id)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	return ra
}

//列出API
func (db *DbInstance) ListApi() (data *[]Api) {
	rows, err := db.Db.Query("select id, name, url from apis order by id ASC")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	apis := make([]Api, 0)
	for rows.Next() {
		var api Api
		rows.Scan(&api.Id, &api.Name, &api.Url)
		apis = append(apis, api)
	}

	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	defer db.Db.Close()

	return &apis
}

//查询单条记录
func (db *DbInstance) RetrieveApi(id int) (row Api) {
	var api Api
	err := db.Db.QueryRow("select id, name, url from apis where id = ?", id).Scan(
		&api.Id, &api.Name, &api.Url,
	)

	if err != nil {
		log.Println(err)
	}
	return api

}
