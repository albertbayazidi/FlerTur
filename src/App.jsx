import { useState } from "react";
import Navbar from "./components/Navbar";
import SearchBarArea from "./components/SearchBarArea";
import Info from "./components/Info";
import Results from "./components/Results";
import Loading from "./components/Loading";
import "./App.css";

function App() {
  const [results, setResults] = useState(null);
  const [loading, setLoading] = useState(false);
  const [hasSearched, setHasSearched] = useState(false); 

  const executeSearch = async (from, to) => {
    setLoading(true);
    setHasSearched(true);
    setResults(null); 

    try {
      const response = await fetch(
        `http://localhost:3001/api/search?from=${encodeURIComponent(from)}&to=${encodeURIComponent(to)}`
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
    return <Info />; 
  };

  return (
    <div className="flex min-h-screen flex-col">
      <Navbar />
      
      <SearchBarArea onSearch={executeSearch} />
      
      {renderContent()}

      <img
        src="/footer_art.png"
        className="mt-auto h-auto w-auto object-contain md:h-auto md:w-full"
        alt=""
      />
    </div>
  );
}

export default App;
