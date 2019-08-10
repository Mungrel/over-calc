import config from './config';
import { userService } from './user-service';

export async function fetchApi(path: string, options: RequestInit): Promise<any> {
    const userItem = localStorage.getItem('user');
    if (!userItem) {
        userService.logout();
        return;
    }

    const user = JSON.parse(userItem);
    if (!user || !user.token) {
        userService.logout();
        return;
    }

    const optionsWithAuth = {
        ...options,
        headers: {
            'Authorization': user.token,
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
    };
    const response = await fetch(url(path), optionsWithAuth);
    return response.text().then(text => {
        const data = text && JSON.parse(text);
        if (!response.ok) {
            if (response.status === 401) {
                userService.logout();
                return;
            }

            const error = (data && data.error) || response.statusText;
            return Promise.reject(error);
        }

        return data;
    })
}

const url = (partialPath: string): string => {
    if (!partialPath.startsWith('/')) {
        partialPath = '/' + partialPath;
    }

    return `${config.apiUrl}${partialPath}`
}
