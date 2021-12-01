FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download curl for hot reload module
RUN apk update && \
    apk add --no-cache git ca-certificates tzdata curl && \
    update-ca-certificates

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code. The wildcard is used to copy all files with a .go
# extension located in the current directory to the current directory inside
# the image.
COPY . ./

# Build a static application binary named twitter-clone located in root of
# our image. Could put anywhere in image, this is easiest.
# RUN go build -o /twitter-clone

# Open the port
EXPOSE 8080

# Run
# CMD [ "/twitter-clone" ]
# CMD ["air", "-c", ".air.toml"]
CMD air
