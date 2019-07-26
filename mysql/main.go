package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "package/log"
)

var (
	sqlparam = "root:123456@/dnakit?charset=utf8"
	mlog     *log.Logger
)

func main() {
	var err error
	mlog, err = log.New("", "DEBUG")
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlclient, err := sql.Open("mysql", sqlparam)
	if err != nil {
		mlog.Error("error open mysql,%s", err.Error())
		return
	}
	defer sqlclient.Close()
	err = sqlclient.Ping()
	if err != nil {
		mlog.Errorln(err)
		return
	}

	//var input string = "00000000000000000000000090650000"
	var input string = "1;drop table dnakit_uid;--"
	//sqlstr := fmt.Sprintf("select pid from dnakit_pid where pid= %s", input)

	sqlstr := "select pid from dnakit_pid where pid= ?"
	smtp, err := sqlclient.Prepare(sqlstr)
	if err != nil {
		mlog.Errorln(err)
		return
	}
	defer smtp.Close()

	rows, err := smtp.Query(input)
	if err != nil {
		mlog.Error("%s", err.Error())
		return
	}
	defer rows.Close()
	var vs string
	for rows.Next() {
		err = rows.Scan(&vs)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		fmt.Printf("----:%s\n", vs)
	}
	fmt.Printf("end\n")
}

func createtable() {
	sqlclient, err := sql.Open("mysql", sqlparam)
	if err != nil {
		fmt.Errorf("error open mysql,%s\n", err.Error())
		return
	}
	defer sqlclient.Close()
	err = sqlclient.Ping()
	if err != nil {
		return
	}
	sqlstr := fmt.Sprintf("create database if not exists dnakit_test default character set utf8")
	_, err = sqlclient.Exec(sqlstr)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
}
