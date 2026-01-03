FROM golang:1.24.1-alpine AS application_builder
WORKDIR /app
COPY go.mod /app/
RUN go mod download
COPY . .
RUN go build -o bank .

FROM alpine:latest AS final_image
WORKDIR /app
COPY --from=application_builder /app/bank /app/
COPY --from=application_builder /app/HTML /app/HTML
COPY --from=application_builder /app/asserts /app/asserts
COPY --from=application_builder /app/migrations /app/migrations
ENTRYPOINT [ "/app/bank" ]

