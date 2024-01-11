FROM golang:1.22rc1-alpine3.19
WORKDIR /app

COPY . .

CMD ["go", "run", "./main.go"]