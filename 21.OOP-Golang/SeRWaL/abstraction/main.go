package main

import "fmt"

//Абстракция позволяет скрыть сложность и показать только необходимые детали.
//В SeRWaL это можно представить как общий интерфейс для всех учеников, которые могут развивать свои навыки.

// Абстракция: интерфейс Student
type Student interface {
	PracticeSpeaking()
	PracticeReading()
	PracticeWriting()
	PracticeListening()
}

// Структура для ученика начальных классов
type ElementaryStudent struct {
	Name string
}

func (e ElementaryStudent) PracticeSpeaking() {
	fmt.Println(e.Name, "практикует разговорный английский с простыми фразами.")
}

func (e ElementaryStudent) PracticeReading() {
	fmt.Println(e.Name, "читает простые тексты и учит базовые слова.")
}

func (e ElementaryStudent) PracticeWriting() {
	fmt.Println(e.Name, "практикует письменный английский с простыми предложениями.")
}

func (e ElementaryStudent) PracticeListening() {
	fmt.Println(e.Name, "смотрит простые мультики с субтитрами и учит базовые слова.")
}

// Структура для ученика старших классов
type HighSchoolStudent struct {
	Name string
}

func (h HighSchoolStudent) PracticeSpeaking() {
	fmt.Println(h.Name, "практикует сложные диалоги и обсуждает актуальные темы.")
}

func (h HighSchoolStudent) PracticeReading() {
	fmt.Println(h.Name, "читает статьи и книги на английском.")
}

func (h HighSchoolStudent) PracticeWriting() {
	fmt.Println(h.Name, "практикует письменный английский со сложными предложениями.")
}

func (h HighSchoolStudent) PracticeListening() {
	fmt.Println(h.Name, "смотрит сериалы в оригинале, учит сложные слова и выражения.")
}

func main() {
	var student Student

	student = ElementaryStudent{Name: "Миша"}
	student.PracticeSpeaking()
	student.PracticeReading()
	student.PracticeWriting()
	student.PracticeListening()

	student = HighSchoolStudent{Name: "Егор"}
	student.PracticeSpeaking()
	student.PracticeReading()
	student.PracticeWriting()
	student.PracticeListening()
}
