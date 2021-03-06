import {apiUrl} from './public/base.js';
import {urlParam} from './public/url.js'
import {storage} from './public/storage.js';

let header = new Vue({
    el: '#header',
    data: {
        title: 'Aulang',
        menus: [],
        keyword: ''
    },
    methods: {
        search: function () {
            let that = this;
            if (that.keyword) {
                let url = './page.html?keyword=' + that.keyword;

                let categoryId = urlParam('category');
                if (categoryId) {
                    url = url + '&category=' + categoryId;
                }

                window.location.assign(url);
            }
        }
    }
});

let author = new Vue({
    el: '#author',
    data: {
        avatar: './images/aulang.jpg',
        author: 'Aulang',
        website: 'https://aulang.cn',
        email: 'aulang@qq.com',
        github: 'https://github.com/aulang',
        wechat: 'aulang88',
        wechatQRCode: './images/wechat.png',
        hitokoto: '醉后不知天在水 满船清梦压星河',
        avatarTmp: null
    },
    methods: {
        showQRCode: function () {
            this.avatarTmp = this.avatar;
            this.avatar = this.wechatQRCode;
        },
        hideQRCode: function () {
            if (this.avatarTmp) {
                this.avatar = this.avatarTmp;
            }
        }
    }
});

let top3Comments = new Vue({
    el: '#top3Comments',
    data: {
        comments: []
    }
});

let top3Articles = new Vue({
    el: '#top3Articles',
    data: {
        articles: []
    }
});

let category = new Vue({
    el: '#category',
    data: {
        categories: []
    }
});

let links = new Vue({
    el: '#links',
    data: {
        links: []
    }
});

let beiAn = new Vue({
    el: '#beiAn',
    data: {
        copyright: '©2018 Aulang',
        miit: {
            no: '鄂ICP备18028762号',
            url: 'http://beian.miit.gov.cn'
        },
        mps: {
            no: '鄂公网安备42011102003833号',
            url: 'http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=42011102003833'
        }
    }
});

function setConfig(config) {
    header.title = config.title;
    header.menus = config.menus;

    author.avatar = config.avatar;
    author.author = config.author;
    author.email = config.email;
    author.github = config.github;
    author.website = config.website;

    links.links = config.links;

    beiAn.copyright = '©' + config.since + ' ' + config.author;
}

function getConfig() {
    let config = storage.load('config');
    if (config) {
        setConfig(config);
        return;
    }

    axios.get(apiUrl + 'config')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (!response.data.data) {
                return;
            }

            config = response.data.data;
            let menus = config.menus;
            if (menus) {
                config.menus = menus.map(e => {
                    if (e.url.startsWith('http')) {
                        e.target = '_blank';
                    } else {
                        e.target = '_self';
                    }
                    return e;
                });
            }

            config = storage.save('config', config);
            setConfig(config);
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

function initCategory() {
    let categories = storage.load('categories');
    if (categories) {
        category.categories = categories;
        return;
    }

    axios.get(apiUrl + 'categories')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (response.data.data) {
                category.categories = storage.save('categories', response.data.data);
            }
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

function initBeiAn() {
    let bei_an = storage.load('bei_an');
    if (bei_an) {
        beiAn.miit = bei_an.miit;
        beiAn.mps = bei_an.mps;
        return;
    }

    axios.get('https://aulang.cn/oauth/api/beian')
        .then(function (response) {
            bei_an = storage.save('bei_an', response.data)
            beiAn.miit = bei_an.miit;
            beiAn.mps = bei_an.mps;
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

function initTop3Articles() {
    axios.get(apiUrl + 'articles/top3')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (response.data.data) {
                top3Articles.articles = response.data.data;
            }
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

function initTop3Comments() {
    axios.get(apiUrl + 'comment/top3')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (response.data.data) {
                top3Comments.comments = response.data.data;
            }
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

function hitokoto() {
    axios.get('https://v1.hitokoto.cn/?encode=text')
        .then(function (response) {
            author.hitokoto = response.data;
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

hitokoto();
getConfig();
initBeiAn();
initCategory();
initTop3Articles();
initTop3Comments();