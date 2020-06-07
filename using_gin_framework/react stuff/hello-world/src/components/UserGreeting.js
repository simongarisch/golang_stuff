import React, { Component } from 'react'

class UserGreeting extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             isLoggenIn: true
        }
    }
    

    render() {
        /*
        if (this.state.isLoggenIn){
            return (
                <div>
                    Welcome Simon, thanks for logging in
                </div>
            )
        } else {
            return (
                <div>
                    Welcome Guest
                </div>
            )
        }
        */

        // alternatively
        /*
        let message
        if(this.state.isLoggenIn){
            message = <div>Welcome Simon, thanks for logging in</div>
        }else{
            message = <div>Welcome Guest</div>
        }
        return <div>{message}</div>
        */

        // yet another alternative
        return (
            this.state.isLoggenIn ?
            <div>Welcome Simon, thanks for logging in</div> :
            <div>Welome Guest</div>
        )

    }
}

export default UserGreeting
