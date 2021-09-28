import './SignupPage.css';
import {useState} from "react";
import {Button, Col, Container, Form, Row, Toast, ToastContainer} from "react-bootstrap";
import {signup} from "../../service/auth-service";

function SignupPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [showError, setShowError] = useState(false);

    const handleSignup = (e) => {
        e.preventDefault();

        setShowError(false);
        if (!email || !password) {
            setShowError(true);
        } else {
            signup(email, password);
        }
    }

    return (
        <div className="signup-page">
            <Container>
                <Row>
                    <Col lg></Col>
                    <Col lg>
                        <div className={"signup-form"}>
                            <Form onSubmit={handleSignup}>
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
                                    Signup
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
    );
}

export default SignupPage;
