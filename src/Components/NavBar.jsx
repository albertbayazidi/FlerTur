import { Fragment } from "react";

function Navbar() {
  return (
    <Fragment>
      <nav class="bg-primary">
        <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
          <div class="relative flex h-16 items-center justify-between">
            <div class="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
              <div class="flex flex-shrink-0 items-center">
                <img class="h-8 w-auto" src="/logo.png" alt="FlerTur" />
              </div>
            </div>
            <div>
              <button
                type="button"
                class="relative flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
                id="user-menu-button"
                aria-expanded="false"
                aria-haspopup="true"
              >
                <span class="absolute -inset-1.5"></span>
                <img
                  class="h-8 w-8 rounded-full"
                  src="https://bayazidi.xyz/pb.jpeg"
                  alt=""
                />
              </button>
            </div>
          </div>
        </div>
      </nav>
    </Fragment>
  );
}

export default Navbar;
