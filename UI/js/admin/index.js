let index = new Vue({
    el: '#index',
    data: {
        isWebsite: true,
        config: {
            id: '',
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
            menus: [],
            links: []
        },
        articles: [],
        page: 1,
        pageSize: 20,
        keyword: '',
        totalPages: 0
    },
    methods: {
        addMenu: function () {
            let order = this.config.menus.length;
            let data = {
                title: '',
                url: '',
                desc: '',
                order: order,
                edit: true
            };
            this.config.menus.push(data);
        },
        editMenu: function (index, edit) {
            let data = this.config.menus[index];
            data.edit = edit;
            this.config.menus.splice(index, 1, data);
        },
        delMenu: function (index) {
            this.config.menus.splice(index, 1);
        },
        addLink: function () {
            let order = this.config.links.length;
            let data = {
                title: '',
                url: '',
                desc: '',
                order: order,
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
            let config = this.config;
            axios.post('admin/config', config)
                .then(response => {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                    } else {
                        alert('保存成功！')
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
        delArticle: function (id, title) {
            let flag = confirm('确认删除【' + title + '】');
            if (!flag) {
                return;
            }

            let page = this.page;
            let size = this.pageSize;
            let keyword = this.keyword;

            let url = `/admin/article/${id}`;
            axios.delete(url)
                .then(response => {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                    } else {
                        getArticles(page, size, keyword);
                    }
                })
                .catch(error => {
                    console.log(error.data || '删除失败！');
                });
        },
        searchArticle: function () {
            let keyword = this.keyword;
            getArticles(1, 20, keyword);
        },
        goPage: function (page) {
            getArticles(page, this.pageSize, this.keyword);
        }
    },
    computed: {
        noPrevious: function () {
            return (this.page === 1);
        },
        noNext: function () {
            return (this.totalPages === 0 || this.page === this.totalPages);
        }
    },
});

function getConfig() {
    axios.get('config')
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

function getArticles(page, size, keyword) {
    let url = `admin/article/page?page=${page}&size=${size}&keyword=${keyword}`;
    axios.get(url)
        .then(function (response) {
            let result = response.data;
            let code = result.code;
            if (code !== 0) {
                alert(result.msg);
                return;
            }

            if (result.data.datas) {
                index.articles = result.data.datas;
            } else {
                index.articles = [];
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

loginHandle(() => {
    getConfig();
    getArticles(1, 20, '');
});
