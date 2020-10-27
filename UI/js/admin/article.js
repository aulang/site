let article = new Vue({
    el: '#article',
    data: {
        title: '',
        subTitle: '',
        categoryId: '',
        summary: '',
        content: ''
    },
    methods: {
        upload: function () {
            alert('上传资源');
        }
    }
});