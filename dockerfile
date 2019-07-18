FROM golang:1.12-alpine as builder

RUN apk --update add git upx

WORKDIR /go/src/gituhb.com/serverlessp/bridle
ENV GO111MODULE=on
ADD go.mod .
ADD go.sum .
RUN go mod download

ADD *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o bridle .

# Make things even smaller!
RUN upx bridle

FROM scratch
COPY --from=builder /go/src/gituhb.com/serverlessp/bridle/ /
CMD ["/bridle"]