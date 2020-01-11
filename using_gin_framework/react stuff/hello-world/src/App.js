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
import UserGreeting from './components/UserGreeting';
import NameList from './components/NameList';
import Stylesheets from './components/Stylesheets';
import Form from './components/Form';
import LifecycleA from './components/LifecycleA'

// importing some style sheets
import "./appStyles.css"
import styles from "./appStyles.module.css"
import FragmentDemo from './components/FragmentDemo';
import PureCompon from './components/PureCompon';
import ParentCompon from './components/ParentCompon'

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
      <hr></hr>
  
      <UserGreeting></UserGreeting>
      <hr></hr>

      <NameList></NameList>
      <hr></hr>

      <Stylesheets primary={true}></Stylesheets>
      <hr></hr>

      <h2 className="error">Error Test</h2>
      <h2 className={styles.success}>And success</h2>
      <hr></hr>

      <Form></Form>
      <hr></hr>

      <LifecycleA></LifecycleA>
      <hr></hr>

      <FragmentDemo></FragmentDemo>
      <hr></hr>

      <PureCompon></PureCompon>
      <ParentCompon></ParentCompon>
      <hr></hr>
    </div>
  );
}

export default App;
