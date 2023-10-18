FROM golang:1.20

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/logservice app/cmd/logservice/*.go

EXPOSE 4000

CMD ["/bin/logservice"]