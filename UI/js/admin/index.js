import {apiUrl} from "../public/base.js";

let index = new Vue({
    el: '#index',
    data: {
        isWebsite: true,
        config: {
            title: '',
            desc: '',
            keywords: '',
            author: '',
            website: '',
            email: '',
            github: '',
            wechat: '',
            wechatQRCode: '',
            avatar: '',
            since: '',
            links: [{
                title: '',
                url: '',
                desc: ''
            }]
        },
        menus: [{
            id: '',
            name: '',
            url: '',
            desc: '',
            order: ''
        }],
        articles: [{
            id: '',
            title: '',
            subTitle: '',
            categoryName: '',
            creationDate: '',
            renew: '',
            commentsCount: 0
        }]
    },
    methods: {
        addMenu: function () {
            alert('添加菜单！');
        },
        addLink: function () {
            alert('添加友情链接！');
        },
        newArticle: function () {
            window.open('./article.html', '_blank');
        },
        editArticle: function (id) {
            window.open(`./article.html?id=${id}`, '_blank');
        }
    }
});

function getConfig() {
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

            index.config = response.data.data;
        })
        .catch(function (error) {
            console.log(error);
        });
}

function getMenus() {
    axios.get(apiUrl + 'menus')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (!response.data.data) {
                return;
            }

            index.menus = response.data.data;
        })
        .catch(function (error) {
            console.log(error);
        });
}

function getArticles(page, size, keyword) {
    let url = apiUrl + `articles/page?page=${page}&size=${size}&keyword=${keyword}`;
    axios.get(url)
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (!response.data.data.datas) {
                return;
            }

            index.articles = response.data.data.datas;
        })
        .catch(function (error) {
            console.log(error);
        });
}

getMenus();
getConfig();
getArticles(1, 10, '');