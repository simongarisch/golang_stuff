import React, { Component } from 'react'
import { FaSearch } from 'react-icons/fa'


class SearchBar extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             searchText: ""
        }
    }
    
    handleChange = (e) => {
        this.setState({
            searchText: e.target.value
        })
    }

    render() {
        return (
            <div className="container mt-5 h-100">
                <div className="d-flex justify-content-center h-100">
                    <div className="searchbar" style={{color: 'tomato'}}>
                    <input 
                        className="search_input"
                        type="text"
                        placeholder="Search..."
                        onChange={this.handleChange}
                        value={this.state.searchText}
                    />
                    < FaSearch size={20} />
                    </div>
                </div>
            </div>
        )
    }
}

export default SearchBar
