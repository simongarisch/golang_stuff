import React, { Component } from 'react'


export class Entry extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             input: ""
        }
    }

    resetState = () => {
        this.setState({
            input: ""
        })
    }

    handleChange = (e) => {
        this.setState({ input: e.target.value });
    }

    handleClick = (e) => {
        let note = this.state.input
        if(note !== "" ){
            this.props.addNote(note)
            this.resetState()
        }
        e.preventDefault();
    }

    render() {
        return (
            <div className="container">
                <form>
                    <div className="form-group">
                        <label>Note</label>
                        <input type="text" onChange={this.handleChange} value={this.state.input} className="form-control" />
                        <small className="form-text text-muted">Go on, you must have some more notes to add...</small>
                    </div>
                    <button type="submit" className="btn btn-primary" onClick={this.handleClick}>Submit</button>
                </form>
            </div>
        )
    }
}

export default Entry
