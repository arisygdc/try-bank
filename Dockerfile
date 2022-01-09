FROM golang:1.17.6-alpine3.14 AS stage
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.14
WORKDIR /app
COPY --from=stage /app/main .
COPY --from=stage /app/config.env .
EXPOSE 8080
CMD [ "./main" ]