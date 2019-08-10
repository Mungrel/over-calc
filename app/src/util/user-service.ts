import config from './config'
import { fetchApi } from './api';

interface User {
    id: string,
    username: string
}

const login = (username: string, password: string): Promise<User> => {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    };

    return fetch(`${config.apiUrl}/sign_in`, requestOptions)
        .then(handleSignInResponse)
        .then(user => {
            localStorage.setItem('user', JSON.stringify(user));
            return user;
        })
}

const logout = (): void => {
    localStorage.removeItem('user');
    window.location.href = '/sign_in';
}

const signUp = (username: string, password: string): Promise<User> => {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    };

    return fetchApi('/sign_up', requestOptions)
        .then(user => {
            localStorage.setItem('user', JSON.stringify(user));
            return user;
        })
}

const handleSignInResponse = (response: Response): Promise<any> => {
    return response.text().then(text => {
        const data = text && JSON.parse(text);
        if (!response.ok) {
            if (response.status === 401) {
                // logout();

                const error = (data && data.error) || response.statusText;
                return Promise.reject(error);
            }
        }

        return data;
    });
}

export const userService = {
    login,
    logout,
    signUp,
}
