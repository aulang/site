let storage = window.sessionStorage;

function save(key, value) {
    let temp = JSON.stringify(value);
    storage.setItem(key, temp);
    return value;
}

function load(key) {
    let value = storage.getItem(key);
    if (value) {
        return JSON.parse(value);
    } else {
        return null;
    }
}

function remove(key) {
    let value = load(key);
    if (value) {
        storage.removeItem(key);
    }
    return value;
}

function clear() {
    storage.clear();
}