FROM golang:1.22 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -o /apiserver cmd/server/apiserver/main.go

FROM golang:1.21

RUN curl -sSf https://atlasgo.sh | sh
RUN apt update && apt install -y jq

COPY --from=build /apiserver /apiserver

WORKDIR /

CMD [ "/apiserver" ]