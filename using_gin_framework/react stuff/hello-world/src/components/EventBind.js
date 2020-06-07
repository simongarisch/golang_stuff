import React, { Component } from 'react'

class EventBind extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             message: "Hello"
        }

        this.clickHandler = this.clickHandler.bind(this)  // best option
    }
    
    clickHandler(){
        // we need to use 'bind' for onClick else this is undefined
        // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/this
        this.setState({
            message: "Yoohoo"
        })
    }

    render() {
        // there are a few ways to bind this
        // binding in the constructor is suggested in the react docs
        return (
            <div>
                <div>{this.state.message}</div>
                {/*<button onClick={this.clickHandler.bind(this)}>Click</button>*/}
                {/*<button onClick={() => this.clickHandler()}>Click</button>*/}
                <button onClick={this.clickHandler}>Click</button>
            </div>
        )
    }
}

export default EventBind
