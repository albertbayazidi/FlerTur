import { Fragment } from "react";
import Route from "./Route.jsx";

function Info() {
    const routes1 = [
        { start: "oslo s", end: "bergen stasjon" },
        { start: "oslo s", end: "stavanger stasjon" },
        { start: "oslo s", end: "trondheim s" },
        { start: "bergen stasjon", end: "fredrikstad stasjon" },
    ];

    const routes2 = [
        { start: "bergen stasjon", end: "kristiansand stasjon" },
        { start: "stavanger stasjon", end: "bergen stasjon" },
        { start: "stavanger stasjon", end: "fredrikstad stasjon" },
        { start: "stavanger stasjon", end: "kristiansand stasjon" },
    ];

  return (
    <Fragment>
      <div className="mx-auto mt-4 w-5/6 md:mx-auto grid grid-cols-1 md:grid-cols-2 gap-8">
        <div className ="w-full mx-auto text-lg px-6 py-6 rounded-xl bg-second ">
            <h1 className="mb-8 text-3xl text-center font-bold text-white">Info:</h1>

            <p className="text-center text-white ">
                <strong>NB!</strong> Denne nettsiden er under konstruksjon.
                        Tallene oppdateres én gang hvert 30. minutt. Det betyr
                        at priser og tilgjengelighet i enkelte tilfeller kan
                        avvike noe fra sanntidsinformasjonen hos EnTUR. Dataen
                        er hentet med en GO-crawleren/scraperen. 
            </p>

            <p className="text-center text-white">
              Denne nettsiden prøver å finne de billigste billettprisene en uke
                        fram i tid. Jeg tjener ikke noe på salget av billetene.
                        Du kan finne hele prosjektet på github og en forklaring
                        på hvorfor den bare ser en uke fram i tid.
            </p>
        </div>

        <div className ="w-full mx-auto text-lg px-6 py-6 rounded-xl bg-second">
            <h1 className="mb-8 text-3xl text-center font-bold text-white">Tilgjengelige Ruter:</h1>
            <div className="grid grid-cols-1 gap-x-4 gap-y-2 lg:grid-cols-2">
                <div className="flex flex-col gap-2">
                  {routes1.map((route, index) => (
                    <Route key={`r1-${index}`} start={route.start} end={route.end} />
                  ))}
                </div>

                <div className="flex flex-col gap-2">
                  {routes2.map((route, index) => (
                    <Route key={`r2-${index}`} start={route.start} end={route.end} />
                  ))}
                </div>
          </div>
        </div>

      </div>
    </Fragment>
  );
}

export default Info;
