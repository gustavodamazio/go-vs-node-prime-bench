FROM node:20 AS builder
RUN mkdir -p /home/node/app/node_modules && chown -R node:node /home/node/app
WORKDIR /home/node/app
COPY package*.json ./
RUN npm i
COPY --chown=node:node ./ ./
RUN npm run build

FROM gcr.io/distroless/nodejs20-debian12
COPY --from=builder /home/node/app /app
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app/lib/index.js"]
