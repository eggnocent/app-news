FROM golang:1.23.2-alphine As builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o main .

FROM gcr.io/disctroless/base-debian10

COPY --from=builder /app/main /app/main

COPY ./docs /app/docs

COPY .env /app/.env

WORKDIR /app/

EXPOSE 8000

CMD [ "/app/main" ]