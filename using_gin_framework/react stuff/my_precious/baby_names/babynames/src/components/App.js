import React, { Component } from 'react'
import '../css/App.css';
import SearchBar from './SearchBar';
import NamesView from './NamesView'


class App extends Component {

  constructor(props) {
    super(props)
  
    this.state = {
       searchText: ''
    }

    this.setSearchText = this.setSearchText.bind(this)
    this.getSearchText = this.getSearchText.bind(this)
  }

  setSearchText(searchText){
    this.setState({
      searchText: searchText
    })
  }

  getSearchText(){
    return this.state.searchText
  }

  render() {
    return (
      <div className="App">
        <SearchBar
          setSearchText={this.setSearchText}
        />
        <NamesView
          getSearchText={this.getSearchText}
        />
    </div>
    )
  }
}

export default App;
