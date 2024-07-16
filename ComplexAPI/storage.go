package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
	CreateAccount(*Account) (int, error)
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=123456 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}
	
	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) CreateAccount(account *Account) (int, error) {
	query := `
	INSERT INTO account (first_name, last_name, number, balance, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING (id)
	`

	var id int

	err := s.db.QueryRow(
		query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt,
	).Scan(&id)

	if err != nil {
		return -1, err
	}

	fmt.Printf("%+v\n", id)

	return id, nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM account WHERE id = $1", id)

	return err
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)

	if err != nil {
		return nil, err
	}

	return account, err
}
