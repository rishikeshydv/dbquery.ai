FROM --platform=linux/amd64 golang:1.18-bullseye
WORKDIR /
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /go-backend
EXPOSE 9002
CMD [ "/go-backend" ]