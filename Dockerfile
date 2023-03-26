FROM golang:1.20

WORKDIR /usr/src/server

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build 

EXPOSE 3000
CMD ["./api"]
#CMD ["bash"]


# docker build -t api .