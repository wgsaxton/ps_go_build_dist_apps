FROM --platform=linux/$TARGETARCH golang:1.20

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o /bin/teacherportalservice app/cmd/teacherportal/*.go

# The executable will run from the WORKDIR and 
# look for files starting from the WORKDIR location
# This is so the executable can find the templates
COPY app/teacherportal/*.gohtml ./teacherportal/

EXPOSE 5001

CMD ["/bin/teacherportalservice"]