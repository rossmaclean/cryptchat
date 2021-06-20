import React from 'react';
import Container from 'react-bootstrap/Container';

import './App.css';
import {LoginPage} from "./components/LoginPage";
import {BrowserRouter, Route} from "react-router-dom";
import {SignupPage} from "./components/SignupPage";

function App() {

    return (
        <Container className="p-3">
            {/*<LoginPage/>*/}
            <BrowserRouter>
                <main>
                    <nav>
                        <ul>

                            <li><a href="/">Home</a></li>
                            <li><a href="/login">Login</a></li>
                            <li><a href="/signup">Signup</a></li>
                        </ul>
                    </nav>

                    {/*<Route path="/" component={App}/>*/}
                    <Route path="/login" component={LoginPage}/>
                    <Route path="/signup" component={SignupPage}/>
                </main>
            </BrowserRouter>

        </Container>
    )
}

export default App;
