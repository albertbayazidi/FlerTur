import GithubIcon from "./custom_svg/GithubIcon";

function Navbar() {
  return (
    <nav className="bg-second">
      <div className="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
        <div className="relative flex h-16 items-center justify-between">
          <a href="/">
            <img className="h-8 w-auto" src="/dark_logo.png" alt="FlerTur" />
          </a>
          <a
            className="text-white"
            href="https://github.com/albertbayazidi/Flertur"
            target="_blank"
            rel="noopener noreferrer"
          >
            <GithubIcon />
          </a>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
