FROM --platform=linux/$TARGETARCH golang:1.20

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o /bin/gradingservice app/cmd/gradingservice/*.go

EXPOSE 6000

CMD ["/bin/gradingservice"]