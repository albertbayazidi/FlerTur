import Route from "./Route.jsx";

function Info({ onRouteSelect }) {
const routes1 = [
    { start: "Oslo S", end: "Bergen stasjon" },
    { start: "Oslo S", end: "Stavanger stasjon" },
    { start: "Oslo S", end: "Trondheim S" },
    { start: "Bergen stasjon", end: "Fredrikstad stasjon" },
  ];

  const routes2 = [
    { start: "Bergen stasjon", end: "Kristiansand stasjon" },
    { start: "Stavanger stasjon", end: "Bergen stasjon" },
    { start: "Stavanger stasjon", end: "Fredrikstad stasjon" },
    { start: "Stavanger stasjon", end: "Kristiansand stasjon" },
  ];

  return (
    <div className="mx-auto mt-4 w-5/6 md:mx-auto grid grid-cols-1 md:grid-cols-2 gap-8">
      <div className="w-full mx-auto text-lg px-6 py-6 rounded-xl bg-second ">
        <h1 className="mb-8 text-3xl text-center font-bold text-white">Info:</h1>

        <p className="text-center text-white ">
          <strong>NB!</strong> Denne nettsiden er under konstruksjon. Den er ment som en demo nettside.</p>

        <p className="text-center text-white">
          Denne nettsiden prøver å finne de billigste billettprisene en uke fram i tid. Jeg tjener ikke noe på salget av
          billetene. Du kan finne hele prosjektet på github og en forklaring på hvorfor den bare ser en uke fram i tid.
        </p>
      </div>

      <div className="w-full mx-auto text-lg px-6 py-6 rounded-xl bg-second">
        <h1 className="mb-8 text-3xl text-center font-bold text-white">Tilgjengelige Ruter:</h1>
        <div className="grid grid-cols-1 gap-x-4 gap-y-2 lg:grid-cols-2">
          <div className="flex flex-col gap-2">
            {routes1.map((route) => (
              <Route key={`${route.start}-${route.end}`} start={route.start} end={route.end} onSelect={onRouteSelect} />
            ))}
          </div>

          <div className="flex flex-col gap-2">
            {routes2.map((route) => (
              <Route key={`${route.start}-${route.end}`} start={route.start} end={route.end} onSelect={onRouteSelect} />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}

export default Info;
