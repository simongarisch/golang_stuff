import React from 'react'

function Person(props) {
    return (
        <div>
            <h2>
                I am {props.person.name}. I know {props.person.skill}
            </h2>
        </div>
    )
}

export default Person
