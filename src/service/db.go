package main

import (
	"database/sql"
	_ "engine/pq"
	"sync"
)

var (
	mutex sync.Mutex
	db    *sql.DB
)

const dsn = `host=127.0.0.1 port=4932 dbname=test user=viney password=admin sslmode=disable`

func init() {
	mutex.Lock()
	defer mutex.Unlock()

	pqdb, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("sql.Open: " + err.Error())
	}

	db = pqdb
}

// -------------
// Close

func Close() error {
	return db.Close()
}

// ------------
// Count

func Count() (uint, error) {
	row := db.QueryRow(countSql)

	var count uint
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

const countSql = `
    select count(*) from user_profile;
`

// -----------
// FindAll

type User struct {
	Uid  int
	Name string
	Age  int
}

func FindAll(currentPage, lineSize uint) ([]User, error) {
	rows, err := db.Query(findAllSql, currentPage, lineSize)
	if err != nil {
		return nil, err
	}

	var (
		user  = &User{}
		users = []User{}
	)
	for rows.Next() {
		defer rows.Close()
		if err := rows.Scan(
			&user.Uid,
			&user.Name,
			&user.Age,
		); err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	return users, nil
}

const findAllSql = `
    select uid,name,age from user_profile order by uid offset $1 limit $2;
`
