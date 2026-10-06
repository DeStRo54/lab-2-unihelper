FROM node:24-alpine AS build
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci --ignore-scripts

COPY . .
RUN npm run build

FROM nginx:1.28.0-alpine
COPY --from=build /app/dist/ /usr/share/nginx/html/
COPY nginx.conf /etc/nginx/conf.d/default.conf

RUN addgroup -S app && adduser -S -G app app \
    && chown -R app:app /usr/share/nginx/html /var/cache/nginx /var/log/nginx /etc/nginx/conf.d \
    && touch /var/run/nginx.pid && chown app:app /var/run/nginx.pid

USER app
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]


