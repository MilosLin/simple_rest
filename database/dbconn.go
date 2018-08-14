package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"simple_rest/config"

	// blank import
	_ "github.com/go-sql-driver/mysql"
	//"log"
)

var (
	//資料庫連線物件
	pool map[string]*sql.DB

	//同步鎖
	mu *sync.Mutex
)

func init() {
	pool = make(map[string]*sql.DB)
	mu = &sync.Mutex{}
}

// GetConn : 依照資料庫名稱取得DB連線
func GetConn(dbName string) *sql.DB {
	mu.Lock()
	defer mu.Unlock()

	if conn, ok := pool[dbName]; ok {
		if err := conn.Ping(); err == nil {
			return conn
		}
		conn.Close()
	}

	pool[dbName] = createConn(dbName)
	return pool[dbName]
}

// createConn : 建立資料庫連線
func createConn(dbName string) *sql.DB {
	c := config.Forge()
	group := "Database." + dbName

	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=%s&timeout=%s",
		c.GetString(group+".Account"),
		c.GetString(group+".Password"),
		c.GetString(group+".IP"),
		c.GetString(group+".Port"),
		c.GetString(group+".DB"),
		c.GetString("Database.Loc"),
		c.GetString("Database.Timeout"),
	)
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		log.Fatalf("Connect Database Failed. err:%s , connStr:%s", err.Error(), connStr)
		return nil
	}

	db.SetMaxOpenConns(c.GetInt("Database.MaxOpenConns"))
	db.SetMaxIdleConns(c.GetInt("Database.MaxIdleConns"))
	db.SetConnMaxLifetime(c.GetDuration("Database.ConnMaxLifeTime"))
	return db
}

// GetConn2 : 依照資料庫名稱取得DB連線,可由外部代入db 名稱
func GetConn2(connection, dbName string) (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	index := fmt.Sprintf("%s_%s", connection, dbName)
	if conn, ok := pool[index]; ok {
		if err := conn.Ping(); err == nil {
			return conn, nil
		}

		conn.Close()
	}

	db, err := createConn2(connection, dbName)

	if err == nil {
		pool[index] = db
		return db, nil
	}

	return db, err
}

// createConn2 : 新版建立連線方式，資料庫名稱由外部代入
func createConn2(connection, dbName string) (db *sql.DB, err error) {
	c := config.Forge()
	group := "Database." + connection

	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=%s&timeout=%s",
		c.GetString(group+".Account"),
		c.GetString(group+".Password"),
		c.GetString(group+".IP"),
		c.GetString(group+".Port"),
		dbName,
		c.GetString("Database.Loc"),
		c.GetString("Database.Timeout"),
	)
	db, err = sql.Open("mysql", connStr)

	if err != nil {
		log.Fatalf("Connect Database Failed. err:%s , connStr:%s", err.Error(), connStr)
		return db, err
	}

	db.SetMaxOpenConns(c.GetInt("Database.MaxOpenConns"))
	db.SetMaxIdleConns(c.GetInt("Database.MaxIdleConns"))
	db.SetConnMaxLifetime(c.GetDuration("Database.ConnMaxLifeTime"))
	return
}

// CloseConn : 關閉所有連線
func CloseConn() {
	mu.Lock()
	defer mu.Unlock()

	for _, v := range pool {
		v.Close()
	}
}
