import React from 'react';
//import logo from './logo.svg';
import './App.css';

// See ES7 React/Redux/GraphQL/React-Native snippets

// Functional vs. Class components
// Go with functional components if you can...
// If you need state then use class components.
// However, with hooks we can introdude state into functional components
// as of React v16.7.0-alpha...

// JavaScript XML (JSX) is an extension to the JavaScript language syntax

import Greet from './components/Greet'
import Welcome from './components/Welcome'
import Message from './components/Message'
import Counter from './components/Counter'
import FunctionClick from './components/FunctionClick'
import ClassClick from './components/ClassClick'
import EventBind from './components/EventBind'
import ParentComponent from './components/ParentComponent';


function App() {
  return (
    <div className="App">
      <Greet name="Simon" followUp="Where you at"></Greet>
      <Greet name="James" children="for the children section"></Greet>
      <Greet name="Ace" followUp="Of Spades"></Greet>

      <button>Ready!</button>

      <Welcome name="Bruce" heroName="Batman"></Welcome>
      <Welcome name="Clark" heroName="Superman"></Welcome>
      <Welcome name="Diana" heroName="Wonder Woman"></Welcome>
      <hr></hr>
  
      <Message></Message>
      <hr></hr>

      <Counter></Counter>
      <hr></hr>

      <FunctionClick></FunctionClick>
      <ClassClick></ClassClick>
      <hr></hr>

      <EventBind></EventBind>
      <hr></hr>

      <ParentComponent></ParentComponent>
    </div>
  );
}

export default App;
