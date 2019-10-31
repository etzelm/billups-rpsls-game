# billups-rpsls-game

### Basic Assumptions For Evaluating Project As Docker Container
```
You have docker installed and running
Build command must be run inside provided repository
```

### Commands For Building & Running Project As Docker Container
```
docker build -t game . 
docker run -d -p 80:80 game
```

### Basic Assumptions For Evaluating Project Locally
```
Both go and npm are installed locally if you dont want to use Docker
Repository is in an appropriate GOPATH (EX: ~/go/src/github.com/etzelm/billups-rpsls-game)
Commands Must Be Run Inside Provided Repository
```

### Commands For Running Project Locally
```
npm install
dep ensure
npm run build
go run main.go handlers.go models.go
```

### Accessing Game & Notes
```
localhost:80
Sometimes the initial connection to the random number api endpoint can take a noticable amount of time
Wait for it to complete or timeout and try again, have never seen it continuously fail/lag
```

### Sources
```
https://github.com/sergivillar/golang-rock-paper-scissors
https://blog.alexellis.io/golang-json-api-client/
https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
https://stackoverflow.com/questions/17948827/reusing-http-connections-in-golang/17948937#17948937
https://stackoverflow.com/questions/39813587/go-client-program-generates-a-lot-a-sockets-in-time-wait-state/39834253#39834253
https://gist.github.com/tsenart/5fc18c659814c078378d
https://codepen.io/cloeff/pen/zEqBVL?editors=1010
https://fontawesome.com/v4.7.0/icons/
https://codesandbox.io/s/31y9jo58p5
https://www.itsolutionstuff.com/post/vue-axios-post-request-exampleexample.html
https://vuejs.org/v2/cookbook/using-axios-to-consume-apis.html
https://medium.com/travis-on-docker/triple-stage-docker-builds-with-go-and-angular-1b7d2006cb88
https://github.com/benc-uk/vuego-demoapp
```
