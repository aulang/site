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
        },
        saveArticle: function () {
            let thiz = this;
            let content = marked(thiz.source);
            let categoryName = getCategoryName();
            let postData = {
                id: thiz.id,
                title: thiz.title,
                subTitle: thiz.subTitle,
                categoryId: thiz.categoryId,
                categoryName: categoryName,
                summary: thiz.summary,
                source: thiz.source,
                content: content
            };
            axios.post('admin/article', postData)
                .then(response => {
                    if (response.data.code !== 0) {
                        alert(response.data.msg);
                        return;
                    }
                    alert('保存成功！');
                    thiz.id = response.data.data.id;
                }).catch(error => {
                console.log(error.data || '保存失败！');
            });
        },
        cancel: function () {
            if (this.id) {
                getArticle(this.id);
            } else {
                window.location.reload();
            }
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
            article.id = id;
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

function getCategoryName() {
    let category = document.getElementById("category");
    let index = category.selectedIndex;
    return category.options[index].text;
}

loginHandle(initCategory);

let id = urlParam('id');
if (id) {
    getArticle(id);
}