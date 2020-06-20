package v1

import (
	"log"
	"net/http"

	"github.com/amagimedia/seshat/repository"
	"github.com/labstack/echo"
)

// ListAgent list all the agents registered
func ListAgent(c echo.Context) error {
	agents, err := repository.Agent{}.List()
	if err != nil {
		log.Println("Failed to get details", err)
		return err
	} else {
		return c.JSONPretty(http.StatusOK, agents, "  ")
	}
}

// CreateAgent create agent
func CreateAgent(c echo.Context) (err error) {
	agent := new(repository.Agent)
	if err = c.Bind(agent); err != nil {
		return
	}

	if err = c.Validate(agent); err != nil {
		log.Println(err)
		return
	}
	err = agent.Create()

	if err = c.Validate(agent); err != nil {
		log.Println(err)
		return
	}
	return c.JSONPretty(http.StatusOK, agent, "  ")
}
