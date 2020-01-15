FROM golang:1.13 AS build
ADD src /go/src/StorageOS-Tools/src
WORKDIR /go/src/StorageOS-Tools/src
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o storageos-tools

FROM alpine
WORKDIR /app
COPY --from=build /go/src/StorageOS-Tools/src/storageos-tools /app/
ENTRYPOINT [ "/app/storageos-tools" ]