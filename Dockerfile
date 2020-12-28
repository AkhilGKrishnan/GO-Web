FROM golang:latest
WORKDIR /app
ADD . .
RUN go get github.com/pilu/fresh
CMD fresh

