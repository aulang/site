import {apiUrl} from './public/base.js';
import {urlParam} from "./public/url.js";

let articleId = urlParam('id');

if (!articleId) {
    window.location.assign('./index.html');
}

let article = new Vue({
    el: '#article',
    data: {
        article: {
        }
    },
    methods: {

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
        // article.comments
    })
    .catch(function (error) {
        console.log(error);
    });
}

getArticle(articleId);