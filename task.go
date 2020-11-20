package cstools

import (
	"fmt"
	"sync"
)

type TaskItem struct {
	name   string `json:"name" bson:"name"`
	Value  string `json:"value" bson:"value"`
	Static string `json:"static" bson:"static"`
}

type TaskManager struct {
	name            string     `json:"name"`
	Items           []TaskItem `json:"items"`
	tmpAllValue     map[string]interface{}
	tmpAllValueLock *sync.RWMutex
}

var TaskMgo Mgo

var AllTask []TaskItem

func (t *TaskManager) InitTask(name string) {
	TaskMgo.InitMgo("task")
	t.name = name
	t.tmpAllValue = make(map[string]interface{})
	t.tmpAllValueLock = new(sync.RWMutex)

	fmt.Println("开始恢复任务状态")
	t.Items = t.GetAllTask()
	fmt.Println("任务恢复完成")
}

func (t *TaskManager) RecordTask(taskItem TaskItem) {
	taskItem.name = t.name
	t.tmpAllValueLock.Lock()
	t.tmpAllValue[taskItem.Value] = ""
	t.tmpAllValueLock.Unlock()

	c := make(map[string]interface{})
	c["value"] = taskItem.Value
	c["name"] = taskItem.name
	go TaskMgo.MgoUpdateOne("task", c, taskItem, true)
}

func (t *TaskManager) CleanTask() {
	c := make(map[string]interface{})
	c["name"] = t.name
	TaskMgo.MgoRemoveAll("task", c)
}

func (t *TaskManager) GetAllTask() []TaskItem {
	tasks := []TaskItem{}
	c := make(map[string]interface{})
	c["name"] = t.name
	allTasks, _ := TaskMgo.MgoGetAll("task", c)
	for _, item := range allTasks {
		taskItem := TaskItem{}
		taskItem.name = t.name
		taskItem.Value = item["value"].(string)
		taskItem.Static = item["static"].(string)
		t.Items = append(t.Items, taskItem)
		//加入临时map，方便搜索任务
		t.tmpAllValue[taskItem.Value] = ""
	}
	return tasks
}

func (t *TaskManager) HasTask(value string) bool {
	t.tmpAllValueLock.Lock()
	defer t.tmpAllValueLock.Unlock()
	return HasKey(t.tmpAllValue, value)
}
