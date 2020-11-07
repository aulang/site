import {apiUrl} from './public/base.js';
import {urlParam} from './public/url.js';

let articleId = urlParam('id');

if (!articleId) {
    window.location.assign('./index.html');
}

let article = new Vue({
    el: '#article',
    data: {
        article: {},
        comment: {
            'name': '',
            'mail': '',
            'content': '',
            'articleId': articleId
        }
    },
    methods: {
        onSubmit: function (e) {
            let formData = JSON.stringify(this.comment);

            e.currentTarget.disabled = true;
            axios.post(apiUrl + 'comment', formData)
                .then(res => {
                    let code = res.data.code;
                    if (code !== 0) {
                        alert(res.data.msg);
                        return;
                    }

                    article.name = '';
                    article.mail = '';
                    article.content = '';

                    article.article.comments.push(res.data.data);

                    article.article.commentsCount += 1;
                })
                .catch(err => {
                    console.log(err);
                })
                .then(() => {
                    e.currentTarget.disabled = false;
                });
        }
    }
});

function getArticle(id) {
    axios.get(apiUrl + 'articles/' + id)
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            article.article = response.data.data;
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

getArticle(articleId);