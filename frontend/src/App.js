import React, { Component } from 'react';
import './App.css';
import { Input, Radio, Table, Divider, Tag } from 'antd';
import { BrowserRouter as Router, Route, Link, Switch } from "react-router-dom";
import ServiceList from "./page/serviceList";
import ConfigList from "./page/configList";
import MainLayout from "./layout/mainLayout"

class App extends Component {
  render() {
    return (
      <div className="App">
        <Router>
          <div>
            <Switch>
              <Route exact path="/" render={() => { return <MainLayout><ServiceList /></MainLayout> }} />
              <Route path="/config" render={() => { return <MainLayout><ConfigList /></MainLayout> }} />
            </Switch>
          </div>
        </Router>
      </div>
    )
  }
}

export default App;
