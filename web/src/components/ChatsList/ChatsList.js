import './ChatsList.css';
import logo from '../../logo.svg';
import {useEffect, useState} from "react";
import {getChatsForUser} from '../../service/chat-service'

function ChatsList(props) {

    const [chats, setChats] = useState([]);

    useEffect(() => {
        getChats()
    }, []);

    const getChats = () => {
        getChatsForUser(props.userId)
            .then(res => {
                res.json()
                    .then(json => {
                        console.log(json);
                        setChats(json);
                    })
                    .catch(err => {
                        console.log(err);
                    })
            })
            .catch(err => {
                console.log(err);
            });
    }

    return (
        <div className="chats-list">
            {chats.map(chat => {
                return (
                    <div key={chat.chatId} onClick={props.chatChanged(chat.chatId)}>
                        <button onClick={props.chatChanged}>X</button>
                        <img src={logo}/>
                        <small>{chat.userId[0]}</small>
                        <small>{chat.timestamp}</small>
                        {/*<small>12:32AM</small>*/}
                    </div>
                )
            })}

        </div>
    );
}

export default ChatsList;
