import { useState } from "react"; 

function SearchBarArea({ onSearch }) { 
  const [from, setFrom] = useState("");
  const [to, setTo] = useState("");

  const handleSearchClick = () => {
    if(from && to) {
        onSearch(from, to);
    }
  };

  const swapStations = () => {
    const temp = from;
    setFrom(to);
    setTo(temp);
  };

  return (
    <div className="bg-second">
      <div className="mx-4 grid gap-4 pt-6 sm:mx-10 sm:grid-cols-3 sm:items-center">
        <select
          className="w-full rounded px-3 py-3 text-primary focus:outline-none"
          type="text"
          placeholder="Fra"
          list="cityname"
          value={from}
          onChange={(e) => setFrom(e.target.value)}
        >
            <option value="Oslo S">Oslo S</option>
            <option value="Trondheim S" >Trondheim S</option>
            <option value="Stavanger stasjon">Stavanger stasjon</option>
            <option value="Bergen stasjon">Bergen stasjon</option>
            <option value="Fredrikstad stasjon" >Fredrikstad stasjon</option>
            <option value="Kristiansand stasjon" >Kristiansand stasjon</option>
        </select>

        <button
          type="button"
          onClick={swapStations} 
          className="w-full rounded bg-primary py-3 text-white hover:bg-second hover:ring hover:ring-primary sm:w-auto"
        >
          Bytt retning
        </button>

        <select
          className="w-full rounded px-3 py-3 text-primary focus:outline-none"
          type="text"
          placeholder="Til"
          list="cityname"
          value={to}
          onChange={(e) => setTo(e.target.value)}
        >
            <option value="Oslo S">Oslo S</option>
            <option value="Trondheim S" >Trondheim S</option>
            <option value="Stavanger stasjon">Stavanger stasjon</option>
            <option value="Bergen stasjon">Bergen stasjon</option>
            <option value="Fredrikstad stasjon" >Fredrikstad stasjon</option>
            <option value="Kristiansand stasjon" >Kristiansand stasjon</option>
        </select>


        {/* I cant be bother finding out the real way now */}
        <div></div> 

        <div>
          <button
            type="button"
            className="w-full rounded bg-primary py-3 text-white hover:bg-second hover:ring hover:ring-primary"
            onClick={handleSearchClick} 
          >
            Søk
          </button>
        </div>

        <div></div>
      </div>


      <img
        src="/search_bar_kunst_remix.png"
        className="w-full pt-1 object-contain"
        alt=""
      />
    </div>
  );
}

export default SearchBarArea;
