FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./auth_api

EXPOSE 8080

CMD [ "./auth_api" ]