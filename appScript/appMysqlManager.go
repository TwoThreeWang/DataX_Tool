package appScript

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

type MysqlManager struct {
	db *sql.DB
}

func NewMysqlManager(jdbcURL string) (*MysqlManager, error) {
	db, err := sql.Open("mysql", jdbcURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &MysqlManager{db: db}, nil
}

func GetTns(jdbcURL, database string) string {
	fmt.Println(jdbcURL)
	// 正则表达式匹配 IP 地址和端口号
	ipPortRegex := `^jdbc:mysql://(\d+\.\d+\.\d+\.\d+):(\d+).*`
	ipPortMatches := regexp.MustCompile(ipPortRegex).FindStringSubmatch(jdbcURL)
	ipAddress := ipPortMatches[1]
	port := ipPortMatches[2]
	// 正则表达式匹配用户名和密码
	credentialRegex := `&username=([^&]+)&password=([^&]+)`
	credentialMatches := regexp.MustCompile(credentialRegex).FindStringSubmatch(jdbcURL)
	username := credentialMatches[1]
	password := credentialMatches[2]
	jdbcURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, ipAddress, port, database)
	return jdbcURL
}

func (mgr *MysqlManager) Close() error {
	return mgr.db.Close()
}

func (mgr *MysqlManager) Insert(query string, args ...interface{}) (int64, error) {
	result, err := mgr.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (mgr *MysqlManager) Update(query string, args ...interface{}) (int64, error) {
	result, err := mgr.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (mgr *MysqlManager) Delete(query string, args ...interface{}) (int64, error) {
	result, err := mgr.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (mgr *MysqlManager) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return mgr.db.Query(query, args...)
}
