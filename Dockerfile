FROM golang:1.22-alpine3.19 AS buildStage
WORKDIR /order-svc
COPY . ./
RUN go mod download
RUN go build -o ./order-svc ./cmd

FROM scratch AS release-stage 
WORKDIR /
COPY --from=buildStage /order-svc/order-svc /order-svc
COPY --from=buildStage /order-svc/dev.env /

EXPOSE 50053

ENTRYPOINT [ "/order-svc" ]