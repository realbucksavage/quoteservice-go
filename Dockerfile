# Build
FROM golang:1.13-alpine AS build

WORKDIR /quoteservice

COPY . .

RUN go mod vendor && go build -o ./app .

# Deployment
FROM alpine:3.7
EXPOSE 8080

WORKDIR /app
RUN mkdir /app/data

COPY --from=build /quoteservice/app ./
COPY ./data/quotes.json ./data/quotes.json

ENTRYPOINT [ "./app" ]

