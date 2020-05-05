let header = new Vue({
    el: '#header',
    data: {
        title: 'Aulang',
        menus: [
            {
                name: '百度',
                url: 'https://www.baidu.com',
                desc: '百度一下',
                target: '_blank'
            },
            {
                name: '吴浪',
                url: 'https://aulang.cn',
                desc: 'Aulang',
                target: '_self'
            }
        ],
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

function hitokoto() {
    axios.get('https://v1.hitokoto.cn/?encode=text')
        .then(function (response) {
            author.hitokoto = response.data;
        })
        .catch(function (error) {
            console.log(error);
        });
}

let recentReplies = new Vue({
    el: '#recentReplies',
    data: {
        replies: [
            {
                name: "user2",
                content: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
                creationDate: "2020-05-04 23:01:12",
            },
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
});

let category = new Vue({
    el: '#category',
    data: {
        categories: [
            {
                id: '1',
                name: '经',
                count: 1
            },
            {
                id: '2',
                name: '史',
                count: 2
            },
            {
                id: '3',
                name: '子',
                count: 3
            },
            {
                id: '4',
                name: '集',
                count: 4
            }
        ]
    }
});

let links = new Vue({
    el: '#links',
    data: {
        links: [
            {
                title: '百度',
                url: 'https://www.baidu.com',
                desc: '百度一下，你就知道'
            },
            {
                title: 'IT之家',
                url: 'https://www.ithome.com',
                desc: 'IT人的生活，尽在IT之家，爱IT，爱这里'
            },
            {
                title: '管理后台',
                url: 'https://aulang.cn',
                desc: '不要瞎点了，不会让你知道的'
            },
        ]
    }
});

let beiAn = new Vue({
    el: '#beiAn',
    data: {
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

function bei_an() {
    axios.get('https://aulang.cn/oauth/api/beian')
        .then(function (response) {
            beiAn.miit = response.data.miit;
            beiAn.mps = response.data.mps;
        })
        .catch(function (error) {
            console.log(error);
        });
}

hitokoto();
bei_an()