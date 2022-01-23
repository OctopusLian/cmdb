function HtmlEncode(str) {
    if(typeof(str) === "undefined") {
        return "";
    }
    if(typeof(str) != typeof("")) {
        str = str.toString();
    }
    return str.replace(/&/g, '&amp;')
            .replace(/"/g, '&quot;')
            .replace(/'/g, '&#39;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;');
};

function FileSize(num) {
    units = ["B", "KB", "MB", "GB", "TB", "PB"];
    index = 0;
    while(num > 1024) {
        num /= 1024;
        index += 1
    }
    return num.toFixed(2) + units[index];
}