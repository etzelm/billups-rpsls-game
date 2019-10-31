package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

//ChoicesResp : Returns all of the different choices players have in game
func ChoicesResp(c *gin.Context) {

	// Variable that holds all possible choices in the rpsls game
	choices := Choices()

	// Call the JSON method of the Context to render the given interface into JSON
	c.JSON(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Pass the data that the response uses
		choices,
	)

}

//ChoiceResp : Returns a random choice out of the available options
func ChoiceResp(client *http.Client) gin.HandlerFunc {

	//Define a handler function that uses the provided http client from main
	fn := func(c *gin.Context) {

		// Variable that holds a randomly generated choice selected from Choices()
		choice, err := RandomChoice(client)
		if choice == nil {
			c.AbortWithStatusJSON(500, map[string]string{
				"error": err.Error(),
			})
			return
		}

		// Call the JSON method of the Context to render the given interface into JSON
		c.JSON(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Pass the data that the response uses
			choice,
		)
	}

	return gin.HandlerFunc(fn)

}

//PlayResp : Takes in player provided choice, plays the rpsls game, and returns the result
func PlayResp(client *http.Client) gin.HandlerFunc {

	//Define a handler function that uses the provided http client from main
	fn := func(c *gin.Context) {

		// Create a PlayerChoice variable to bind POST body JSON to, then do just that
		pc := PlayerChoice{}
		if err := c.ShouldBindWith(&pc, binding.JSON); err != nil {
			log.Errorln(err)
			c.AbortWithStatusJSON(422, map[string]string{
				"error": err.Error(),
			})
			return
		}

		// Check for an empty or invalid player choice
		if pc.Player == 0 || pc.Player > 5 {
			log.Errorln("Empty or invalid player choice.")
			c.AbortWithStatusJSON(422, map[string]string{
				"error": "Empty or invalid player choice",
			})
			return
		}

		// Grab all choices and use them to get full context for players choice
		choices := Choices()
		choice := choices[pc.Player-1]

		// Play rpsls and return the result
		result, err := PlayGame(client, choice)
		if result == nil {
			c.AbortWithStatusJSON(500, map[string]string{
				"error": err.Error(),
			})
			return
		}

		// Call the JSON method of the Context to render the given interface into JSON
		c.JSON(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Pass the data that the response uses
			result,
		)
	}

	return gin.HandlerFunc(fn)

}
