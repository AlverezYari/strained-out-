# Build frontend
FROM node:20-alpine AS frontend-build
WORKDIR /frontend
COPY frontend/package.json ./
RUN npm install
COPY frontend ./
RUN npm run build

# Build backend
FROM golang:1.22-alpine AS backend-build
WORKDIR /backend
COPY backend/go.mod ./
RUN go mod download
COPY backend ./
RUN go build -o server ./cmd/server

# Final image
FROM alpine:latest
RUN apk add --no-cache nginx
WORKDIR /app
COPY --from=backend-build /backend/server ./
COPY --from=frontend-build /frontend/dist /usr/share/nginx/html
COPY docker/start.sh /start.sh
EXPOSE 8080 80
CMD ["/start.sh"]
