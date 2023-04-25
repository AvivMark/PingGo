FROM golang:1.19 AS gobuilder

WORKDIR /go/src/
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY ../ .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app ./cmd/pinggo/. 

FROM node:19-alpine3.16 as nodebuilder
WORKDIR /app
COPY ../web/ui/package.json .
COPY ../web/ui/package-lock.json .
RUN npm install
COPY ../web/ui/. .
RUN npm run build

FROM nginx:latest
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=nodebuilder /app/build /usr/share/nginx/html
COPY ../web/ui/nginx/nginx.conf /etc/nginx/conf.d

WORKDIR /root
COPY --from=gobuilder /go/src/app ./
COPY --from=gobuilder /go/src/hosts.json ./
COPY ./entrypoint.sh ./

EXPOSE 8080
EXPOSE 5000

ENTRYPOINT [ "./entrypoint.sh"]
