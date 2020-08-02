package service

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

const queryInsert = `INSERT INTO public.notifications(id, message) VALUES ($1,$2)`

//Notifications{} contains pointer to database
type Notifications struct {
	db *sql.DB
}

// NewNotifications() returns Notifications with db
func NewNotifications(data *sql.DB) *Notifications {
	return &Notifications{db: data}
}

// Add() adds new user to database
func (n *Notifications) Add(message string) error {
	ui, err := uuid.NewRandom()

	if err != nil {
		return fmt.Errorf("uuid.NewRandom error: %w", err)
	}

	_, err = n.db.Exec(queryInsert, ui, message)

	return err
}
