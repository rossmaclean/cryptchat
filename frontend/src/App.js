import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import LoginPage from "./components/LoginPage/LoginPage";
import SignupPage from "./components/SignupPage/SignupPage";
import {BrowserRouter, Link, Route, Switch} from "react-router-dom";
import ChatsPage from "./components/ChatsPage/ChatsPage";

function App() {
    return (
        <div>
            <h3>Cryptchat</h3>
            <BrowserRouter>
                <Link to="/home">Home</Link>
                <Link to="/login">Login</Link>
                <Link to="/signup">Signup</Link>
                <Link to="/chats">Chats</Link>
                <br/><br/>
                <Switch>
                    {/*<Route exact path="/" component={Home} />*/}
                    <Route path="/login" component={LoginPage}/>
                    <Route path="/signup" component={SignupPage}/>
                    <Route path="/chats" component={ChatsPage}/>
                </Switch>
                <br/><br/>
            </BrowserRouter>
        </div>
    );
}

export default App;
