# STEP 1 build executable binary
FROM golang:alpine as builder
# Install git + SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/Journerist/drone-backend/
WORKDIR $GOPATH/src/Journerist/drone-backend/
#get dependancies
RUN go get ./...
#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -installsuffix cgo -o /go/bin/drone-backend ./main/main.go

# STEP 2 build a small image
# start from scratch
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/drone-backend /go/bin/drone-backend
USER appuser
ENTRYPOINT ["/go/bin/drone-backend"]