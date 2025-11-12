package main

import (
	"fmt"
)

type Ticket struct {
	ID            int
	PassengerName string
	Destination   string
}

type BookingSystem struct {
	ticket map[int]Ticket
}

func (bs *BookingSystem) BookTicket(id int, name, destination string) error {
	if _, exists := bs.ticket[id]; exists {
		return fmt.Errorf("билет с ID %d уже существует", id)
	}
	// Добавляем билет если он не существует
	bs.ticket[id] = Ticket{
		ID:            id,
		PassengerName: name,
		Destination:   destination,
	}
	return nil
}

func (bs *BookingSystem) GetTicket(id int) (Ticket, error) {
	ticket, exists := bs.ticket[id]
	if !exists {
		return Ticket{}, fmt.Errorf("билет с ID %d не найден", id)
	}
	return ticket, nil
}

func (bs *BookingSystem) CancelTicket(id int) error {
	if _, exists := bs.ticket[id]; !exists {
		return fmt.Errorf("билет с ID %d не найден", id)
	}
	delete(bs.ticket, id)
	return nil
}

func main() {
	system := &BookingSystem{ticket: make(map[int]Ticket)}

	err := system.BookTicket(1, "Иван", "Москва")
	if err != nil {
		fmt.Println("Ошибка бронирования:", err)
		return
	}

	ticket, err := system.GetTicket(1)
	if err != nil {
		fmt.Println("Ошибка получения:", err)
		return
	}

	fmt.Printf("Билет: ID=%d, Пассажир=%s, Направление=%s\n",
		ticket.ID, ticket.PassengerName, ticket.Destination)
}
