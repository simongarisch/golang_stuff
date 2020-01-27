import React, { Component } from 'react'
import { FaSearch } from 'react-icons/fa'


class SearchBar extends Component {

    constructor(props) {
        super(props)
    
        this.setSearchText = props.setSearchText

        this.state = {
             searchText: ""
        }
    }
    
    handleChange = (e) => {
        const searchText = e.target.value
        this.setState({
            searchText: searchText
        })
        this.setSearchText(searchText)
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
