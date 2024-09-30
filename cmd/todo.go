package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)



type Item struct {
	Task string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todos []Item

func (t *Todos) AddTask(task string) {
	item := Item{
	 Task : task,
	 Completed: false,
	 CreatedAt: time.Now(),
	 UpdatedAt: time.Time{},		

	}

	*t  = append(*t, item)
}

func (t *Todos) DoneTask(index int) error {
	ls := *t
	if index < 0 || index >= len(ls) {
		return errors.New("invalid index")
	}
	ls[index].Completed = true
	ls[index].UpdatedAt = time.Now()

	return nil
}


func (t *Todos) DeleteTask(index int) error {
	ls := *t
	if index < 0 || index >= len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index], ls[index+1:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
	 file, err := os.ReadFile(filename)
        if  err != nil {
           return err
	}

	if len(file) == 0 {
         return errors.New("any task does not exit")
	}
	
	err = json.Unmarshal(file ,t)
	if err != nil {
		return err
	}

	return nil
	
}