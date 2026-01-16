import { sql } from "bun";
import { seedData } from "./data.js";

async function seed() {
  console.log("Starting database seeding...");

  try {
    await sql.begin(async (tx) => {
      await tx`DELETE FROM page_data_results`;
      await tx`DELETE FROM page_data_wrappers`;

      for (const entry of seedData) {
        const [wrapper] = await tx`
          INSERT INTO page_data_wrappers (start_station, end_station, retrieval_time)
          VALUES (${entry.startStation}, ${entry.endStation}, ${entry.retrievalTime})
          RETURNING id
        `;

        const wrapperId = wrapper.id;

        for (const res of entry.results) {
          await tx`
            INSERT INTO page_data_results (
              wrapper_id, 
              duration, 
              start_time, 
              price, 
              number_of_trains, 
              train_ids, 
              url
            )
            VALUES (
              ${wrapperId}, 
              ${res.duration}, 
              ${res.startTime}, 
              ${res.price}, 
              ${res.numberOfTrains}, 
              ${sql.array(res.trainIds)}, 
              ${res.url}
            )
          `;
        }
      }
    });

    console.log("Database seeded successfully.");
  } catch (error) {
    console.error("Error during seeding:", error);
    process.exit(1);
  }
}

seed();
