FROM golang:alpine as builder
WORKDIR /app 
ADD . /app 
RUN go mod tidy 
RUN cd /app && go build -o main-svc


FROM alpine 
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* 
WORKDIR /app 
COPY --from=builder /app/main-svc /app

EXPOSE 3001 
ENTRYPOINT [ "./main-svc" ]