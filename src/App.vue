<template id="app-template">
  <div id="app">

    <h2>Rock, Paper, Scissors, Lizard, Spock</h2>
    
    <h1>

      <i class="fa fa-2x fa-fw" v-bind:class="computedUser"></i>
      <i class="fa fa-2x fa-fw" v-bind:class="computedRand"></i>

    </h1>

    <h3>{{ result }}</h3>
  
    <div>

      <button class="button" @click="choose(1)">
        <i class="fa fa-hand-rock-o"></i>
      </button>
  
      <button class="button" @click="choose(2)">
        <i class="fa fa-hand-paper-o"></i>
      </button>

      <button class="button" @click="choose(3)">
          <i class="fa fa-hand-scissors-o"></i>
      </button>

      <button class="button" @click="choose(4)">
          <i class="fa fa-hand-lizard-o"></i>
      </button>

      <button class="button" @click="choose(5)">
          <i class="fa fa-hand-spock-o"></i>
      </button>

    </div>

    <h1>{{ userScore }} - {{ computerScore }}</h1>

    <button class="button" @click="reset()">
          <i class="fa fa-undo"></i>
    </button>

  </div>
</template>

<script>
import axios from "axios";

export default { 
  name: 'App',
  data: function () {
    return {
      userPick: null,
      randPick: null,
      userScore: 0,
      computerScore: 0,
      result: "",
    };
  },
  methods: {
    choose: function (pick) {

      const picks = ['rock', 'paper', 'scissors', 'lizard', 'spock'];
      
      axios.post('http://localhost:80/play', {

          player: pick,

      }).then(response => {
        
        console.log(response);

        this.userPick = picks[pick-1];
        this.randPick = picks[response.data.computer-1];

        if (response.data.results == 'won') {
          this.userScore++;
          this.result = "You won!";
        } else if (response.data.results == 'lose') {
          this.computerScore++;
          this.result = "You lost. :(";
        } else {
          this.result = "It's a tie!";
        }
        
      }).catch(error => {

        console.log(error);

      });

    },
    reset: function () {

      this.userScore = 0;
      this.computerScore = 0;
      this.userPick = null;
      this.randPick = null;
      this.result = "";

    }
  },
  computed: {
    computedUser: function () {
      return {
        "fa-circle-o-notch fa-spin": this.userPick === null,
        "fa-hand-rock-o": this.userPick === "rock",
        "fa-hand-paper-o": this.userPick === "paper",
        "fa-hand-scissors-o": this.userPick === "scissors",
        "fa-hand-lizard-o": this.userPick === "lizard",
        "fa-hand-spock-o": this.userPick === "spock"
      }
    },
    computedRand: function () {
      return {
        "fa-circle-o-notch fa-spin": this.randPick === null,
        "fa-hand-rock-o": this.randPick === "rock",
        "fa-hand-paper-o": this.randPick === "paper",
        "fa-hand-scissors-o": this.randPick === "scissors",
        "fa-hand-lizard-o": this.randPick === "lizard",
        "fa-hand-spock-o": this.randPick === "spock"
      }
    }
  } 
}; 
</script>

<style>

#app {
    height: 75vh;
    display: flex;
    flex-direction: column;
    max-width: 500px;
    justify-content: center;
    align-items: center;
    margin: 0 auto;
    font-family: "Open Sans", sans-serif;
    background-color: lightgrey;
}

.button {
    background-color: #000000;
    color: #FFFFFF;
    padding: 10px;
    border-radius: 10px;
    -moz-border-radius: 10px;
    -webkit-border-radius: 10px;
    margin:10px;
    width: 50px;
    height: 50px;
    font-size : 20px;
}

</style>
