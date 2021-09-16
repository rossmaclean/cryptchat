import './LoginPage.css';
import {Button, Col, Container, Form, Row, Toast, ToastContainer} from "react-bootstrap";
import {useState} from "react";

function LoginPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [showError, setShowError] = useState(false);

    const handleLogin = (e) => {
        e.preventDefault();

        setShowError(false);
        if (!email || !password) {
            setShowError(true);
        } else {
            console.log("Logging in with email " + email + " and password " + password);
        }
    }

    return (
        <div className="login-page">
            <Container>
                <Row>
                    <Col lg></Col>
                    <Col lg>
                        <div className={"login-form"}>
                            <Form onSubmit={handleLogin}>
                                <Form.Group className="mb-3" controlId="formEmail">
                                    <Form.Label>Email</Form.Label>
                                    <Form.Control type="email" placeholder="Email"
                                                  onChange={e => setEmail(e.target.value)}/>
                                    <Form.Text className="text-muted">
                                        We'll never share your email with anyone else.
                                    </Form.Text>
                                </Form.Group>

                                <Form.Group className="mb-3" controlId="formPassword">
                                    <Form.Label>Password</Form.Label>
                                    <Form.Control type="password" placeholder="Password"
                                                  onChange={e => setPassword(e.target.value)}/>
                                </Form.Group>
                                <Button variant="primary" type="submit">
                                    Login
                                </Button>
                            </Form>
                        </div>
                    </Col>
                    <Col lg></Col>
                </Row>
            </Container>

            <ToastContainer position={"top-end"}>
                <Toast bg={"danger"} show={showError} onClick={() => setShowError(false)}>
                    <Toast.Header>
                        <strong className="me-auto">Cryptchat</strong>
                    </Toast.Header>
                    <Toast.Body className={"danger"}>Email and password are required</Toast.Body>
                </Toast>
            </ToastContainer>
        </div>
    )
        ;
}

export default LoginPage;
