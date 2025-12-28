import "./App.css";
import SearchBarArea from "./Components/Searchbar.jsx";
import Navbar from "./Components/NavBar.jsx";
import Info from "./Components/info.jsx";

function App() {
  return (
    <div className="flex min-h-screen flex-col">
      <Navbar />
      <SearchBarArea />
      <Info />
      <img
        src="/footer_art.png"
        className="mt-auto h-auto w-auto md:h-auto md:w-full object-contain"
        alt=""
      />
    </div>
  );
}

export default App;
