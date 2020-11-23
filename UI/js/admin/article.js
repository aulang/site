import {storage} from "../public/storage.js";

marked.setOptions({
    highlight: function (code) {
        return hljs.highlightAuto(code).value;
    }
});

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
        upload: function (e) {
            let thiz = this;
            let id = thiz.id;
            if (!id) {
                alert('请先保存文章，然后再上传资源!');
                return;
            }

            let config = {
                headers: {
                    'Content-Type': 'multipart/form-data;boundary = ' + new Date().getTime()
                }
            }
            let formData = new FormData();
            formData.append('file', e.target.files[0]);

            axios.post('admin/resource/subject/' + id, formData, config)
                .then(response => {
                    this.resources.push(toResource(response.data.data))
                })
                .catch(error => {
                    alert(error.data || '上传资源失败');
                });
        },
        delResource: function (index, id) {
            let thiz = this;
            axios.delete('admin/resource/' + id)
                .then(response => {
                    let code = response.data.code;
                    if (code !== 0) {
                        alert(response.data.msg);
                    } else {
                        thiz.resources.splice(index, 1);
                    }
                })
                .catch(error => {
                    console.log(error.data || '删除失败！');
                });
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

function getResource(id) {
    axios.get('admin/resource/subject/' + id)
        .then(response => {
            let code = response.data.code;
            if (code !== 0) {
                alert(response.data.msg);
                return;
            }

            if (!response.data.data) {
                return;
            }

            article.resources = response.data.data.map(e => toResource(e));
        }).catch(error => {
        alert(error.data || '获取资源失败!');
    });
}

function getCategoryName() {
    let category = document.getElementById("category");
    let index = category.selectedIndex;
    return category.options[index].text;
}

function toResource(result) {
    let id = result.id
    let filename = result.filename;

    let url = `${baseUrl}resource/${id}`;
    return {
        id: id,
        filename: filename,
        url: url
    }
}

loginHandle(initCategory);

let id = urlParam('id');
if (id) {
    getArticle(id);
    getResource(id);
}