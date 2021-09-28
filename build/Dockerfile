# pull official base image
FROM node:13.12.0-alpine as react-build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
ADD ./web ./
RUN npm install --silent
RUN npm run test -- --watchAll=false
RUN npm run build

FROM golang:alpine as go-build
WORKDIR /app
ADD ./cmd/cryptchat ./cmd/cryptchat
ADD ./configs ./configs
ADD ./internal ./internal
ADD ./go.mod .
ADD ./go.sum .
ADD ./vendor ./vendor
RUN #go test -v ./...
RUN go build -o cryptchat /app/cmd/cryptchat

FROM golang:alpine
WORKDIR /app/code
COPY --from=react-build /app/build ./frontend/build/
COPY --from=go-build /app/cryptchat ./api/
ADD ./configs/*.properties ./configs/

RUN adduser -S -D -H -h /app appuser
USER appuser
EXPOSE 8000
CMD ["./api/cryptchat"]