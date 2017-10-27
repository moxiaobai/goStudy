package models

import (
	"log"
	"time"
	db "github.com/moxiaobai/goStudy/api/database"
)

type Api struct {
	Id    int        `json:"id" form:"id"`
	Name  string     `json:"name" form:"name" binding:"required"`
	Url   string     `json:"url" form:"url" binding:"required"`
}

//新增数据
func (a Api) Add() (Id int, err error){
	rs, err := db.SqlDB.Exec("INSERT INTO apis(name, url, addtime) VALUES (?, ?, ?)", a.Name, a.Url, time.Now().Unix())
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
func (a Api) Delete() (rows int, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM apis WHERE id=?", a.Id)
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
func (a Api) Update() (rows int, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE apis SET name=?, url=? WHERE id=?")
	if err != nil {
		return
	}

	rs, err := stmt.Exec(a.Name, a.Url, a.Id)
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
func (a Api) List(offset, size int) (apis []Api, err error) {
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
