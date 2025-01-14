package main

import (
	"api/functions"

	"github.com/google/uuid"
)

func main() {
	id :=uuid.New().String()
	functions.DisplayId(id)
	
}