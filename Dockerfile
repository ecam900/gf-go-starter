# Just for running Air Hotreload
FROM golang:latest
RUN go get -u github.com/cosmtrek/air
WORKDIR /app
ENTRYPOINT ["air"]