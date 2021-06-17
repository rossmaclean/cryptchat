export const signup = function (username, confirmUsername, password, confirmPassword) {
    const data = {
        username: username,
        confirmUsername: confirmUsername,
        password: password,
        confirmPassword: confirmPassword
    }
    fetch('/api/v1/signup', {
        method: 'POST',
        body: JSON.stringify(data)
    }).then(r => console.log(r));
}

export const login = function (username, password) {
    const data = {
        username: username,
        password: password
    }
    fetch('/api/v1/login', {
        method: 'POST',
        body: JSON.stringify(data)
    }).then(r => console.log(r));
}

