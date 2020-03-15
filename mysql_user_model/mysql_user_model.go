package mysql_user_model

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

const CONN_HOST = "localhost"
const CONN_PORT = "3306"
const DRIVER_NAME = "mysql"
const DATA_SOURCE_NAME = "root:Ms!(**0212@/test"
var Db *sql.DB
var connectionError error

func init() {
    Db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
    if connectionError != nil {
        log.Fatal("error connecting to database: ", connectionError)
    }
}

type MysqlUserModel struct {
    id int
    username string
    password string
    nickname string
    avatar []byte
}

func (m *MysqlUserModel) Create(username, password string) error {
    m.username = username
    m.password = password
    statement := "INSERT into user (username, password, nickname) values (?, ?, ?)"
    stmt, err := Db.Prepare(statement)
    defer stmt.Close()
    if err != nil {
        return err // prepare error
    }
    res, err := stmt.Exec(username, password, "")
    // Set nickname to "" because scan cannot convert null to string
    if err != nil {
        return err // user exists error
    }
    id, err := res.LastInsertId()
    m.id = int(id)

    return nil
}

func (m *MysqlUserModel) FromUsername(username string) (err error) {
    statement := "select id, password, nickname, avatar from user where username = (?)"
    m.username = username
    err = Db.QueryRow(statement, username).Scan(&m.id, &m.password, &m.nickname, &m.avatar)
    return // not found
}

func (m *MysqlUserModel) Get(userId int) (err error) {
    statement := "select username, password, nickname, avatar from user where id = (?)"
    err = Db.QueryRow(statement, userId).Scan(&m.username, &m.password, &m.nickname, &m.avatar)
    if err != nil {
        return
    }
    m.id = userId
    return // not found
}

func (m *MysqlUserModel) GetPassword() string {
    return m.password
}

func (m *MysqlUserModel) GetId() int {
    return m.id
}

func (m *MysqlUserModel) GetUsername() string {
    return m.username
}

func (m *MysqlUserModel) GetNickname() string {
    return m.nickname
}

func (m *MysqlUserModel) GetAvatar() []byte {
    return m.avatar
}

func (m *MysqlUserModel) SetNickname(nickname string) (err error) {
    statement := "update user set nickname = ? where id = ?"
    m.nickname = nickname
    _, err = Db.Exec(statement, nickname, m.id)
    return err
}

func (m *MysqlUserModel) SetAvatar(avatar []byte) (err error) {
    statement := "update user set avatar = ? where id = ?"
    m.avatar = avatar
    _, err = Db.Exec(statement, avatar, m.id)
    return err
}

func (m *MysqlUserModel) Delete() (err error) {
    statement := "delete from user where id = ?"
    _, err = Db.Exec(statement, m.id)
    return
}
