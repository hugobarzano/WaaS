import React, { Component } from "react";
import "./App.css";
import Header from './components/Header/Header';
import CommunityHistory from './components/CommunityHistory/CommunityHistory';


import { connect, sendMsg } from "./api";

class App extends Component {
    constructor(props) {
        super(props);
    }

    send() {
        console.log("hello");
        sendMsg("hello");
    }

    componentDidMount() {
        connect((msg) => {
            console.log("New Message")
            this.setState(prevState => ({
                communityHistory: [...this.state.communityHistory, msg]
            }))
            console.log(this.state);
        });
    }

    render() {
        return (
            <div className="App">
                <Header />
                <CommunityHistory communityHistory={this.state.communityHistory} />
                <button onClick={this.send}>Hit</button>
            </div>
        );
    }
}

export default App;