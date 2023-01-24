############################
# STEP 1 build executable binary
############################
FROM golang:1.18-alpine AS builder

RUN apk --update --no-cache add \
    openssl \
    git \
    curl \
    tzdata \
    ca-certificates \
    && update-ca-certificates

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -v -o anymind ./cmd/app/main.go

############################
# STEP 2 build a small image
############################
FROM alpine:3.15
COPY --from=builder /app/resources /resources
COPY --from=builder /app/migration /migration
COPY --from=builder /app/anymind /app/anymind
COPY --from=builder /app/wait-for.sh /app/wait-for.sh

RUN pwd
RUN ls -al

EXPOSE 8811
CMD [ "/app/anymind", "--port", "8811" ]
