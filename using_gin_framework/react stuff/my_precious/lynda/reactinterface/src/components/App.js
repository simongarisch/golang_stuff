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
      formDisplay: false,
      orderBy: 'petName',
      orderDir: 'asc',
      queryText: '',
      lastIndex: 0
    }

    this.deleteAppointment = this.deleteAppointment.bind(this)
    this.toggleForm = this.toggleForm.bind(this)
    this.addAppointment = this.addAppointment.bind(this)
    this.changeOrder = this.changeOrder.bind(this)
    this.searchApts = this.searchApts.bind(this)
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

  changeOrder(order, dir){
    this.setState({
      orderBy: order,
      orderDir: dir
    })
  }

  searchApts(query){
    this.setState({
      queryText: query
    })
  }

  render() {

    let order
    let filteredApts = this.state.myAppointments
    if(this.state.orderDir === 'asc'){
      order = +1
    }else{
      order = -1
    }

    filteredApts = filteredApts.sort((a, b) => {
      if (a[this.state.orderBy].toLowerCase() <
          b[this.state.orderBy].toLowerCase()
      ){
        return -1 * order
      }else{
        return +1 * order
      }
    }).filter(eachItem => {
      return (
        eachItem['petName']
        .toLowerCase()
        .includes(this.state.queryText.toLowerCase()) ||
        eachItem['ownerName']
        .toLowerCase()
        .includes(this.state.queryText.toLowerCase()) ||
        eachItem['aptNotes']
        .toLowerCase()
        .includes(this.state.queryText.toLowerCase())
      )
    })

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

                <SearchAppointments
                  orderBy={this.state.orderBy}
                  orderDir={this.state.orderDir}
                  changeOrder={this.changeOrder}
                  searchApts={this.searchApts}
                />
    
                <ListAppointments
                  appointments={filteredApts}
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
