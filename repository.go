package main

import (
	"os"
	"path"
	"time"

	"github.com/ostafen/clover"
)

const colName = "tasks"
const dbName = "task-db"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Time        string    `json:"time"`
	Reminder    time.Time `json:"reminder"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskRepository struct{}

func getDb() (*clover.DB, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	db, err := clover.Open(path.Join(dir, dbName))
	if err != nil {
		return nil, err
	}

	isExists, err := db.HasCollection(colName)
	if err != nil {
		return nil, err
	}
	if !isExists {
		err = db.CreateCollection(colName)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func (tr *TaskRepository) GetTasks() ([]Task, error) {
	var tasks []Task

	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	docs, err := db.Query(colName).FindAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		task := &Task{}
		err = doc.Unmarshal(task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *task)
	}

	return tasks, nil
}

func (tr *TaskRepository) StoreTask(t Task) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()

	//TODO: change this get task implementation using better performance
	tasks, err := tr.GetTasks()
	if err != nil {
		return err
	}
	nextId := 1
	if len(tasks) > 0 {
		nextId = tasks[len(tasks)-1].ID + 1
	}

	doc := clover.NewDocument()
	doc.Set("id", nextId)
	doc.Set("description", t.Description)
	doc.Set("time", t.Time)
	doc.Set("created_at", time.Now())

	err = db.Insert(colName, doc)
	if err != nil {
		return err
	}

	return nil
}
