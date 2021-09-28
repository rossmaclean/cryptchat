import './ChatsPage.css';
import {Col, Container, Row} from "react-bootstrap";
import ChatsList from "../ChatsList/ChatsList";
import ChatDetails from "../ChatDetails/ChatDetails";
import {useEffect, useState} from "react";
import {isAuthSet} from "../../service/auth-service";
import {useHistory} from "react-router-dom";

function ChatsPage() {
    const history = useHistory();
    const [chatId, setChatId] = useState('chat123');
    const userId = "user1";

    const chatChanged = (chatId) => {
        console.log("ChatID now ");
        console.log(chatId)
        setChatId(chatId);
    }

    useEffect(() => {
        isLoggedIn()
    }, []);

    const isLoggedIn = () => {
        console.log("Logged in?");
        if (!isAuthSet()) {
            console.log("No");
            history.push("/login");
        } else {
            console.log("yes");
        }
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
