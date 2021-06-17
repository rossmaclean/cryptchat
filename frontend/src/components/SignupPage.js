import Jumbotron from "react-bootstrap/Jumbotron";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import React from "react";
import {signup} from "../service/auth-service";

export class SignupPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            confirmUsername: '',
            password: '',
            confirmPassword: ''
        }

        this.handleUsernameChange = this.handleUsernameChange.bind(this);
        this.handleConfirmUsernameChange = this.handleUsernameChange.bind(this);
        this.handlePasswordChange = this.handleConfirmPasswordChange.bind(this);
        this.handleConfirmPasswordChange = this.handleConfirmPasswordChange.bind(this);
        this.handleSignup = this.handleSignup.bind(this);
    }

    handleUsernameChange(event) {
        this.setState({username: event.target.value});
    }

    handleConfirmUsernameChange(event) {
        this.setState({confirmUsername: event.target.value});
    }

    handlePasswordChange(event) {
        this.setState({password: event.target.value});
    }

    handleConfirmPasswordChange(event) {
        this.setState({confirmPassword: event.target.value});
    }

    handleSignup(event) {
        event.preventDefault();
        signup(this.state.username, this.state.password)
    }

    render() {
        return (
            <Container className="p-3">
                <Jumbotron>
                    <h1 className="header">Welcome To Cryptchat</h1>
                    <h2 className="header">Signup</h2>
                    <Form>
                        <Form.Group controlId="formUsername">
                            <Form.Label>Username</Form.Label>
                            <Form.Control value={this.state.username} onChange={this.handleUsernameChange} type="text"
                                          placeholder="Enter Username"/>
                            <Form.Text className="text-muted">
                                We'll never share your information with anyone else.
                            </Form.Text>
                        </Form.Group>

                        <Form.Group controlId="formConfirmUsername">
                            <Form.Label>Confirm Username</Form.Label>
                            <Form.Control value={this.state.confirmUsername} onChange={this.handleConfirmUsernameChange}
                                          type="text"
                                          placeholder="Confirm Username"/>
                        </Form.Group>

                        <Form.Group controlId="formPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control value={this.state.password} onChange={this.handlePasswordChange}
                                          type="password" placeholder="Password"/>
                        </Form.Group>

                        <Form.Group controlId="formConfirmPassword">
                            <Form.Label>Confirm Password</Form.Label>
                            <Form.Control value={this.state.confirmPassword} onChange={this.handleConfirmPasswordChange}
                                          type="password" placeholder="Confirm Password"/>
                        </Form.Group>

                        <Form.Group controlId="formBasicCheckbox">
                            <Form.Check type="checkbox" label="Check me out"/>
                        </Form.Group>

                        <Button variant="primary" onClick={this.handleSignup}>
                            Signup
                        </Button>
                    </Form>
                </Jumbotron>
            </Container>
        )
    }

}