# use official Golang image
FROM golang:1.16.3-alpine3.13

# set working directory
WORKDIR /app

# Copia el script wait-for-it.sh a /usr/local/bin/
COPY wait-for-it.sh /usr/local/bin/

# Asegúrate de que el script es ejecutable
RUN chmod +x /usr/local/bin/wait-for-it.sh

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

#EXPOSE the port
EXPOSE 3000

CMD ["./api"]