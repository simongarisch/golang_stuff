import React, { Component } from 'react'

// props are immutable
class Welcome extends Component {
    render (){
        const {name, heroName} = this.props  // destructuring props
        return (
            <h1>Welcome {name} a.k.a {heroName}</h1>
        )
    }
}

export default Welcome
