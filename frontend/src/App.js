import React from 'react';
import Container from 'react-bootstrap/Container';

import './App.css';
import {LoginPage} from "./components/LoginPage";

function App() {

    return (
        <Container className="p-3">
            <LoginPage/>
        </Container>
    )
}

export default App;
