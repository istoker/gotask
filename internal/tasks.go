package task

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const path = "/Users/idzstoke/.gotasklist"

type Task struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
	Id        int    `json:"id"`
}

func taskfileExists() bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetTaskList() ([]Task, error) {
	var list []Task
	if taskfileExists() {
		raw, err := ioutil.ReadFile(path)
		if err != nil {
			log.Print(err.Error())
			return nil, err
		} else if len(raw) == 0 {
			return list, nil
		}
		err = json.Unmarshal(raw, &list)
		if err != nil {
			log.Print(err.Error())
			return nil, err
		}
	}
	return list, nil
}

func toJson(t interface{}) []byte {
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Print(err.Error())
		return make([]byte, 0, 0)
	}
	return bytes
}

func SaveTaskList(list []Task) error {
	f, err := os.Create(path)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	defer f.Close()
	jsonTasks := toJson(list)
	_, err = f.Write(jsonTasks)
	return err
}
