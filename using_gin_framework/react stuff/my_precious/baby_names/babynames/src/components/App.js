import React from 'react';
import '../css/App.css';
import SearchBar from './SearchBar';
import NamesView from './NamesView'

function App() {
  return (
    <div className="App">
      <SearchBar />
      <NamesView />
    </div>
  );
}

export default App;
