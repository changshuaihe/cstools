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

func (mc *MysqlClient) QuerySql(SqlString string, args ...interface{}) ([]map[string]string, error) {
	if len(args) > 0 {
		return mc.Conn.SQL(SqlString, args...).QueryString()
	} else {
		return mc.Conn.SQL(SqlString).QueryString()
	}

}
