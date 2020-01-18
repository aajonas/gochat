import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    
    return (
      <div className="App">
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );

  }

  componentDidMount() {
    connect((msg) => {
        console.log("New Message")
        this.setState(prevState => ({
            chatHistory: [...this.state.chatHistory, msg]
        }))
        console.log(this.state);
    });
 }

}

export default App;

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

// export default App;
