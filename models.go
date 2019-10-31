package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

//Choice : structure used to represent different options in rpsls game
type Choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//PlayerChoice : structure used to unmarshal the players choice from api call
type PlayerChoice struct {
	Player int `json:"player"`
}

//RandomNumber : structure used to unmarshal random int from api call
type RandomNumber struct {
	RandomInt int `json:"random_number"`
}

//Result : structure used to represent the result of playing one round of rpsls
type Result struct {
	Result   string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}

//Choices : returns a slice of all possible choices in the rpsls game
func Choices() []Choice {

	// create a slice that holds all possible choices in the rpsls game
	choices := []Choice{
		{1, "rock"},
		{2, "paper"},
		{3, "scissors"},
		{4, "lizard"},
		{5, "spock"},
	}

	return choices

}

//RandomChoice : returns a random choice out of the available options
func RandomChoice(client *http.Client) (*Choice, error) {

	// make a call to the provided random number generator
	resp, err := client.Get("https://codechallenge.boohma.com/random")
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	// read the response body into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	// close the response body, helps reuse connections(netstat shows no reuse :/ )
	resp.Body.Close()

	// create a RandomNumber variable and unmarshal the response body into it
	rand := RandomNumber{}
	err = json.Unmarshal(body, &rand)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	// grab a slice that holds all possible choices in the rpsls game
	choices := Choices()

	// using mod to get an even distribution, return a random choice
	return &choices[rand.RandomInt%5], nil

}

//PlayGame : takes player choice, determines computer choice, and returns result
func PlayGame(client *http.Client, pChoice Choice) (*Result, error) {

	// grab a random choice for the computer and isolate choice names
	cChoice, err := RandomChoice(client)
	if cChoice == nil {
		return nil, err
	}
	computer := cChoice.Name
	player := pChoice.Name

	// map that holds the rules to the game, maps a choice to what it beats
	rules := map[string]string{
		"rock":     "scissors,lizard",
		"paper":    "rock,spock",
		"scissors": "paper,lizard",
		"lizard":   "paper,spock",
		"spock":    "rock,scissors",
	}

	// instantie a string to hold result and figure out what happens this round
	var result string
	if player == computer {
		result = "tie"
	} else if strings.Contains(rules[player], computer) {
		result = "won"
	} else {
		result = "lose"
	}

	// return a Result variable with the determined outcome of this round
	return &Result{Result: result, Player: pChoice.ID, Computer: cChoice.ID}, nil

}
