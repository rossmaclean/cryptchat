import './ChatDetails.css';
import logo from '../../logo.svg';
import {Col, Container, Row} from "react-bootstrap";
import {useEffect, useState} from "react";
import {getMessagesForChat} from "../../service/chat-service";

function ChatDetails(props) {
    const [messages, setMessages] = useState([]);

    useEffect(() => {
        getMessages()
    }, []);

    const getMessages = () => {
        getMessagesForChat(props.chatId)
            .then(res => {
                res.json()
                    .then(json => {
                        console.log("Json")
                        console.log(json);
                        setMessages(json.messages);
                    })
                    .catch(err => {
                        console.log(err);
                    })
            })
            .catch(err => {
                console.log(err);
            });
    }

    const getAlignClass = (message) => {
        return message.fromMe ? "from-me" : "from-other";
    }
    return (
        <div className="chat-details">
            {messages.map(message => {
                return (
                    <div className={"message-container"} key={message.messageId}>
                        {message.userId === props.userId ? (
                            <div className={"from-me"} key={message.messageId}>
                                <Container>
                                    <Row>
                                        <Col sm={11}>
                                            {message.message}
                                        </Col>
                                        <Col sm={1}>
                                            <img className={"user-icon"} src={logo}/>
                                        </Col>
                                    </Row>
                                </Container>
                            </div>

                        ) : (
                            <div className={"from-other"}>
                                <Container>
                                    <Row>
                                        <Col sm={1}>
                                            <img className={"user-icon"} src={logo}/>
                                        </Col>
                                        <Col sm={11}>
                                            {message.message}
                                        </Col>
                                    </Row>
                                </Container>
                            </div>
                        )}
                    </div>
                )
            })}
        </div>
    );
}

export default ChatDetails;
