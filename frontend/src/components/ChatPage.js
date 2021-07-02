import React from "react";
import Container from "react-bootstrap/Container";

export class ChatPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            chatMessages: props.chatMessages, me: {
                "userId": "b123",
                "username": "Ross"
            }
        }
        console.log(props.chatMessages)
    }

    render() {
        return (
            <Container>
                <h3 className="header"></h3>
                {this.state.chatMessages.map(message => {
                    // return <li>{message.userId}: {message.message}</li>
                    return <span>
                        <span style={{
                            float: this.state.me.userId === message.userId ? "right" : "left",
                            // margin: "0.1em"
                        }}>{message.username}</span>
                        <span style={{
                            float: this.state.me.userId === message.userId ? "right" : "left",
                            // margin: "0.1em",
                            border: "2px solid lightblue",
                            borderRadius: "0.5em"
                        }}>{message.message}</span>
                        <br/>
                    </span>
                })}
            </Container>
        )
    }

}