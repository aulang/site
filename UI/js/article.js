import {apiUrl} from './public/base.js';
import {urlParam} from "./public/url.js";
import {storage} from './public/storage.js';

let articleId = urlParam('id');

if (articleId) {
    alert(articleId)
}