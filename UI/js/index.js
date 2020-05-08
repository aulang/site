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
            alert(that.keyword);
        }
    }
});

let article = new Vue({
    el: '#article',
    data: {
        id: 'articleId',
        title: "title",
        subTitle: 'subTitle',
        summary: '<h1 class="display-4">Hello, world!</h1>',
        commentsCount: 4,
        comments: [
            {
                name: "user1",
                content: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                creationDate: "2020-05-04 22:12:32",
                replies: null
            },
            {
                name: "user2",
                content: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
                creationDate: "2020-05-04 22:12:32",
                replies: [
                    {
                        name: "user3",
                        content: "cccccccccccccccccccccccccccccccccccccc",
                        creationDate: "2020-05-04 23:01:12",
                    },
                    {
                        name: "user4",
                        content: "dddddddddddddddddddddddddddddddddddddd",
                        creationDate: "2020-05-04 23:03:15",
                    }
                ]
            }
        ],
        preArticle: null,
        nextArticle: {
            id: 'articleId',
            title: "测试文章3"
        }
    },
    methods: {
        goPre: function () {
            let that = this;
            if (!that.preArticle || !that.preArticle.id) {
                return
            }

            alert(that.preArticle.id);
        },
        goNext: function () {
            let that = this;
            if (!that.nextArticle || !that.nextArticle.id) {
                return
            }

            alert(that.nextArticle.id);
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
        avatar: 'https://aulang.cn/oauth/images/nologo400.png',
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

            let data = response.data.data;
            header.title = data.title;

            author.avatar = data.avatar;
            author.author = data.author;
            author.email = data.email;
            author.github = data.github;
            author.website = data.website;

            links.links = data.links;

            beiAn.copyright = '©' + data.since + ' ' + data.author;
        })
        .catch(function (error) {
            console.log(error);
        });
}

function initMenus() {
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

            header.menus = response.data.data.map(menu => {
                let target = menu.url.toLowerCase().startsWith('http') ? '_blank' : '_self';
                return {
                    name: menu.name,
                    url: menu.url,
                    desc: menu.desc,
                    target: target
                }
            });
        })
        .catch(function (error) {
            console.log(error);
        });
}

function initCategory() {
    axios.get(apiUrl + 'categories')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (response.data.data) {
                category.categories = response.data.data;
            }
        })
        .catch(function (error) {
            console.log(error);
        });
}

function initBeiAn() {
    axios.get('https://aulang.cn/oauth/api/beian')
        .then(function (response) {
            beiAn.miit = response.data.miit;
            beiAn.mps = response.data.mps;
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
initCategory();
initTop3Articles();
initTop3Comments()
