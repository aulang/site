let storage = {
    storage: window.sessionStorage,
    save: function (key, value) {
        let temp = JSON.stringify(value);
        this.storage.setItem(key, temp);
        return value;
    },
    load: function (key) {
        let value = this.storage.getItem(key);
        if (value) {
            return JSON.parse(value);
        } else {
            return null;
        }
    },
    remove: function (key) {
        let value = this.load(key);
        if (value) {
            this.storage.removeItem(key);
        }
        return value;
    },
    clear: function clear() {
        this.storage.clear();
    }
}

export {storage};