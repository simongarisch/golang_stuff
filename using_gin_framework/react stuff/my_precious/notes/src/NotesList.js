import React from 'react'


function NotesList(props) {

    const notesList = props.notes.map((note, index) =>
        <li key={index}>{note}</li>
    );

    return (
        <div className="container">
            <h3 className="text-secondary mt-5">Your notes:</h3>
            <ul>{notesList}</ul>,
        </div>
    )
}

export default NotesList
