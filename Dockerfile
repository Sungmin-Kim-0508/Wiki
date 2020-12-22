FROM golang:1.15 as goBuilder
WORKDIR /go/src/goApp
COPY ./backend ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main.exe
EXPOSE 9090

FROM node:12 as nodeBuilder
WORKDIR /webApp
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend ./
RUN npm run build

# FROM nginx
# COPY ./frontend/nginx/default.conf /etc/nginx/conf.d/default.conf
# COPY --from=nodeBuilder /webApp/build /usr/share/nginx/html

FROM nginx
COPY --from=nodeBuilder /webApp/build/ /usr/share/nginx/html
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=goBuilder /go/src/goApp .
COPY --from=nodeBuilder /webApp/build /build
ENTRYPOINT ./main.exe
# CMD ["./main.exe"]
# CMD ["./main.exe"]  

# CMD ["./main.exe"]
