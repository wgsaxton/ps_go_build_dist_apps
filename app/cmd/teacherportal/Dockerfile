FROM --platform=linux/$TARGETARCH golang:1.20

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o teacherportalservice app/cmd/teacherportal/*.go

LABEL org.opencontainers.image.source=https://github.com/wgsaxton/ps_go_build_dist_apps

# The executable will run from the WORKDIR and 
# look for files starting from the WORKDIR location
# This is so the executable can find the templates
COPY app/teacherportal/*.gohtml ./teacherportal/

# Trying this to see if the binary can find the files on startup
COPY app/teacherportal/*.gohtml /bin/teacherportal/

EXPOSE 5001

CMD ["/build/teacherportalservice"]