FROM golang:1.24 AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -o /out/task-api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/task-api /task-api
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/task-api"]
