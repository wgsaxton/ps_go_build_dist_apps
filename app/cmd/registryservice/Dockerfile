FROM golang:1.20

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/registryservice app/cmd/registryservice/*.go

EXPOSE 3000

CMD ["/bin/registryservice"]