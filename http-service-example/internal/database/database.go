package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DataAdapter struct {
	DbConnectionPool *sql.DB
}

type Identity struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age int `json:"age"`
}

func NewDataAdapter() (*DataAdapter, error) {
	db, err := sql.Open("mysql", "root:@/customer")
	if err != nil {
		return nil, err
	}
	// since sqlOpen does not introduce any connection errors, use Ping() to check
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DataAdapter{
		DbConnectionPool: db, 
	}, nil
}

func (dataAdapter *DataAdapter) GetUserByUsername(username string) (*Identity, error) {
	statement, err := dataAdapter.DbConnectionPool.Prepare("select * from identity where username=?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	
	var identity Identity
	err = statement.QueryRow(username).Scan(&identity.Username, &identity.Password, &identity.Age)
	if err != nil {
		return nil, err
	}

	return &identity, nil
}

func (dataAdapter *DataAdapter) GetUserByUsernameAndPassword(username string, password string) (*Identity, error) {
	statement, err := dataAdapter.DbConnectionPool.Prepare("select * from identity where username=? and password=?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	
	var identity Identity
	err = statement.QueryRow(username, password).Scan(&identity.Username, &identity.Password, &identity.Age)
	if err != nil {
		return nil, err
	}

	return &identity, nil
}