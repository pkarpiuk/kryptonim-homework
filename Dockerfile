FROM golang:1.21.3
WORKDIR /kryptonim-homework
COPY . ./
RUN go mod tidy
RUN go build -o /homework
EXPOSE 8080
CMD ["/homework"]

