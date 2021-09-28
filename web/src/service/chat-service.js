// const baseUrl = 'localhost:8000/api/v1';
import {getCurrentAuth} from "./auth-service";

const baseUrl = '/api/v1';

export const getChatsForUser = (userId) => {
    const url = baseUrl + '/chats'
    const payload = {userId: userId}

    return fetch(url, {
        method: 'POST',
        headers: {
            Authorization: getCurrentAuth(),
            Accept: 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    });
}

export const getMessagesForChat = (chatId) => {
    const url = baseUrl + '/chat_messages'
    const payload = {chatId: chatId}

    return fetch(url, {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    });
}