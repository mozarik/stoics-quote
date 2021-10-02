FROM golang:alpine as builder
WORKDIR /app 
ADD . /app 
RUN go mod tidy 
RUN cd /app && go build -o auth-api-gateway


FROM alpine 
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* 
WORKDIR /app 
COPY --from=builder /app/auth-api-gateway /app

EXPOSE 3002
ENTRYPOINT [ "./auth-api-gateway" ]