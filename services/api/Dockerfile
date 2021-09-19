FROM node:12-alpine3.12

WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY ./src ./src
COPY ./server.js ./server.js

EXPOSE 8080
CMD [ "npm", "run", "start" ]
