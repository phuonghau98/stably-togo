FROM node:14.15 as react-app-builder
WORKDIR /app/ui
COPY ui .
RUN npm install
RUN CI="" GENERATE_SOURCEMAP=false NODE_ENV=production npm run build

# syntax=docker/dockerfile:1
FROM golang:1.16 as server-builder
WORKDIR /go/src/github.com/phuonghau98/stably-togo/
COPY . .
RUN go get \
  && CGO_ENABLED=0 GOOS=linux go build -o server .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/stably-togo
COPY --from=server-builder /go/src/github.com/phuonghau98/stably-togo/ .
COPY --from=react-app-builder /app/ui/build /root/stably-togo/ui/build
ENV ENV=production
CMD ["./server"]