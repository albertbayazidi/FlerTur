import { Fragment } from "react";

function Navbar() {
    function refreshPage(){ window.location.reload(false);}
  return (
    <Fragment>
      <nav class="bg-primary">
        <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
          <div class="relative flex h-16 items-center justify-between">
            <div class="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
              <div class="flex flex-shrink-0 items-center">
                <button onClick={refreshPage}>
                  <img class="h-8 w-auto" src="/logo.png" alt="FlerTur" />
                </button>
              </div>
            </div>
            <div>
              <button
                type="button"
                class="relative flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
                id="user-menu-button"
                aria-expanded="false"
                aria-haspopup="true">
                <a href="https://bayazidi.xyz"  target="_blank" >
                <img
                  class="h-8 w-8 rounded-full"
                  src="https://bayazidi.xyz/pb.jpeg"
                  alt="profile image"
                />
                </a>
              </button>
            </div>
          </div>
        </div>
      </nav>
    </Fragment>
  );
}

export default Navbar;
