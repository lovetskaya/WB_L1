// Реализовать паттерн «адаптер» на любом примере.
// У меня есть различные типы телефонов и адаптер для интеграции их с общей системой продаж

package main

import (
	"fmt"
)

// Phone - интерфейс, который определяет метод для продажи телефона.
type Phone interface {
	Sell() string
}

// OldPhone - структура, представляющая старый телефон.
type OldPhone struct {
	model string
}

func (p *OldPhone) Sell() string {
	return fmt.Sprintf("Selling old phone model: %s", p.model)
}

// NewPhone - структура, представляющая новый телефон.
type NewPhone struct {
	model string
}

func (p *NewPhone) Sell() string {
	return fmt.Sprintf("Selling new phone model: %s", p.model)
}

// SmartPhone - структура, представляющая смартфон, который имеет дополнительную функциональность.
type SmartPhone struct {
	model string
}

func (p *SmartPhone) Sell() string {
	return fmt.Sprintf("Selling smartphone model: %s", p.model)
}

// PhoneAdapter - адаптер, который делает SmartPhone совместимым с интерфейсом Phone.
type PhoneAdapter struct {
	smartPhone *SmartPhone
}

// Sell - метод адаптера, который вызывает метод Sell для SmartPhone.
func (a *PhoneAdapter) Sell() string {
	return a.smartPhone.Sell()
}

// PhoneStore - структура, представляющая магазин телефонов.
type PhoneStore struct {
	phones []Phone
}

// AddPhone - метод для добавления телефона в магазин.
func (s *PhoneStore) AddPhone(phone Phone) {
	s.phones = append(s.phones, phone)
}

// SellPhones - метод для продажи всех телефонов в магазине.
func (s *PhoneStore) SellPhones() {
	for _, phone := range s.phones {
		fmt.Println(phone.Sell())
	}
}

func main() {
	store := &PhoneStore{}

	// Добавляем старые телефоны
	store.AddPhone(&OldPhone{model: "Nokia 3310"})
	store.AddPhone(&OldPhone{model: "Motorola Razr"})

	// Добавляем новые телефоны
	store.AddPhone(&NewPhone{model: "iPhone 14"})
	store.AddPhone(&NewPhone{model: "Samsung Galaxy S21"})

	// Добавляем смартфоны через адаптер
	smartPhone1 := &SmartPhone{model: "Google Pixel 6"}
	adapter1 := &PhoneAdapter{smartPhone: smartPhone1}
	store.AddPhone(adapter1)

	smartPhone2 := &SmartPhone{model: "OnePlus 9"}
	adapter2 := &PhoneAdapter{smartPhone: smartPhone2}
	store.AddPhone(adapter2)

	// Продаем все телефоны в магазине
	store.SellPhones()
}
