import React, { Component } from 'react'
import ReguCompon from './ReguCompon'
import PureCompon from './PureCompon'

/*
Pure component implements shouldComponentUpdate with a
shallow props and state comparison.
*/
class ParentCompon extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             name: "Simon"
        }
    }
    
    componentDidMount(){
        setInterval(() => this.setState({
            name: "Simon"
        }), 2000)
    }

    render() {
        //console.log("*** Parent Component Render ***")
        return (
            <div>
                Parent Component
                <ReguCompon name={this.state.name}></ReguCompon>
                <PureCompon name={this.state.name}></PureCompon>
            </div>
        )
    }
}

export default ParentCompon
