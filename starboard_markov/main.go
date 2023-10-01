package main

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mb-14/gomarkov"
)

const order = 10

var separator = ""

func main() {
	db, err := sql.Open("sqlite3", "bot.db")
	if err != nil {
		panic(err)
	}

	var entries []string

	rows, err := db.Query("SELECT content FROM posts WHERE content != ''")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var entry string
		if err := rows.Scan(&entry); err != nil {
			panic(err)
		}
		entries = append(entries, entry)
	}

	chain := gomarkov.NewChain(order)

	for _, entry := range entries {
		chain.Add(strings.Split(entry, separator))
	}

	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - order):])
		tokens = append(tokens, next)
	}
	println(strings.Join(tokens[order:len(tokens)-1], separator))
}
