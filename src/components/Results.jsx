function Results({ results }) {
    const formatDate = (baseDateString, daysToAdd) => {
    const date = new Date(baseDateString);
    date.setDate(date.getDate() + daysToAdd); 
    
    return date.toLocaleDateString("no-NO", { 
      day: "2-digit", 
      month: "short" 
    });
  };

  if (!results || results.length === 0) {
    return (
      <div className="mx-4 mt-4 text-center text-black">
        OBS, Ingen resultater funnet.
      </div>
    );
  }
    console.log(results)

  return (
    <div className="mx-4 mt-4 flex flex-col gap-6 sm:mx-10">
      {results.map((route, index) => (
        <div key={index} className="rounded-xl bg-second p-4 text-white shadow-lg">
          <h2 className="mb-4 text-xl font-bold text-gray-200">
            {route.startStation} ➝ {route.endStation}
          </h2>
          
            <div className={` space-y-4 max-h-[240px] overflow-y-auto `} >
            {route.pageDataResults.map((ticket, i) => (
              <div key={i} className="flex flex-col justify-between rounded-lg
                                bg-primary p-3 sm:flex-row sm:items-center">
                <div>
                    <div className="text-lg font-bold text-white">
                    {formatDate(route.retrievalTime, i)} Kl. {ticket.startTime}
                    </div>
                  <div className="text-sm">Varighet: {ticket.duration}</div>
                  <div className="text-xs">
                    Tog: {ticket.trainIds.join(", ")}
                  </div>
                </div>
                
                <div className="mt-4 flex items-center justify-between gap-4 sm:mt-0">
                  <span className="text-xl ">{ticket.price} kr</span>
                  <a 
                    href={ticket.url} 
                    target="_blank" 
                    rel="noreferrer"
                    className="rounded-lg hover:text-accentColor px-4 py-2
                                        text-sm font-semibold
                                        hover:bg-opacity-90"
                  >
                    Kjøp
                  </a>
                </div>
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
}

export default Results;
