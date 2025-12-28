import { useState } from "react";
import "../App.css";

function SearchBarArea() {
  const [date, setDate] = useState(""); // state to store the date

  return (
    <div className="bg-second">
      <div className="mx-4 grid gap-4 py-6 sm:mx-10 sm:grid-cols-3 sm:items-center">
        <input
          className="w-full rounded px-3 py-3 text-primary focus:outline-none"
          type="text"
          placeholder="Fra"
          list="cityname"
        />

        <button
          type="button"
          className="w-full rounded bg-primary py-3 text-white hover:bg-second hover:ring hover:ring-primary sm:w-auto"
        >
          Bytt retning
        </button>

        <input
          className="w-full rounded px-3 py-3 text-primary focus:outline-none"
          type="text"
          placeholder="Til"
          list="cityname"
        />
      </div>

      <div className="mx-4 grid gap-4 pb-4 sm:mx-10 sm:grid-cols-3 sm:items-center">
        <input
          type="date"
          className="w-full rounded px-3 py-3 text-primary focus:outline-none"
          value={date}
          onChange={(e) => setDate(e.target.value)} // update state when user selects a date
        />

        <div className="sm:col-span-2">
          <button
            type="button"
            className="w-full rounded bg-primary py-3 text-white hover:bg-second hover:ring hover:ring-primary"
            onClick={() => console.log(date)} // print the selected date
          >
            Søk
          </button>
        </div>
      </div>

      <datalist id="cityname">
        <option value="Oslo%20S" />
        <option value="Trondheim%20S" />
        <option value="Lillehammer%20stasjon" />
        <option value="Stavanger%20stasjon" />
        <option value="Bergen%20stasjon" />
        <option value="Lillestrøm%20stasjon" />
        <option value="Fredrikstad%20stasjon" />
        <option value="Kristiansand%20stasjon" />
        <option value="Kongsvinger%20stasjon" />
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
