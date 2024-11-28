FROM golang:1.23
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main .
EXPOSE 8080
# Command to run the executable
CMD ["./main"]