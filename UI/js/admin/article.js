import {storage} from "../public/storage.js";
import {urlParam} from "../public/url.js";

let article = new Vue({
    el: '#article',
    data: {
        title: '',
        subTitle: '',
        categoryId: '',
        summary: '',
        source: '',

        resources: [],
        categories: [],
        isEdit: true,
        content: ''
    },
    methods: {
        upload: function () {
            alert('上传资源');
        },
        view: function () {
            // markdown转html
            this.content = marked(this.source);
        }
    }
});

function initCategory() {
    let categories = storage.load('categories');
    if (categories) {
        article.categories = categories;
        return;
    }

    axios.get('categories')
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (response.data.data) {
                article.categories = storage.save('categories', response.data.data);
            }
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

function getArticle(id) {
    axios.get('articles/' + id)
        .then(function (response) {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            article.title = response.data.data.title;
            article.subTitle = response.data.data.subTitle;
            article.categoryId = response.data.data.categoryId;
            article.summary = response.data.data.summary;
            article.source = response.data.data.source;

            if (!response.data.data.source && response.data.data.content) {
                article.source = response.data.data.content;
            }
        })
        .catch(function (error) {
            console.log(error.data);
        });
}

loginHandle();
initCategory();

let id = urlParam('id');
if (id) {
    getArticle(id);
}