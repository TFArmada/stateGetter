FROM golang:latest as builder
WORKDIR /stateGetter
COPY . .
RUN CGO_ENABLED=0 go build
FROM gcr.io/distroless/static AS final
COPY --from=builder /stateGetter/stateGetter /bin/stateGetter
ENTRYPOINT ["/bin/stateGetter"]
