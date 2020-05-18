let apiUrl = '/site/';

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
                window.location.assign('./page.html?keyword=' + that.keyword);
            }
        }
    }
});

let article = new Vue({
    el: '#article',
    data: {
        currPage: 1,
        currArticle: {},
        preArticle: {},
        nextArticle: {}
    },
    methods: {
        goPre: function () {
            let that = this;
            if (!that.preArticle || !that.preArticle.id) {
                return;
            }

            let page = that.currPage + 2;

            axios.get(apiUrl + 'articles/page?size=1&page=' + page)
                .then(function (response) {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                        return;
                    }

                    that.nextArticle = that.currArticle;
                    that.currArticle = that.preArticle;

                    if (!response.data.data) {
                        that.preArticle = null;
                        return;
                    }

                    that.currPage = that.currPage + 1;

                    let data = response.data.data;
                    that.preArticle = data[0];
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        goNext: function () {
            let that = this;
            if (!that.nextArticle || !that.nextArticle.id) {
                return;
            }

            let page = that.currPage - 1;
            if (page < 1) {
                that.preArticle = that.currArticle;
                that.currArticle = that.nextArticle;

                that.nextArticle = null;
                return;
            }

            axios.get(apiUrl + 'articles/page?size=1&page=' + page)
                .then(function (response) {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                        return;
                    }

                    that.preArticle = that.currArticle;
                    that.currArticle = that.nextArticle;

                    if (!response.data.data) {
                        that.nextArticle = null;
                        return;
                    }

                    that.currPage = that.currPage - 1;

                    let data = response.data.data;
                    that.nextArticle = data[0];
                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    },
    computed: {
        pre: function () {
            let that = this;
            if (that.preArticle && that.preArticle.id) {
                return {
                    id: that.preArticle.id,
                    title: that.preArticle.title,
                    success: true
                }
            }
            return {
                title: '没有了',
                success: false
            };
        },
        next: function () {
            let that = this;
            if (that.nextArticle && that.nextArticle.id) {
                return {
                    id: that.nextArticle.id,
                    title: that.nextArticle.title,
                    success: true
                }
            }
            return {
                title: '没有了',
                success: false
            };
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
        hitokoto: '醉后不知天在水 满船清梦压星河'
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

    author.avatar = config.avatar;
    author.author = config.author;
    author.email = config.email;
    author.github = config.github;
    author.website = config.website;

    links.links = config.links;

    beiAn.copyright = '©' + config.since + ' ' + config.author;
}

function getConfig() {
    let config = load('config');
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

            config = save('config', response.data.data);
            setConfig(config);
        })
        .catch(function (error) {
            console.log(error);
        });
}

function initMenus() {
    let menus = load('menus');
    if (menus) {
        header.menus = menus;
        return;
    }

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

            menus = response.data.data.map(menu => {
                let target = menu.url.toLowerCase().startsWith('http') ? '_blank' : '_self';
                return {
                    name: menu.name,
                    url: menu.url,
                    desc: menu.desc,
                    target: target
                }
            });

            header.menus = save('menus', menus);
        })
        .catch(function (error) {
            console.log(error);
        });
}

function initArticle() {
    axios.get(apiUrl + 'articles/page?page=1&size=2')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (!response.data.data) {
                return;
            }

            let data = response.data.data;

            article.currArticle = data[0];

            if (data.length > 1) {
                article.preArticle = data[1];
            }
        })
        .catch(function (error) {
            console.log(error);
        });
}

function initCategory() {
    let categories = load('categories');
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
                category.categories = save('categories', response.data.data);
            }
        })
        .catch(function (error) {
            console.log(error);
        });
}

function initBeiAn() {
    let bei_an = load('bei_an');
    if (bei_an) {
        beiAn.miit = bei_an.miit;
        beiAn.mps = bei_an.mps;
        return;
    }

    axios.get('https://aulang.cn/oauth/api/beian')
        .then(function (response) {
            bei_an = save('bei_an', response.data)
            beiAn.miit = bei_an.miit;
            beiAn.mps = bei_an.mps;
        })
        .catch(function (error) {
            console.log(error);
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
            console.log(error);
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
            console.log(error);
        });
}

function hitokoto() {
    axios.get('https://v1.hitokoto.cn/?encode=text')
        .then(function (response) {
            author.hitokoto = response.data;
        })
        .catch(function (error) {
            console.log(error);
        });
}

hitokoto();
getConfig();
initMenus();
initBeiAn();
initArticle();
initCategory();
initTop3Articles();
initTop3Comments()
