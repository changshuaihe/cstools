package cstools

import (
	"fmt"
	"github.com/amorist/mango"
)

type Mgo struct {
	DB          string
	mgoSession  *mango.Session
	ThreadCtl   Thread
	ThreadCount int
	ServerIp    string
	Port        string
}

type MgoGlobalConf struct {
	ServerIp    string
	Port        string
	Password    string
	ThreadCount int
}

//使用前初始化
//dataPoolMgo := lib.Mgo{}
//dataPoolMgo.InitMgo("data")
//SyncGitPro(&dataPoolMgo)

var GlobalMgoConf MgoGlobalConf

func (m *Mgo) InitMgo(dbName string) {
	m.ThreadCtl.Init(GlobalMgoConf.ThreadCount)
	m.DB = dbName
	m.mgoSession = mango.New("mongodb://" + GlobalMgoConf.ServerIp + ":" + GlobalMgoConf.Port)
	m.mgoSession.SetPoolLimit(1000)

	if err := m.mgoSession.Connect(); err != nil {
		fmt.Println(err)
	}
}

func (m *Mgo) MgoInsertOne(col string, data interface{}) error {
	m.ThreadCtl.Start()
	err := m.mgoSession.DB(m.DB).Collection(col).Insert(data)
	m.ThreadCtl.End()
	return err
}

func (m *Mgo) MgoInsertMany(col string, datas []interface{}) error {
	m.ThreadCtl.Start()
	err := m.mgoSession.DB(m.DB).Collection(col).InsertAll(datas)

	m.ThreadCtl.End()
	return err
}

func (m *Mgo) MgoRemoveAll(col string, condition map[string]interface{}) error {
	m.ThreadCtl.Start()
	err := m.mgoSession.DB(m.DB).Collection(col).RemoveAll(condition)
	m.ThreadCtl.End()
	return err
}

func (m *Mgo) MgoGetAll(col string, condition map[string]interface{}) ([]map[string]interface{}, error) {
	m.ThreadCtl.Start()
	res := []map[string]interface{}{}
	err := m.mgoSession.DB(m.DB).Collection(col).Find(condition).All(&res)
	if err != nil {
		fmt.Println(err)
	}
	m.ThreadCtl.End()
	return res, err
}

func (m *Mgo) MgoGetList(col string, condition map[string]interface{}, page int64, size int64) ([]map[string]interface{}, error) {
	m.ThreadCtl.Start()
	res := []map[string]interface{}{}
	page = page - 1
	if page < 0 {
		page = 0
	}
	err := m.mgoSession.DB(m.DB).Collection(col).Find(condition).Limit(size).Skip(size * page).All(&res)
	m.ThreadCtl.End()
	return res, err
}

func (m *Mgo) MgoGetOne(col string, condition interface{}) (map[string]interface{}, error) {
	m.ThreadCtl.Start()
	res := make(map[string]interface{})
	err := m.mgoSession.DB(m.DB).Collection(col).Find(condition).One(&res)
	m.ThreadCtl.End()
	return res, err
}

func (m *Mgo) MgoUpdateOne(col string, condition interface{}, data interface{}, upset bool) error {
	m.ThreadCtl.Start()
	updateData := make(map[string]interface{})
	updateData["$set"] = data

	err := m.mgoSession.DB(m.DB).Collection(col).Update(condition, updateData, upset)
	m.ThreadCtl.End()
	return err
}

func (m *Mgo) MgoUpdateMany(col string, condition interface{}, data interface{}, upset bool) error {
	m.ThreadCtl.Start()
	updateData := make(map[string]interface{})
	updateData["$set"] = data
	_, err := m.mgoSession.DB(m.DB).Collection(col).UpdateAll(condition, updateData, upset)
	if err != nil {
		fmt.Println(err)
	}
	m.ThreadCtl.End()
	return err
}
