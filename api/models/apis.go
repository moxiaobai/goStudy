package models

import (
	"time"
	db "github.com/moxiaobai/goStudy/api/database"
)

type Api struct {
	Id    int        `json:"id" form:"id"`
	Name  string     `json:"name" form:"name" binding:"required"`
	Url   string     `json:"url" form:"url" binding:"required"`
}

//新增数据
func (api Api) Add() (Id int, err error){
	stmt, err := db.SqlDB.Prepare("INSERT INTO apis(name, url, addtime) VALUES (?, ?, ?)")
	if err != nil {
		return
	}

	rs, err := stmt.Exec(api.Name, api.Url, time.Now().Unix())
	if err != nil {
		return
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return
	}

	Id = int(id)
	return
}

//删除数据
func (api Api) Delete() (rows int, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM apis WHERE id=?", api.Id)
	if err != nil {
		return
	}

	row, err := rs.RowsAffected()
	if err != nil {
		return
	}

	rows = int(row)
	return
}

//更新数据
func (api Api) Update() (rows int, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE apis SET name=?, url=? WHERE id=?")
	if err != nil {
		return
	}

	rs, err := stmt.Exec(api.Name, api.Url, api.Id)
	if err != nil {
		return
	}

	row, err := rs.RowsAffected()
	if err != nil {
		return
	}

	rows = int(row)
	defer stmt.Close()
	return
}

//列出API
func (api Api) List(offset, size int) (apis []Api, err error) {
	apis = make([]Api, 0)

	page := (offset - 1) * size
	rows, err := db.SqlDB.Query("select id, name, url from apis order by id ASC limit ?,?", page, size)
	if err != nil {
		return
	}

	for rows.Next() {
		var api Api
		rows.Scan(&api.Id, &api.Name, &api.Url)
		apis = append(apis, api)
	}

	if err = rows.Err(); err != nil {
		return
	}

	defer rows.Close()
	return
}

//统计数据
func Count() (count int, err error) {
	err = db.SqlDB.QueryRow("select COUNT(*) from apis").Scan(&count)
	if err != nil {
		return
	}
	return
}

//查询单条记录
func (a Api) Retrieve() (api Api, err error) {
	err = db.SqlDB.QueryRow("select id, name, url from apis where id = ?", a.Id).Scan(
		&api.Id, &api.Name, &api.Url,
	)

	if err != nil {
		return
	}

	return
}
