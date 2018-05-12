package todoapi

import (
	"errors"
	"os"
)

type Env struct {
	Bind      string
	MasterURL string
	SlaveURL  string
}

func CreateEnv() (*Env, error) {

	env := Env{}

	bind := os.Getenv("TODO_BIND")
	if bind == "" {
		env.Bind = ":8080"
	}
	env.Bind = bind

	masterURL := os.Getenv("TODO_MASTER_URL")
	if masterURL == "" {
		return nil, errors.New("TODO_MASTER_URL is not specified")
	}
	env.MasterURL = masterURL

	slaveURL := os.Getenv("TODO_SLAVE_URL")
	if slaveURL == "" {
		return nil, errors.New("TODO_SLAVE_URL is not specified")
	}
	env.SlaveURL = slaveURL

	return &env, nil
}
