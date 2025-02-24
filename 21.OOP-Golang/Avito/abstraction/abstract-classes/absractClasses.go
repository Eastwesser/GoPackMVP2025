package main

import "fmt"

// Абстракция: интерфейс Task
type Task interface {
	Execute()
	GetDetails() string
}

// #####################################################################################################################

// Структура для простой задачи
type SimpleTask struct {
	Name        string
	Description string
}

func (s *SimpleTask) Execute() {
	fmt.Println("Выполнена простая задача:", s.Name)
}

func (s *SimpleTask) GetDetails() string {
	return fmt.Sprintf("Простая задача: %s, Описание: %s", s.Name, s.Description)
}

// #####################################################################################################################

// Структура для задачи с дедлайном
type DeadlineTask struct {
	Name        string
	Description string
	Deadline    string
}

func (d *DeadlineTask) Execute() {
	fmt.Println("Выполнена задача с дедлайном:", d.Name)
}

func (d *DeadlineTask) GetDetails() string {
	return fmt.Sprintf("Задача с дедлайном: %s, Описание: %s, Дедлайн: %s", d.Name, d.Description, d.Deadline)
}

// #####################################################################################################################

func main() {
	var task Task

	task = &SimpleTask{
		Name:        "Написать отчет",
		Description: "Подготовить отчет по итогам квартала",
	}
	task.Execute()
	fmt.Println(task.GetDetails())

	task = &DeadlineTask{
		Name:        "Подготовить презентацию",
		Description: "Создать презентацию для совещания",
		Deadline:    "2023-10-31",
	}
	task.Execute()
	fmt.Println(task.GetDetails())
}
