FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./questionsandanswers"]

