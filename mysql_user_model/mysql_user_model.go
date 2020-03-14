package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "log"

const CONN_HOST = "localhost"
const CONN_PORT = "3306"
const DRIVER_NAME = "mysql"
const DATA_SOURCE_NAME = ""
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

func testCreate() {
    username := "jiao"
    password := "1111"
    model := &MysqlUserModel{}
    model.Create(username, password)
    fmt.Println(model.id)
}

func testFromUsername() {
    username := "jiaoo"
    model := &MysqlUserModel{}
    err := model.FromUsername(username)
    fmt.Println(err)

    username = "jiao"
    model = &MysqlUserModel{}
    err = model.FromUsername(username)
    fmt.Println(model.GetId())
    fmt.Println(model.GetUsername())
    fmt.Println(model.GetNickname())
    fmt.Println(model.GetPassword())
    fmt.Println(model.GetAvatar())
}

func testGet() {
    userId := 2
    model := &MysqlUserModel{}
    err := model.Get(userId)
    fmt.Println(err)

    userId = 5
    model = &MysqlUserModel{}
    err = model.Get(userId)
    fmt.Println(model.GetId())
    fmt.Println(model.GetUsername())
    fmt.Println(model.GetNickname())
    fmt.Println(model.GetPassword())
    fmt.Println(model.GetAvatar())
}

func testSetNicknameError(){
    userId := 2
    model := &MysqlUserModel{}
    model.Get(userId)

    nickname := "lalal"
    model.SetNickname(nickname)

    model = &MysqlUserModel{}
    model.Get(userId)
    fmt.Println(model.GetNickname())

}

func testSetNickname(){
    userId := 5
    model := &MysqlUserModel{}
    model.Get(userId)

    nickname := "lalalsss"
    model.SetNickname(nickname)

    model = &MysqlUserModel{}
    model.Get(userId)
    fmt.Println(model.GetNickname())
}

func testSetAvatar(){
    userId := 5
    model := &MysqlUserModel{}
    model.Get(userId)

    avatar := []byte{'a', 'b'}
    model.SetAvatar(avatar)

    model = &MysqlUserModel{}
    model.Get(userId)
    fmt.Println(model.GetAvatar())
}

func main() {
    // testCreate()
    // testFromUsername()
    // testGet()
    // testSetNicknameError()
    // testSetNickname()
    testSetAvatar()
}
