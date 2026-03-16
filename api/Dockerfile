FROM oven/bun:alpine

WORKDIR /app

COPY package.json bun.lock ./

RUN bun install --production

COPY . .

EXPOSE 3005

CMD ["sh", "-c", "bun run scripts/migrate.ts && bun run server.ts"]
#CMD ["sh", "-c", "bun run scripts/migrate.ts && bun run scripts/seed.ts && bun run server.ts"]
