FROM golang:1.20 as builder

WORKDIR /go/src/github.com/rajatjindal/goodfirstissue
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go test ./... -cover
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w" -o bin/goodfirstissue main.go

FROM alpine:3.17.2

RUN mkdir -p /home/app

# Add non root user
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

WORKDIR /home/app

USER app

COPY --from=builder /go/src/github.com/rajatjindal/goodfirstissue/bin/goodfirstissue /usr/local/bin/

ENTRYPOINT "/usr/local/bin/goodfirstissue"