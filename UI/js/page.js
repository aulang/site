import {apiUrl} from './public/base.js';
import {urlParam} from './public/url.js'
import {storage} from './public/storage.js';

let keyword = urlParam('keyword');
let categoryId = urlParam('category');

if (keyword) {
    /** 关键字搜索 **/
}

if (categoryId) {
    /** 类别查找 **/
}

var articles = new Vue({
    el: '#articles',
    data: {
        page: 0,
        size: 20,
        articles: [],
        totalPages: 0
    },
    computed: {
        noPrevious: function () {
            return (this.totalPages <= 1 || this.page === 0);
        },
        noNext: function () {
            return (this.totalPages <= 1 || this.page === this.totalPages);
        }
    },
    methods: {
        goPage: function (page) {
            getArticles(page, this.size);
        }
    }
});

function getArticles(page, size) {
    page = page || 0;
    size = size || 20;

    let url = 'articles/page?page=' + page + '&size=' + size;
    if(keyword) {
        url = url +  '&keyword=' + keyword;
    }
    if (categoryId) {
        url = url +  '&category=' + categoryId;
    }

    axios.get(apiUrl + url)
            .then(function (response) {
                let code = response.data.code;
                if (code !== 0) {
                    alert(response.data.msg);
                    return;
                }

                let page = response.data.data;

                articles.page = page.pageNo;
                articles.size = page.pageSize;
                articles.articles = page.datas;
                articles.totalPages = page.totalPages;
            })
            .catch(function (error) {
                console.log(error);
            });
}

getArticles(0, 20);