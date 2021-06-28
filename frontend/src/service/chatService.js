export const getChats = function (userId) {
    const data = {userId: userId}
    return fetch('/api/v1/chats', {
        method: 'POST',
        body: JSON.stringify(data)
    });
}
