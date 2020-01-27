import React from 'react'


const Square = (props) => {
    return (
        <div id={props.id} onClick={props.squareClick} className="square center">
            <button
                className="btn"
            >
                {props.text}
            </button>
        </div>
    )
}


const squareClick = () => {
    console.log("This still needs to be connected")
}


Square.defaultProps = {
    text: '',
    squareClick: squareClick
}

export default Square
