import {apiUrl} from './public/base.js';

let article = new Vue({
    el: '#article',
    data: {
        currPage: 1,
        currArticle: {},
        preArticle: {},
        nextArticle: {}
    },
    methods: {
        goArticle: function () {
            let url = './article.html?id=' + this.currArticle.id;
            window.location.assign(url);
        },
        goPre: function () {
            let that = this;
            if (!that.preArticle || !that.preArticle.id) {
                return;
            }

            let page = that.currPage + 2;

            axios.get(apiUrl + 'articles/page?pageSize=1&page=' + page)
                .then(function (response) {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                        return;
                    }

                    that.nextArticle = that.currArticle;
                    that.currArticle = that.preArticle;

                    if (!response.data.data.content) {
                        that.preArticle = null;
                        return;
                    }

                    that.currPage = page;

                    let articles = response.data.data.content;
                    that.preArticle = articles[0];
                })
                .catch(function (error) {
                    console.log(error.data);
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

            axios.get(apiUrl + 'articles/page?pageSize=1&page=' + page)
                .then(function (response) {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                        return;
                    }

                    that.preArticle = that.currArticle;
                    that.currArticle = that.nextArticle;

                    if (!response.data.data.content) {
                        that.nextArticle = null;
                        return;
                    }

                    that.currPage = page;

                    let articles = response.data.data.content;
                    that.nextArticle = articles[0];
                })
                .catch(function (error) {
                    console.log(error.data);
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


function initArticle() {
    axios.get(apiUrl + 'articles/page?pageSize=2&page=1')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (!response.data.data.content) {
                return;
            }

            let articles = response.data.data.content;

            article.currArticle = articles[0];

            if (articles.length > 1) {
                article.preArticle = articles[1];
            }
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

initArticle();
