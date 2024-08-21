
var vditor = null;
window.onload = function() {
    vditor = new Vditor(document.getElementById('vditor'), {
        // placeholder: "placeholder",
        height: window.innerHeight - 40,
        lang: "zh_CN",
        value: "{{ .article.MdContent }}",
        cache: {
            enable: false
        },
        "mode": "sv",
        "preview": {
            "mode": "both"
        }
    });

}
function saveContent() {
    var htmlContent = vditor.getHTML();
    var re = new RegExp("<h1>(.*?)<");
    // todo 没有匹配到结果的情况
    var title = htmlContent.match(re)[1];
    fetch("/article/update", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "id": "{{ .article.ID }}",
            "nickname":"",
            "title": title,
            "type":"",
            "mdcontent": vditor.getValue(),
            "htmlcontent": htmlContent,
            // "ishtml": true,
        })
    })
    .then(response => callback(response))
    // .then(response => response.json())
    // .then(data => console.log(data.message))
    vditor.destory();
}
function callback(response) {
    console.log(response)
    if (response.redirected) {
        alert("保存成功");
        window.location.href = response.url
    }
}
