FROM --platform=linux/$TARGETARCH golang:1.20

ARG TARGETARCH

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o /bin/logservice app/cmd/logservice/*.go

EXPOSE 4000

CMD ["/bin/logservice"]