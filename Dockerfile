FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code. The wildcard is used to copy all files with a .go
# extension located in the current directory to the current directory inside
# the image.
COPY *.go ./

# Build a static application binary named twitter-clone located in root of
# our image. Could put anywhere in image, this is easiest.
RUN go build -o /twitter-clone

# Open the port
EXPOSE 8080

# Run
CMD [ "/twitter-clone" ]
