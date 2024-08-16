FROM golang:alpine

WORKDIR /app

COPY . .

# Download and install the dependencies:
RUN go get -d -v ./...

# Install Air:
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air"]
