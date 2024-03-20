package main

import (
	"dimasfadilah/go-restful-api/helper"
	"fmt"
	"testing"
)

func TestInitializedServer(t *testing.T) {
	server := InitializedServer()

	err := server.ListenAndServe()
	fmt.Println("Server is running on localhost:3000")
	helper.PanicIfError(err)
}
