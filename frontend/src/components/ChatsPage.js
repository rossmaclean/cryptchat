import Jumbotron from "react-bootstrap/Jumbotron";
import Container from "react-bootstrap/Container";
import React from "react";
import * as chatService from "../service/chatService";
import {ChatPage} from "./ChatPage";
import {Col, Row} from "react-bootstrap";
import {ChatCard} from "./ChatCard";

export class ChatsPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {username: "Ross", chats: [1,2,3,4]}
    }

    componentDidMount() {
        console.log('Component mounted');
        chatService.getChats(this.state.username)
            .then(res => res.json())
            .then(json => {
                console.log(json);
                this.setState({chats: json.chats});
            })
            .catch(err => {
                console.log("Unable to get chats");
                console.log(err);
            });
    }

    render() {
        return (
            <Container className="p-3">
                <Jumbotron>
                    {/*<h1 className="header">Chats Cryptchat</h1>*/}
                    <Container>
                        <Row>
                            <Col xs lg="3">
                                Nav Here
                                {this.state.chats.map(chat =>
                                    <ChatCard></ChatCard>
                                )}
                            </Col>
                            <Col>Big Chat</Col>
                        </Row>
                    </Container>

                    {/*{this.state.chats.map(chat =>*/}
                    {/*    <span>*/}
                    {/*        <ul key={chat.chatId}>*/}
                    {/*            <li>User: {chat.user}</li>*/}
                    {/*            <li>Timestamp: {chat.timestamp}</li>*/}
                    {/*            <li>Chat Id: {chat.chatId}</li>*/}
                    {/*        </ul>*/}
                    {/*        <hr/>*/}
                    {/*    </span>*/}
                    {/*)}*/}
                    {/*<ChatPage></ChatPage>*/}
                </Jumbotron>
            </Container>
        )
    }

}