import { sql } from "bun";

const PORT = Number(process.env.PORT) || 3005;
const PUBLIC_DIR = "./public";

console.log(`API Server running at http://localhost:${PORT}`);

Bun.serve({
	port: PORT,
	async fetch(req) {
		const url = new URL(req.url);

		if (req.method === "OPTIONS") {
			return new Response(null, {
				headers: {
					"Access-Control-Allow-Origin": "*",
					"Access-Control-Allow-Methods": "GET, OPTIONS",
					"Access-Control-Allow-Headers": "Content-Type",
				},
			});
		}

		if (url.pathname === "/api/search") {
			const from = url.searchParams.get("from");
			const to = url.searchParams.get("to");

			if (!from || !to) {
				return new Response("Missing 'from' or 'to' parameters", {
					status: 400,
				});
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
                ) ORDER BY r.start_time
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

		if (url.pathname.startsWith("/api/")) {
			return new Response("API Endpoint Not Found", { status: 404 });
		}

		const filePath = url.pathname === "/" ? "/index.html" : url.pathname;
		const file = Bun.file(`${PUBLIC_DIR}${filePath}`);

		if (await file.exists()) {
			return new Response(file);
		}

		const indexFile = Bun.file(`${PUBLIC_DIR}/index.html`);
		if (await indexFile.exists()) {
			return new Response(indexFile);
		}

		return new Response("Not Found", { status: 404 });
	},
});
