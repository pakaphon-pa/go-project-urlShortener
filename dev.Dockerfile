FROM golang:1.16-alpine
RUN mkdir /app
ADD .. /app/
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT ["air"]