# Multi-stage build for React application
FROM node:18-alpine AS base

# Install dependencies only when needed
FROM base AS deps
WORKDIR /app

# Copy package files
COPY package.json package-lock.json* ./

# Install dependencies
RUN npm ci

# Rebuild the source code only when needed
FROM base AS builder
WORKDIR /app

# Accept base path as build argument
ARG BASE_PATH=/
ENV BASE_PATH=${BASE_PATH}
# Set API base URL to use the same base path
ENV VITE_API_BASE_URL=${BASE_PATH}/api

COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Build application with base path and API URL
RUN npm run build

# Production image, copy all the files and run next
FROM nginx:alpine AS runner
WORKDIR /app

# Copy built assets from builder
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port
EXPOSE 3000

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
