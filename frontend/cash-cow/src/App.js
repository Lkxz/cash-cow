import React, { Component } from 'react';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      records: [],
    };
  }
  componentDidMount() {
    fetch('http://localhost:8000/records').then(response => response.json()).then(records => this.setState({records}))
  }
  render() {
    return (
      <div className="App">
        <h1>Cash 🐮</h1>
        <table>
          <thead>
            <tr>
              <th>Source</th>
              <th>Target</th>
              <th>Amount</th>
            </tr>
          </thead>
          <tbody>
            {this.state.records.map((record, i) => <tr key={i}><td>{record.Source}</td><td>{record.Target}</td><td>{record.Amount}</td></tr>)}
          </tbody>
        </table>
      </div>
    );
  }
}

export default App;
