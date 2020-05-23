import {apiUrl} from './public/base.js';
import {urlParam} from "./public/url.js";
import {storage} from './public/storage.js';

let keyword = urlParam('keyword');
let categoryId = urlParam('category');

if (keyword) {
    alert(keyword);
}

if (categoryId) {
    alert(categoryId);
}
