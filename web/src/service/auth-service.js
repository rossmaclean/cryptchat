const authField = 'auth';
const baseUrl = '/api/v1';

export const signup = (email, password) => {
    const url = baseUrl + '/signup'
    const payload = {
        email: email,
        password: password
    }

    return fetch(url, {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    });
}

export const login = (email, password) => {
    const url = baseUrl + '/login'
    const payload = {
        email: email,
        password: password
    }

    return fetch(url, {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    }).then(res => {
        res.json().then(auth => {
            setAuth(auth);
        });
    }).catch(err => {
        console.log(err);
    });
}

export const setAuth = (auth) => {
    auth.token = 'Bearer ' + auth.token
    localStorage.setItem(authField, auth);
}

export const getCurrentAuth = () => {
    return localStorage.getItem(authField);
}

export const isAuthSet = () => {
    return !!localStorage.getItem(authField);
}