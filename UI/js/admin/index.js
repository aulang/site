let index = new Vue({
    el: '#index',
    data: {
        isWebsite: true,
        config: {
            title: '',
            desc: '',
            keywords: '',
            author: '',
            website: '',
            email: '',
            github: '',
            wechat: '',
            wechatQRCode: '',
            avatar: '',
            since: '',
            links: [{
                title: '',
                url: '',
                desc: ''
            }]
        },
        menus: [{
            id: '',
            name: '',
            url: '',
            desc: '',
            order: ''
        }],
        articles: [{
            id: '',
            title: '',
            subTitle: '',
            categoryName: '',
            creationDate: '',
            renew: '',
            commentsCount: 0
        }]
    }
});