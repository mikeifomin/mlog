FROM node:8 as front

COPY ./front/package.json /app/front/package.json
COPY ./front/package-lock.json /app/front/package-lock.json

RUN cd /app/front/ && npm i --loglevel="warn" 

COPY front /app/front

WORKDIR /app/front
RUN npm run build

FROM golang:1.10 as compiler
COPY ./ /go/src/github.com/mikeifomin/mlog
WORKDIR /go/src/github.com/mikeifomin/mlog


COPY --from=front /app/front/dist /var/www/html

RUN go build -o /bin/mlog cmd/server/main.go

FROM debian:jessie
ENV ADMIN_DIR=/var/www/html
COPY --from=compiler /bin/mlog /bin/mlog
EXPOSE 80 
ENTRYPOINT ["bin/mlog"]
