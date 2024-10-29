# Stage 1: Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY . .

RUN go mod download
RUN go build -o ./mentorapi

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/mentorapi ./mentorapi
CMD ["/app/mentorapi"]