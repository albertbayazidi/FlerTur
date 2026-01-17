import { sql } from "bun";

const PORT = 3001;

console.log(`API Server running at http://localhost:${PORT}`);

Bun.serve({
  port: PORT,
  async fetch(req) {
    const url = new URL(req.url);

    // 1. Handle CORS (Allow React to talk to Bun)
    if (req.method === "OPTIONS") {
      return new Response(null, {
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods": "GET, OPTIONS",
          "Access-Control-Allow-Headers": "Content-Type",
        },
      });
    }

    // 2. The Search Route
    if (url.pathname === "/api/search") {
      const from = url.searchParams.get("from");
      const to = url.searchParams.get("to");

      if (!from || !to) {
        return new Response("Missing 'from' or 'to' parameters", { status: 400 });
      }

      try {
        const results = await sql`
          SELECT 
            w.start_station as "startStation",
            w.end_station as "endStation",
            w.retrieval_time as "retrievalTime",
            COALESCE(
              json_agg(
                json_build_object(
                  'duration', r.duration,
                  'startTime', r.start_time,
                  'price', r.price,
                  'numberOfTrains', r.number_of_trains,
                  'trainIds', r.train_ids,
                  'url', r.url
                ) ORDER BY r. start_time
              ) FILTER (WHERE r.id IS NOT NULL), 
              '[]'
            ) AS "pageDataResults"
          FROM page_data_wrappers w
          LEFT JOIN page_data_results r ON w.id = r.wrapper_id
          WHERE w.start_station ILIKE ${from} 
            AND w.end_station ILIKE ${to}
          GROUP BY w.id
          ORDER BY w.retrieval_time DESC
          LIMIT 1 
        `;

        return Response.json(results, {
          headers: { "Access-Control-Allow-Origin": "*" },
        });

      } catch (error) {
        console.error(error);
        return new Response("Database Error", { status: 500 });
      }
    }

    return new Response("Not Found", { status: 404 });
  },
});
