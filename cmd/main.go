package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gihyodocker/todoapi"
)

func main() {

	env, err := todoapi.CreateEnv()
	if err != nil {
		fmt.Println(env.Bind)
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	masterDB, err := todoapi.CreateDbMap(env.MasterURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.MasterURL)
		return
	}

	slaveDB, err := todoapi.CreateDbMap(env.SlaveURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.SlaveURL)
		return
	}

	mux := http.NewServeMux()

	hc := func(w http.ResponseWriter, r *http.Request) {
		log.Println("[GET] /hc")
		w.Write([]byte("OK"))
	}

	todoHandler := todoapi.NewTodoHandler(masterDB, slaveDB)

	mux.Handle("/todo", todoHandler)
	mux.HandleFunc("/hc", hc)

	s := http.Server{
		Addr:    env.Bind,
		Handler: mux,
	}

	log.Printf("Listen HTTP Server")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
