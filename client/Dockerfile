FROM node:18-alpine as client

WORKDIR /app

COPY yarn.lock package.json ./

RUN yarn install --frozen-lockfile

COPY . .

RUN yarn build

FROM nginx:1.24-alpine as server

COPY --from=client /app/dist /usr/share/nginx/html

COPY /nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80