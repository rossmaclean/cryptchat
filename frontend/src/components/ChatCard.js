import Jumbotron from "react-bootstrap/Jumbotron";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import React from "react";
import {login} from "../service/auth-service";
import {Card} from "react-bootstrap";

export class ChatCard extends React.Component {
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
            <Card>
                <Card.Title>Joe Bloggs</Card.Title>
                <Card.Text>How man, how are you?</Card.Text>
            </Card>
        )
    }

}