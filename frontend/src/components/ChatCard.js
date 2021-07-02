import React from "react";
import {Card} from "react-bootstrap";

export class ChatCard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {chat: props.chat}
    }

    render() {
        return (
            <Card style={{margin: "5%"}}>
                <Card.Title>{this.state.chat.users[0].username}</Card.Title>
                <Card.Text>{this.state.chat.lastMessage}</Card.Text>
            </Card>
        )
    }

}