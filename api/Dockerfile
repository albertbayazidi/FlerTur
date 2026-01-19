FROM oven/bun:latest

WORKDIR /app

COPY package.json bun.lock ./
RUN bun install --production

COPY . .

EXPOSE 3001

CMD ["sh", "-c", "bun run scripts/migrate.ts && bun run server.ts"]
