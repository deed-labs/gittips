FROM node:19-alpine as build

RUN npm install -g serve

WORKDIR /app
RUN apk add --no-cache bash
COPY package.json /app/package.json
RUN npm install
COPY . /app
RUN npm run build

CMD serve -p $PORT -s build
