package cstools

import (
	"fmt"
	"github.com/amorist/mango"
)

var DB string

type Mgo struct {
	S *mango.Session
}

func (m *Mgo) InitMgo() {
	DB = "plat"
	m.S = mango.New("mongodb://127.0.0.1")
	m.S.SetPoolLimit(10)

	if err := m.S.Connect(); err != nil {
		fmt.Println(err)
	}
}

func (m *Mgo) MgoInsertOne(col string, data interface{}) {
	err := m.S.DB(DB).Collection(col).Insert(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (m *Mgo) MgoInsertMany(col string, datas []interface{}) {
	err := m.S.DB(DB).Collection(col).InsertAll(datas)
	if err != nil {
		fmt.Println(err)
	}
}
