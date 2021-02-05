package cstools

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var GlobalMysqlClient MysqlClient

type MysqlClient struct {
	Host     string
	Port     string
	UserName string
	Password string
	Db       string
	Conn     *xorm.Engine
}

func (mc *MysqlClient) InitConn() (err error) {
	mc.Conn, err = xorm.NewEngine("mysql", mc.UserName+":"+mc.Password+"@tcp("+mc.Host+":"+mc.Port+")/"+mc.Db+"?charset=utf8")
	return err
}

func (mc *MysqlClient) QuerySql(SqlString string, args ...interface{}) *xorm.Session {
	if len(args) > 0 {
		return mc.Conn.SQL(SqlString, args...)
	} else {
		return mc.Conn.SQL(SqlString)
	}

}
