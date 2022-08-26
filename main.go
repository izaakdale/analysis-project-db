package main

import (
	db "github.com/izaakdale/analysis-project-db/db/access"
)

func main() {
	db.NewPostgres("", "", "")
}
