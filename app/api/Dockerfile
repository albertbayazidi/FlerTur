FROM oven/bun:latest AS build-frontend
WORKDIR /build/frontend
COPY frontend/package.json frontend/bun.lock ./
RUN bun install
COPY frontend/ ./
RUN bun run build

FROM oven/bun:latest AS runtime
WORKDIR /app
COPY api/package.json api/bun.lock ./
RUN bun install --production
COPY api/ ./

COPY --from=build-frontend /build/frontend/dist ./public

EXPOSE 3005
CMD ["sh", "-c", "bun run scripts/migrate.ts && bun run server.ts"]
