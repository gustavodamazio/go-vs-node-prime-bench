FROM golang:1.22 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go mod tidy
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app -a -installsuffix cgo /go/src/app/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /app
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app"]
