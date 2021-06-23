export const getChats = function (username) {
    const data = {username: username}
    return fetch('/api/v1/chats', {
        method: 'POST',
        body: JSON.stringify(data)
    });
}
