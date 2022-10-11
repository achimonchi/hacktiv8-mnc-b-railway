package repository

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Contact   string
	CreatedAt time.Time
}

type Product struct {
	ID   int
	Name string
}

type Person struct {
	ID int
}

type Order struct {
	Id uint
}

type Item struct {
	Id      int
	OrderId uint
}
