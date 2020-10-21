FROM golang:latest AS buildContainer
WORKDIR /go/src/app
COPY . .
#flags: -s -w to remove symbol table and debug info
#CGO_ENALBED=0 is required for the code to run properly when copied alpine
RUN CGO_ENABLED=0 GOOS=linux go build -v -mod mod -ldflags "-s -w" -o restapi ./app

FROM alpine:latest
WORKDIR /app
COPY --from=buildContainer /go/src/app/restapi .

ENV DBUSER root
ENV DBPASSWORD jobs@123
ENV DBNAME api
ENV DBHOST localhost
ENV DBPORT 3306

EXPOSE 8080

CMD ["./restapi"]