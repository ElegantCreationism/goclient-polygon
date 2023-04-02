#################################
# STEP 1 Build Executable Binary
##################################
FROM golang:alpine as builder

WORKDIR /app
COPY . ./

# Build the binary
RUN go mod download; \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -o ./bin/svc -ldflags="-s -w"

################################
# STEP 2 Build a small image
################################
FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /app/bin/svc /svc

# Port on which the service will be exposed.
EXPOSE 8080

# Run the svc binary.
CMD ["./svc"]
