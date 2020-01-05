import React from 'react'

/*
function Greet (){
    return <h1>Hello there!</h1>
}
*/

// instead using the arrow function syntax
// convention is to use the word 'props' for props
const Greet = (props) => {
    console.log(props)
    return (
        <div>
            <h1>Hello {props.name}! {props.followUp}</h1>
            {props.children}
        </div>
    )
}

export default Greet