FROM golang:alpine

WORKDIR /app

COPY . .

# Download and install the dependencies:
RUN go get -d -v ./...

# Install Air:
RUN go install github.com/air-verse/air@latest

# Install Goose:
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 8080

CMD ["air"]
