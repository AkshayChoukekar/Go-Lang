package main

import (
	"log"
	"strconv"
)

func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("InitPool Error: %s", err)
	}
	conn1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool loan error: %s", err)
	}
	conn2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan error: %s", err)
	}
	pool.receive(conn1)
	pool.receive(conn2)
}
