package main

import pgx "github.com/jackc/pgx/v5"

func main() {
	pgx.Connect(nil, "")
}
