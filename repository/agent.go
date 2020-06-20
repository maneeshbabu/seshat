package repository

import (
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"
)

type (
	Agent struct {
		ID      int64     `localIndex:"ID-Seq-index,range"`
		Code    string    `dynamo:"code,hash" json:"code"`
		Secret  string    `json:"-"`
		Time    time.Time `dynamo:"time,range" json:"time"`
		Name    string    `dynamo:"name" json:"name" validate:"required"`
		Type    string    `dynamo:"type" json:"type" validate:"required"`
		Blip    string    `dynamo:"blip" json:"blip" validate:"required"`
		Account string    `dynamo:"account" json:"account" validate:"required"`
		Feed    string    `dynamo:"feed" json:"feed" validate:"required"`
	}
	CustomValidator struct {
		Validator *validator.Validate
	}
)

// Validate validates input data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// TableName returns the table name
func (a Agent) TableName() string {
	isLambda := strings.ToUpper(os.Getenv("LAMBDA"))

	if isLambda == "TRUE" {
		return os.Getenv("AGENT_TABLE")
	} else {
		return "agents"
	}
}

// List all the agents
func (a Agent) List() (agents []Agent, err error) {
	err = Table(a.TableName()).Scan().All(&agents)
	return agents, err
}

// Create agent
func (a *Agent) Create() (err error) {
	a.Code = UUID()
	a.Time = time.Now()
	a.Secret = Token()
	err = Table(a.TableName()).Put(a).Run()
	return err
}
