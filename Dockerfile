FROM public.ecr.aws/docker/library/golang:1.21.0 as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  llama-user

WORKDIR /app
ARG GOPROXY=direct
COPY go.* ./
RUN go mod download
RUN go mod verify
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /server .

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /server .

USER llama-user:llama-user

CMD ["./server"]

