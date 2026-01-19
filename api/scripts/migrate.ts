import { sql } from "bun";

async function migrate() {
  console.log("Starting database migration...");

  try {
    await sql.begin(async (tx) => {
      await tx`
        CREATE TABLE IF NOT EXISTS page_data_wrappers (
          id SERIAL PRIMARY KEY,
          start_station TEXT NOT NULL,
          end_station TEXT NOT NULL,
          retrieval_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );
      `;

      await tx`
        CREATE TABLE IF NOT EXISTS page_data_results (
          id SERIAL PRIMARY KEY,
          wrapper_id INTEGER REFERENCES page_data_wrappers(id) ON DELETE CASCADE,
          duration TEXT,
          start_time TEXT,
          price INTEGER,
          number_of_trains INTEGER,
          train_ids TEXT[], 
          url TEXT,
          created_at TIMESTAMPTZ DEFAULT NOW()
        );
      `;

      await tx`
        CREATE INDEX IF NOT EXISTS idx_route_search 
        ON page_data_wrappers (start_station, end_station);
      `;
    });

    console.log("Migration completed successfully.");
  } catch (error) {
    console.error("Migration failed:", error);
    process.exit(1);
  }
}

migrate();
