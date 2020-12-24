# FROM golang:1.15 as goBuilder
FROM golang:1.15
WORKDIR /go/src/goApp
COPY ./backend ./
ARG GROUP_ID
ARG USER_ID
RUN CGO_ENABLED=0 GOOS=linux go build -o main.exe
EXPOSE 9090

# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=goBuilder /go/src/goApp .
# # RUN chmod +x ./main.exe
# ENTRYPOINT ./main.exe
# EXPOSE 9090
# CMD ["./main.exe"]

FROM node:12 as nodeBuilder
WORKDIR /webApp
COPY ./frontend/package.json ./
RUN npm install
COPY ./frontend ./
RUN npm run build

FROM nginx:latest
EXPOSE 8080
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=nodeBuilder /webApp/build /usr/share/nginx/html
