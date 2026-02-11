import { useState } from "react";
import Info from "./components/Info";
import Loading from "./components/Loading";
import Navbar from "./components/Navbar";
import Results from "./components/Results";
import SearchBarArea from "./components/SearchBarArea";
import "./App.css";

function App() {
  const [results, setResults] = useState(null);
  const [loading, setLoading] = useState(false);
  const [hasSearched, setHasSearched] = useState(false);

  const [from, setFrom] = useState("Trondheim S");
  const [to, setTo] = useState("Oslo S");

  const handleRouteSelect = (start, end) => {
    setFrom(start);
    setTo(end);
  };

  const executeSearch = async (searchFrom, searchTo) => {
    setLoading(true);
    setHasSearched(true);
    setResults(null);

    try {
      const response = await fetch(
        `http://localhost:3001/api/search?from=${encodeURIComponent(searchFrom)}&to=${encodeURIComponent(searchTo)}`,
      );

      if (!response.ok) throw new Error("Nooooooo :(");

      const data = await response.json();
      setResults(data);
    } catch (error) {
      console.error("error:", error);
    } finally {
      setLoading(false);
    }
  };

  const renderContent = () => {
    if (loading) return <Loading />;
    if (hasSearched && results) return <Results results={results} />;
    
    return <Info onRouteSelect={handleRouteSelect} />;
  };

  return (
    <div className="flex min-h-screen flex-col">
      <Navbar />

      <SearchBarArea 
        onSearch={executeSearch} 
        from={from} 
        setFrom={setFrom} 
        to={to} 
        setTo={setTo} 
      />

      {renderContent()}

      <img src="/footer_art.png" className="mt-auto h-auto w-auto object-contain md:h-auto md:w-full" alt="" />
    </div>
  );
}

export default App;
