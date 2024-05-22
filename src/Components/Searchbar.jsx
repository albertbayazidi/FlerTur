import "../App.css";

function SearchBarArea() {
  return (
    <div class="bg-second">
      <div class="grid grid-cols-3 justify-items-center bg-second px-10 py-8">
        <div>
          <input
            class="rounded px-1 py-2 text-primary focus:outline-none"
            type="text"
            placeholder="Fra"
            list="cityname"/>
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
        </div>
        <div>
          <button
            type="button"
            class=" rounded  bg-primary p-2 text-white hover:bg-second hover:ring hover:ring-primary">
            Bytt retning
          </button>
        </div>

        <div>
          <input
            class="rounded px-1 py-2 text-primary focus:outline-none"
            type="text"
            placeholder="Til"
            list="cityname"/>
        </div>
      </div>
      <img src="/search_bar_kunst_remix.png" />
    </div>
  );
}

export default SearchBarArea;
