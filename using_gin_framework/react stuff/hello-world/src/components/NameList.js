import React from 'react'
import Person from './Person'

function NameList() {

    //const names = ["Bruce", "Clark", "Diana"]

    const persons = [
        {
            id: 1,
            name: "Bruce",
            age: 30,
            skill: "React"
        },
        {
            id: 2,
            name: "Clark",
            age: 25,
            skill: "Angular"
        }
    ]

    //return (
        /*
        <div>
            <h2>{names[0]}</h2>
            <h2>{names[1]}</h2>
            <h2>{names[2]}</h2>
        </div>
        */

        /*
        <div>
            {
                names.map(name => <h2>{name}</h2>)
            }
        </div>
        */
    //)

    const personList = persons.map(person => <Person key={person.id} person={person}></Person>)
    return <div>{personList}</div>
}

export default NameList
