import React, { Component } from 'react'
import Header from './Header';
import Entry from './Entry';
import NotesList from './NotesList';

export class Notes extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             notes: []
        }
    }

    addNote = (note) => {
        this.setState(prevState => {
            let notes = prevState.notes
            notes.push(note)
            return {
                notes: notes
            }
          });
    }

    render() {
        return (
            <div>
                <Header />
                <Entry addNote={this.addNote}/>
                <NotesList notes={this.state.notes}/>
            </div>
        )
    }
}

export default Notes
