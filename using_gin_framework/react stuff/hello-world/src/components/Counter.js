import React, { Component } from 'react'

class Counter extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             count: 0
        }
    }
    
    increment(){
        // don't modify the state directly - changes will not render
        // by using the setState method you can ensure that the browser updates

        /*
        this.setState({  // calls to setState are also asynchronous... we can add a callback function after...
            count: this.state.count+1  // this can bite you - see below
        },
        () => {console.log("Callback value", this.state.count)}  // callback arrow function
        )
        */
    
        // be sure to pass the previous state into a function
        this.setState(prevState => ({
            count: prevState.count + 1
        }))
    }

    render() {
        return (
            <div>
                <div>count - {this.state.count}</div>
                <button onClick={() => this.increment()}>Increment</button>
            </div>
        )
    }
}

export default Counter
