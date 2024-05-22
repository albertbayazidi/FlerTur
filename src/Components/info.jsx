import { Fragment } from "react";

function Info() {
  return (
    <Fragment>
      <div class="mx-10 mt-2 rounded-xl bg-primary px-10 py-5">
        <h1 class="pb-5 text-center text-xl text-white">Info</h1>
        <p class="text-center text-white">
          Denne nettsiden prøver å finne de billigeste bilettprisene en månded
          fram i tid.
        </p>
      </div>
      <img src="/footer_art.png" />
    </Fragment>
  );
}

export default Info;
