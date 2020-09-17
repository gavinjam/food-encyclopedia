package main

import (
	"database/sql"
)

type Store interface {
	CreateFood(food *Food) error
	GetFood() ([]*Food, error)
	DeleteFood(deleteIds string) error
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateFood(food *Food) error {
	_, err := store.db.Query("INSERT INTO food(species, description) VALUES ($1, $2)", food.Species, food.Description)
	return err
}

func (store *dbStore) GetFood() ([]*Food, error) {
	rows, err := store.db.Query("SELECT id, species, description from food")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	food := []*Food{}
	for rows.Next() {
		singleFood := &Food{}

		if err := rows.Scan(&singleFood.Id, &singleFood.Species, &singleFood.Description); err != nil {
			return nil, err
		}

		food = append(food, singleFood)
	}
	return food, nil
}

func (store *dbStore) DeleteFood(deleteIds string) error {

	_, err := store.db.Query("DELETE FROM food WHERE id IN ($1)", deleteIds)

	/*
	if err != nil {
		return err
	}
	defer rows.Close()

	
	for rows.Next() {
	}*/

	return err
}

var store Store

func InitStore(s Store) {
	store = s
}
