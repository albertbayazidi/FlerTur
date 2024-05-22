import { Fragment } from "react"

function Info() {
  return (
    <Fragment>
    <div class = 'py-5 px-10 mt-2 mx-10 bg-primary rounded-xl'>
        <h1 class = 'text-white text-center text-xl pb-5'>Info</h1>
        <p class = 'text-white text-center'> 
        Denne nettsiden prøver å finne de billigeste bilettprisene en månded fram i tid.</p>
    </div>
    <img src='/footer_art.png' />
    </Fragment>
  );
}

export default Info;