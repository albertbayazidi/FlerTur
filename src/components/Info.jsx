import { Fragment } from "react";

function Info() {
  return (
    <Fragment>
      <div className="mx-4 mt-4 rounded-xl bg-second px-6 py-6 sm:mx-10 sm:px-10">
        <h1 className="pb-4 text-center text-xl text-white">Info:</h1>
        <p className="text-center text-white">
          Denne nettsiden prøver å finne de billigste billettprisene to uker fram i tid.
        </p>
        <p className="text-center text-white">
        Jeg tjener ikke noe på salget av billetene. Du kan finne hele prosjektet på github.
        </p>

      </div>
    </Fragment>
  );
}

export default Info;
