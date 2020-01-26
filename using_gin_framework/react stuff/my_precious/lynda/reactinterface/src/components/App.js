// npm install --save bootstrap react-icons lodash jquery popper.js moment react-moment
import React, { Component } from 'react';
import '../css/App.css';
import AddAppointments from './AddAppointments';
import ListAppointments from './ListAppointments';
import SearchAppointments from './SearchAppointments';

import {without} from 'lodash'


class App extends Component {

  constructor(props) {
    super(props)
  
    this.state = {
      myAppointments: [],
      lastIndex: 0,
      formDisplay: false
    }

    this.deleteAppointment = this.deleteAppointment.bind(this)
    this.toggleForm = this.toggleForm.bind(this)
    this.addAppointment = this.addAppointment.bind(this)
  }
  
  componentDidMount(){
    fetch('./data.json')
      .then(response => response.json())
      .then(result => {
        const apts = result.map(item => {
          item.aptId = this.state.lastIndex;
          this.setState({ lastIndex: this.state.lastIndex + 1})
          return item;
        })

        this.setState({
          myAppointments: apts
        });
      });
  }

  toggleForm(){
    this.setState({
      formDisplay: !this.state.formDisplay
    })
  }

  addAppointment(apt){
    let tempApts = this.state.myAppointments
    apt.aptId = this.state.lastIndex;
    tempApts.unshift(apt)  // add to appointments array at beginning
    this.setState({
      myAppointments: tempApts,
      lastIndex: this.state.lastIndex + 1
    })
  }

  deleteAppointment(apt){
    //console.log("deleting...")
    let tempApts = this.state.myAppointments
    tempApts = without(tempApts, apt)  // lodash without.js

    this.setState({
      myAppointments: tempApts
    })
  }

  render() {
    return (
      <main className="page bg-white" id="petratings">
        <div className="container">
          <div className="row">
            <div className="col-md-12 bg-white">
              <div className="container">

                <AddAppointments
                  formDisplay={this.state.formDisplay}
                  toggleForm={this.toggleForm}
                  addAppointment={this.addAppointment}
                />

                <SearchAppointments />
    
                <ListAppointments
                  appointments={this.state.myAppointments}
                  deleteAppointment={this.deleteAppointment}
                />

              </div>
            </div>
          </div>
        </div>
      </main>
    );
  }
}

export default App;
