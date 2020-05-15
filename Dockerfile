# Build
FROM golang:1.14-alpine3.11 AS build

WORKDIR /quoteservice

COPY . .

RUN go build -o ./app .

# Deployment
FROM alpine:3.11
EXPOSE 8080

WORKDIR /app
RUN mkdir /app/data

COPY --from=build /quoteservice/app ./
COPY ./data/quotes.json ./data/quotes.json

ENTRYPOINT [ "./app" ]

