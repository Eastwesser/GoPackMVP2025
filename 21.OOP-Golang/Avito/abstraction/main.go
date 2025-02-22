package main

import "fmt"

//Абстракция позволяет скрыть сложность и показать только необходимые детали.
//В Avito это можно представить как общий интерфейс для всех объявлений, которые могут быть опубликованы.

// Абстракция: интерфейс Advertisement
type Advertisement interface {
	Publish()
	GetDetails() string
}

// #####################################################################################################################

// Структура для объявления о продаже
type SaleAd struct {
	Title string
	Price float64
}

func (s *SaleAd) Publish() {
	fmt.Println("Опубликовано объявление о продаже:", s.Title)
}

func (s *SaleAd) GetDetails() string {
	return fmt.Sprintf("Продажа: %s, Цена: %.2f руб.", s.Title, s.Price)
}

// #####################################################################################################################

// Структура для объявления об аренде
type RentAd struct {
	Title  string
	Price  float64
	Period string
}

func (r *RentAd) Publish() {
	fmt.Println("Опубликовано объявление об аренде:", r.Title)
}

func (r *RentAd) GetDetails() string {
	return fmt.Sprintf("Аренда: %s, Цена: %.2f руб./%s", r.Title, r.Price, r.Period)
}

// #####################################################################################################################

func main() {
	var ad Advertisement

	ad = &SaleAd{Title: "Продам iPhone 13", Price: 70000}
	ad.Publish()
	fmt.Println(ad.GetDetails())

	ad = &RentAd{Title: "Сдам квартиру", Price: 30000, Period: "месяц"}
	ad.Publish()
	fmt.Println(ad.GetDetails())
}
