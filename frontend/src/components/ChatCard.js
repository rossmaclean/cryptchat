import React from "react";
import {Card} from "react-bootstrap";

export class ChatCard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {chat: props.chat}
    }

    render() {
        return (
            <Card>
                <Card.Title>{this.state.chat.userId}</Card.Title>
                <Card.Text>{this.state.chat.message}?</Card.Text>
            </Card>
        )
    }

}