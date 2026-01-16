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
        <input
          className="w-full rounded px-3 py-3 text-primary focus:outline-none"
          type="text"
          placeholder="Fra"
          list="cityname"
          value={from}
          onChange={(e) => setFrom(e.target.value)}
        />

        <button
          type="button"
          onClick={swapStations} 
          className="w-full rounded bg-primary py-3 text-white hover:bg-second hover:ring hover:ring-primary sm:w-auto"
        >
          Bytt retning
        </button>

        <input
          className="w-full rounded px-3 py-3 text-primary focus:outline-none"
          type="text"
          placeholder="Til"
          list="cityname"
          value={to}
          onChange={(e) => setTo(e.target.value)}
        />

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

      <datalist id="cityname">
        <option value="Oslo S" />
        <option value="Trondheim S" />
        <option value="Lillehammer stasjon" />
        <option value="Stavanger stasjon" />
        <option value="Bergen stasjon" />
        <option value="Lillestrøm stasjon" />
        <option value="Fredrikstad stasjon" />
        <option value="Kristiansand stasjon" />
        <option value="Kongsvinger stasjon" />
      </datalist>

      <img
        src="/search_bar_kunst_remix.png"
        className="w-full object-contain"
        alt=""
      />
    </div>
  );
}

export default SearchBarArea;
