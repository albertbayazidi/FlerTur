import { Fragment } from "react"

function Info() {
  return (
    <Fragment>
    <div class = 'p-10 mt-5 mx-10 bg-primary rounded-xl'>
        <h1 class = 'text-white text-center text-xl pb-5'>Info</h1>
        <p class = 'text-white text-center'> Denne nettsiden prøver å finne de billigeste bilettprisene pr dag i en månded og framstiller dem i en priskalender.</p>
    </div>
    <img src='/footer_art.png' />
    </Fragment>
  );
}

export default Info;