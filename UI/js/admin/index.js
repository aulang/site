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
            order: 1
        }],
        articles: [{
            id: '',
            title: '',
            subTitle: '',
            categoryName: '',
            creationDate: '',
            renew: '',
            commentsCount: 0
        }],
        page: 1,
        pageSize: 20,
        keyword: '',
        totalPages: 0
    },
    methods: {
        addMenu: function () {
            let data = {
                id: '',
                name: '',
                url: '',
                desc: '',
                order: '',
                edit: true,
            };
            this.menus.push(data);
        },
        editMenu: function (index, edit) {
            let data = this.menus[index];
            data.edit = edit;
            this.menus.splice(index, 1, data);
        },
        delMenu: function (index) {
            this.menus.splice(index, 1);
        },
        addLink: function () {
            let data = {
                title: '',
                url: '',
                desc: '',
                edit: true
            };
            this.config.links.push(data);
        },
        editLink: function (index, edit) {
            let data = this.config.links[index];
            data.edit = edit;
            this.config.links.splice(index, 1, data);
        },
        delLink: function (index) {
            this.config.links.splice(index, 1);
        },
        cancel: function () {
            window.location.reload();
        },
        saveConfig: function () {
            let menus = this.menus;
            let config = this.config;
            axios.post(apiUrl + 'admin/config', {
                config: config,
                menus: menus
            })
                .then(response => {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                    }
                })
                .catch(error => {
                    console.log(error.data);
                });
        },
        newArticle: function () {
            window.open('./article.html', '_blank');
        },
        editArticle: function (id) {
            window.open(`./article.html?id=${id}`, '_blank');
        },
        delArticle: function (id) {
            let url = apiUrl + `/admin/article/${id}`;
            axios.delete(url)
                .then(response => {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                    }
                })
                .catch(error => {
                    console.log(error.data);
                })
        },
        searchArticle: function () {
            let keyword = this.keyword;
            getArticles(1, 20, keyword);
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
            console.log(error.data);
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
            console.log(error.data);
        });
}

function getArticles(page, size, keyword) {
    let url = apiUrl + `admin/article/page?page=${page}&size=${size}&keyword=${keyword}`;
    axios.get(url)
        .then(function (response) {
            let result = response.data;
            let code = result.code;
            if (code !== 0) {
                alert(result.msg);
                return;
            }

            if (!result.data.datas) {
                return;
            }

            index.articles = result.data.datas;
            index.page = result.data.pageNo;
            index.pageSize = result.data.pageSize;
            index.totalPages = result.data.totalPages;
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

loginHandle();

getMenus();
getConfig();
getArticles(1, 20, '');