(function (d) {
    const blank_suf = "__blank/";
    var base_url = document.baseURI;
    base_url = base_url.substring(0, base_url.indexOf("algorithm-in-action"));
    const mains = d.getElementsByTagName("main");
    if (mains.length == 1) {
        const as = mains.item(0).getElementsByTagName("a");
        for (const element of as) {
            const a = element;
            const url = a.href;
            if (!url.startsWith(base_url)) {
                a.target = "_blank"
            } else if (url.endsWith(blank_suf)) {
                a.href = url.substring(0, url.length - blank_suf.length)
                a.target = "_blank"
            }
        }
    }
})(document);
