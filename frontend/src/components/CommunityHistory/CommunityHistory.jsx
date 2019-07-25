import React, { Component } from "react";
import "./CommunityHistory.scss";

class CommunityHistory extends Component {
    render() {
        const messages = this.props.communityHistory.map((msg, index) => (
            <p key={index}>{msg.data}</p>
        ));

        return (
            <div className="CommunityHistory">
                <h2>History</h2>
                {messages}
            </div>
        );
    }
}

export default CommunityHistory;