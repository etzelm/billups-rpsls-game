# Stage 1 : Build Golang Backend API
FROM golang:1.9.2-alpine3.7 as go-build
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep
COPY . src/github.com/etzelm/billups-rpsls-game/
WORKDIR src/github.com/etzelm/billups-rpsls-game/
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o billups-rpsls-game .

# Stage 2 : Build Vue.js Frontend Service
FROM node:latest as vue-build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY ./ .
RUN npm run build

# Stage 3 : Combine them together into a single image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=vue-build /app/dist /dist 
COPY --from=go-build /go/src/github.com/etzelm/billups-rpsls-game/billups-rpsls-game .
CMD ["./billups-rpsls-game"]

EXPOSE 80