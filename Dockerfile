# pull official base image
FROM node:13.12.0-alpine as react-build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
ADD ./frontend ./
RUN npm install --silent
RUN npm run test -- --watchAll=false
RUN npm run build

FROM golang:alpine as go-build
WORKDIR /app
ADD ./api ./
#RUN go test
RUN go build -o main .

FROM golang:alpine
WORKDIR /app/code
COPY --from=react-build /app/build ./frontend/build
COPY --from=go-build /app/main ./api/
COPY ./api/properties/*.properties ./properties/

RUN adduser -S -D -H -h /app appuser
USER appuser
EXPOSE 8000
CMD ["./api/main"]