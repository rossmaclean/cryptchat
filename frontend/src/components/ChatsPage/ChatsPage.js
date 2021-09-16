import './ChatsPage.css';
import {Col, Container, Row} from "react-bootstrap";
import ChatsList from "../ChatsList/ChatsList";
import ChatDetails from "../ChatDetails/ChatDetails";
import {useState} from "react";

function ChatsPage() {

    const [chatId, setChatId] = useState('chat123');
    const userId = "user1";

    const chatChanged = (chatId) => {
        console.log("ChatID now ");
        console.log(chatId)
        setChatId(chatId);
    }

    return (
        <div className="chats-page">
            <Container>
                <Row>
                    <Col lg></Col>
                    <Col lg={11}>
                        <div className={"chats-container"}>
                            <Container>
                                <Row>
                                    <Col sm={3}><ChatsList
                                        userId={userId}
                                        chatChanged={chatChanged}
                                    /></Col>
                                    <Col sm={9}><ChatDetails
                                        userId={userId}
                                        chatId={chatId}
                                    /></Col>
                                </Row>
                            </Container>
                        </div>
                    </Col>
                    <Col lg></Col>
                </Row>
            </Container>
        </div>
    );
}

export default ChatsPage;
