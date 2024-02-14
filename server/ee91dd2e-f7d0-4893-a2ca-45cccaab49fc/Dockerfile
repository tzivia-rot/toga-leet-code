FROM golang:latest
	WORKDIR /app
	COPY . .
	RUN go build -o function
	CMD ["./function"]