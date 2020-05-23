function urlParam(name) {
    let search = window.location.search;
    if (!search) {
        return null;
    }

    let query = search.substring(1);
    let params = query.split("&");

    for (let i = 0; i < params.length; i++) {
        let pair = params[i].split("=");
        if (pair[0] == name) {
            return pair[1];
        }
    }

    return null;
}

export {urlParam}