import React, { Component } from 'react'
import { Container, Row, Col} from 'react-bootstrap';
import Square from './Square.js'


class Board extends Component {

    constructor(props) {
        super(props)

        let squares = []
        let counter = 0
        for(let i=0; i<3; i++){
            let row = []
            for(let j=0; j<3; j++){
                row.push({text:'', id:counter})
                counter++
            }
            squares.push(row)
        }
        //console.log(squares)

        this.state = {
            turnX: true, // X always goes first
            infoText: "X turn",
            gameOver: false,
            squares: squares
        }

        this.squareClick = this.squareClick.bind(this)
        this.checkGameOver = this.checkGameOver.bind(this)
    }


    checkGameOver(){
        let squares = this.state.squares
        let winCombinations = [
            "012", // horizontal
            "345",
            "678",
            "036", // vertical
            "147",
            "258",
            "048", // diagonal
            "246"
        ]

        let xstr = ""
        let ostr = ""
        for(let i=0; i<3; i++){
            for(let j=0; j<3; j++){
                let square = squares[i][j]
                let idstr = square.id.toString()
                let text = square.text
                if(text === "X"){
                    xstr = xstr + idstr
                }
                if(text === "O"){
                    ostr = ostr + idstr
                }
            }
        }

        //console.log(xstr)
        let gameOver = false
        let winText = ""
        for(let i=0; i<winCombinations.length; i++){
            let combination = winCombinations[i]
            //console.log(combination)
            let xwin = true
            let owin = true
            for (let j=0; j < combination.length; j++) {
                let combChar = combination.charAt(j);
                if(!xstr.includes(combChar)){
                    xwin = false
                }
                if(!ostr.includes(combChar)){
                    owin = false
                }
            }

            if(xwin || owin){
                gameOver = true
                if(xwin){
                    winText = "X wins"
                }else{
                    winText = "O wins"
                }
                break
            }
        }

        if(gameOver){
            this.setState({
                gameOver: true,
                infoText: winText
            })
        }
    }

    squareClick(e){
        let target = e.target
        let id = target.id
        //console.log(id)

        // are we still playing
        if(this.state.gameOver){
            return
        }

        let row = Math.floor(id / 3)
        let col = id % 3
        //console.log(id, row, col)

        let squares = this.state.squares
        let square = squares[row][col]
        //console.log(square)

        let currentText = square['text']
        if(currentText !== ''){
            return // this spot is already taken
        }

        let newText = this.state.turnX ? 'X' : 'O'
        square['text'] = newText

        squares[row][col] = square
        let turnX = !this.state.turnX
        this.setState({
            squares: squares,
            turnX: turnX,
            infoText: turnX ? "X turn" : "O turn"
        })

        this.checkGameOver()
    }

    render() {
        /*
        console.log(this.state.squares)
        this.state.squares.map(row => (
            console.log(row)
        ))

        this.state.squares.map(row => (
            row.map(item => (
                console.log(item)
            ))
        ))
        */

        return (
            <div>
                <Container className="board mt-5">
                {this.state.squares.map(row => (
                    <Row key={row[0].id}>
                    {row.map(item => (
                        <Col key={item.id}>
                            <Square
                                id={item.id}
                                text={item.text}
                                squareClick={this.squareClick}
                            />
                        </Col>
                    ))}
                    </Row>
                ))}
                <h3>{this.state.infoText}</h3>
                </Container>
            </div>
        )
    }
}

export default Board
