FROM golang:1.18 as egts-builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make

FROM busybox

COPY --from=egts-builder /app/bin /app/
COPY --from=egts-builder /app/configs/receiver.yaml /etc/egts-receviver/config.yaml

ENTRYPOINT ["/app/receiver", "-c", "/etc/egts-receviver/config.yaml"]
