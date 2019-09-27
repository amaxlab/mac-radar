FROM golang
WORKDIR /go/src/github.com/amaxlab/mac-radar
ADD . /go/src/github.com/amaxlab/mac-radar

RUN go get -u github.com/amaxlab/go-lib/log github.com/amaxlab/go-lib/config github.com/go-chi/chi github.com/go-routeros/routeros
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o /go/bin/mac-radar

FROM scratch
COPY --from=0 /go/bin/mac-radar /go/bin/mac-radar

ENV APP_DEBUG="false"
ENV APP_PORT="8080"
ENV APP_MIKROTIK_ADDRESS="192.168.88.1:8728"
ENV APP_MIKROTIK_USERNAME="admin"
ENV APP_MIKROTIK_PASSWORD="admin"

EXPOSE 8080

ENTRYPOINT ["/go/bin/mac-radar"]
