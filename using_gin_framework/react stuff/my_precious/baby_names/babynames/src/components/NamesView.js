import React, { Component } from 'react'


export default class NamesView extends Component {
    constructor(props) {
        super(props)
    
        this.state = {
             
        }
    }
    
    componentDidMount(){
        fetch("./data.json")
        .then(
            function(response) {
              if (response.status !== 200) {
                console.log('Looks like there was a problem. Status Code: ' + response.status);
                return;
              }
        
              // Examine the text in the response
              response.json().then(function(data) {
                console.log(data);
              });
            }
        )
        .catch(function(err) {
            console.log('Fetch Error ', err);
        });
    }

    render() {
        return (
            <div>
                
            </div>
        )
    }
}
