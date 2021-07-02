import Jumbotron from "react-bootstrap/Jumbotron";
import Container from "react-bootstrap/Container";
import React from "react";
import {ChatPage} from "./ChatPage";
import {Col, Row} from "react-bootstrap";
import {ChatCard} from "./ChatCard";

export class ChatsPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            username: "Ross", userId: "Ross", chatInfo: {
                chats: [
                    {
                        "chatId": "cb123",
                        "users": [
                            {
                                "userId": "a123",
                                "username": "Joe"
                            },
                            {
                                "userId": "b123",
                                "username": "Ross"
                            }
                        ],
                        "lastMessage": "Yeah me too!",
                        "timestamp": "2021-02-15:15-00-00"
                    },
                    {
                        "chatId": "cb1234",
                        "users": [
                            {
                                "userId": "b123",
                                "username": "Ross"
                            },
                            {
                                "userId": "c123",
                                "username": "Tom"
                            }
                        ],
                        "lastMessage": "Good thanks, and you?",
                        "timestamp": "2021-02-15:15-00-00"
                    }
                ],
                chatMessages: [
                    {
                        "userId": "c123",
                        "username": "Tom",
                        "chatId": "cb123",
                        "messageId": "abc123",
                        "message": "Hey",
                        "timestamp": "2021-02-15:15-00-00"
                    },
                    {
                        "userId": "b123",
                        "username": "Ross",
                        "chatId": "cb123",
                        "messageId": "abc123",
                        "message": "Hey, how are you?",
                        "timestamp": "2021-02-15:15-01-00"
                    },
                    {
                        "userId": "c123",
                        "username": "Tom",
                        "chatId": "cb123",
                        "messageId": "abc123",
                        "message": "Good thanks, and you?",
                        "timestamp": "2021-02-15:15-02-00"
                    }
                ]
            }
        }
    }

    componentDidMount() {
        console.log('Component mounted');
        // chatService.getChats(this.state.username)
        //     .then(res => res.json())
        //     .then(json => {
        //         console.log(json);
        //         this.setState({chats: json.chats});
        //     })
        //     .catch(err => {
        //         console.log("Unable to get chats");
        //         console.log(err);
        //     });
    }

    render() {
        return (
            <Container className="p-3">
                <Jumbotron>
                    {/*<h1 className="header">Chats Cryptchat</h1>*/}
                    <Container>
                        <Row>
                            <Col xs lg="3">
                                <Jumbotron style={{backgroundColor: "white", padding: "0.2em"}}>
                                    {this.state.chatInfo.chats.map(chat =>
                                        <ChatCard chat={chat}></ChatCard>
                                    )}
                                </Jumbotron>
                            </Col>
                            <Col>
                                <Jumbotron style={{backgroundColor: "white", padding: "0.2em"}}>
                                    <ChatPage chatMessages={this.state.chatInfo.chatMessages}></ChatPage>
                                </Jumbotron>
                            </Col>
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