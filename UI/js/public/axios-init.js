let baseUrl = 'http://office.aulang.cn';
let clientId = '5f37d9f4c4155cda795f8fe5';
let redirectUri = encodeURI('http://office.aulang.cn/index.html');

// 配置请求baseURL
axios.defaults.baseURL = baseUrl;
// 允许跨域
axios.defaults.crossDomain = true;
// 允许跨域携带Cookie
axios.defaults.withCredentials = false;
// 设置请求头为 Authorization
axios.defaults.headers.common['Authorization'] = '';


// 设置Authorization请求头Token
axios.interceptors.request.use(
    config => {
        let token = ttlLocalStorage.getItem('access_token');
        token && (config.headers.Authorization = 'Bearer ' + token);
        return config;
    },
    error => {
        return Promise.error(error);
    }
);


axios.interceptors.response.use(
    response => {
        // 如果返回的状态码为200，说明接口请求成功，可以正常拿到数据，否则的话抛出错误
        if (response.status === 200) {
            return Promise.resolve(response);
        } else {
            return Promise.reject(response);
        }
    },
    // 服务器状态码不是2开头的的情况
    error => {
        if (error.response.status) {
            switch (error.response.status) {
                // 401: 未登录，跳转登录页面
                case 401:
                    // 清除token
                    ttlLocalStorage.removeItem('access_token');
                    // 重定向登录页面
                    redirectLoginUrl();
                    break;
                default:
                    return Promise.reject(error.response);
            }
            return Promise.reject(error.response);
        }
    }
);

let scope = 'basic';
let responseType = 'token';

function redirectLoginUrl() {
    let state = random(6);
    ttlLocalStorage.setItem('oauth_state', state, 600);

    let url = `https://aulang.cn/oauth/authorize?client_id=${clientId}&response_type=${responseType}&scope=${scope}&redirect_uri=${redirectUri}&state=${state}`;
    window.location.replace(url);
}

function urlParam(name) {
    let search = window.location.search;
    if (!search) {
        return null;
    }

    let query = search.substring(1);
    let params = query.split('&');

    for (let i = 0; i < params.length; i++) {
        let pair = params[i].split('=');
        if (pair[0] === name) {
            return pair[1];
        }
    }

    return null;
}

function random(length) {
    let str = Math.random().toString(36).substr(2);
    if (str.length >= length) {
        return str.substr(0, length);
    }
    str += random(length - str.length);
    return str;
}

function loginHandle(loginFunc) {
    let token = ttlLocalStorage.getItem('access_token');
    if (token) {
        // 已登录
        loginFunc();
        return;
    }

    let state = urlParam('state');
    let expires_in = urlParam('expires_in');
    let access_token = urlParam('access_token');

    if (!state || !access_token || !expires_in) {
        // 未登录没有state、access_token和expires_in
        redirectLoginUrl();
        return;
    }

    let currState = ttlLocalStorage.getItem('oauth_state');
    if (state !== currState) {
        // 不合法的state
        alert('不合法的state！');
        redirectLoginUrl();
        return;
    }

    // 清空state
    ttlLocalStorage.removeItem('oauth_state');
    // 登录成功
    ttlLocalStorage.setItem('access_token', access_token, parseInt(expires_in) - 10);
    // 清空URL参数
    window.history.replaceState({}, '', redirectUri);
    // 登录成功执行操作
    loginFunc();
}

// 带TTL的localStorage
const ttlLocalStorage = {
    setItem: function (key, value, ttl) {
        let expires = null;
        if (!Number.isNaN(ttl)) {
            expires = new Date().getTime() + ttl * 1000;
        }

        let obj = {
            key: key,
            value: value,
            expires: expires
        }

        window.localStorage.setItem(key, JSON.stringify(obj));
    },
    getItem: function (key) {
        let item = window.localStorage.getItem(key);

        if (!item) {
            return null;
        }

        let obj = JSON.parse(item);

        if (obj.expires) {
            let date = new Date().getTime();
            if (date > obj.expires) {
                window.localStorage.removeItem(name);
                return null;
            }
        }

        return obj.value;
    },
    removeItem: function (key) {
        window.localStorage.removeItem(key);
    },
    clear: function () {
        window.localStorage.clear();
    }
}