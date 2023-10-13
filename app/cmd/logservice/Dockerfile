FROM golang:1.20

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o logservice /app/cmd/logservice/*

EXPOSE 4000

CMD ["/logservice"]