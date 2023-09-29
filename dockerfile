FROM golang:latest as builder

WORKDIR /app

COPY .  .

RUN GOOS="linux" CGO_ENABLED=0  go build -ldflags="-w -s" -o server ./src/
 

FROM scratch


COPY --from=builder /app .

# CMD ["sleep infinity"]

CMD [ "./server" ]