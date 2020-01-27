import React, { Component } from 'react'


function getNames(){
    let names = []
    let json = require('./data.json')
    let genders = Object.keys(json)

    let counter = 0
    for (let genderIndex in genders) {
        let gender = genders[genderIndex]
        let genderNames = json[gender]
        for (let nameIndex in genderNames) {
            counter++
            let name = genderNames[nameIndex]
            names.push({'id': counter, 'name': name, 'gender': gender})
        }
    }
    return names
}


export default class NamesView extends Component {
    constructor(props) {
        super(props)
    
        this.getSearchText = props.getSearchText
        this.state = {
            names: []   
        }
    }
    
    componentDidMount(){
        let names = getNames()
        //console.log(names[0])
        this.setState({
            names: names
        })
    }

    render() {
        let searchText = this.getSearchText().toLowerCase()
        //console.log(searchText)

        let filteredNames = this.state.names
        filteredNames = filteredNames.sort((a, b) => {
            if(a.name.toLowerCase() < b.name.toLowerCase()){
                return -1
            }else{
                return +1
            }
        }).filter(item => {
            return (
                item.name.toLowerCase().includes(searchText)
            )
        })

        return (
            <div className='container mt-5'>
                {filteredNames.map(item => (
                    <span 
                        className={
                            'badge mx-1 ' + 
                            (item.gender === 'boy' ? 'badge-primary': 'badge-success')
                        }
                        key={item.id}
                    >
                        {item.name}
                    </span>
                )
                )}
            </div>
        )
    }
}
