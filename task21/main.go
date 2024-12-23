package main

import "fmt"

// Target представляет собой нужный нам интерфейс, который должен реализовывать адаптер
type Target interface {
	GetReport()
}

// Database наша исходная структура, которая не обладает Target функциональностью
type Database struct {
}

func (d *Database) FindAllUsers() []int {
	return []int{1, 2, 3, 4, 5}
}

func (d *Database) FindName(id int) string {
	names := map[int]string{
		1: "Vasya",
		2: "Petya",
		3: "Maksim",
		4: "Igor",
		5: "Pavel",
	}
	return names[id]
}

type DatabaseAdapter struct {
	*Database
}

// GetReport реализует целевой интерфейс для взаимодействия с использованием уже существующей функциональностью
func (d *DatabaseAdapter) GetReport() {
	users := d.FindAllUsers()
	for _, id := range users {
		fmt.Printf("User id: %d, name: %s\n", id, d.FindName(id))
	}
}

func main() {
	db := &Database{}

	adapter := DatabaseAdapter{db}

	adapter.GetReport()
}
