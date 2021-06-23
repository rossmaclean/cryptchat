import Jumbotron from "react-bootstrap/Jumbotron";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import React from "react";
import {login} from "../service/auth-service";

export class ChatPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {username: '', password: ''}

        this.handleUsernameChange = this.handleUsernameChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleLogin = this.handleLogin.bind(this);
    }

    handleUsernameChange(event) {
        this.setState({username: event.target.value});
    }

    handlePasswordChange(event) {
        this.setState({password: event.target.value});
    }

    handleLogin(event) {
        event.preventDefault();
        login(this.state.username, this.state.password);
    }

    render() {
        return (
            <Container className="p-3">
                <Jumbotron>
                    <h1 className="header">Welcome To Cryptchat</h1>
                    <Form>
                        <Form.Group controlId="formUsername">
                            <Form.Label>Username</Form.Label>
                            <Form.Control value={this.state.username} onChange={this.handleUsernameChange} type="text" placeholder="Enter Username"/>
                            <Form.Text className="text-muted">
                                We'll never share your information with anyone else.
                            </Form.Text>
                        </Form.Group>

                        <Form.Group controlId="formPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control value={this.state.password} onChange={this.handlePasswordChange} type="password" placeholder="Password"/>
                        </Form.Group>
                        <Form.Group controlId="formBasicCheckbox">
                            <Form.Check type="checkbox" label="Check me out"/>
                        </Form.Group>
                        <Button variant="primary" onClick={this.handleLogin}>
                            Login
                        </Button>
                    </Form>
                </Jumbotron>
            </Container>
        )
    }

}