import {storage} from "../public/storage.js";
import {apiUrl} from "../public/base.js";

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

    axios.get(apiUrl + 'categories')
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
            console.log(error);
        });
}

initCategory();