FROM node:16-alpine

WORKDIR /usr/src/app

COPY rollup.config.js ./
COPY package*.json ./
COPY postcss.config.js ./
COPY tailwind.config.js ./
COPY tsconfig.json ./

RUN npm install

COPY ./src ./src
COPY ./public ./public

RUN npm run build

EXPOSE 5000

ENV HOST=0.0.0.0

CMD [ "npm", "start" ]