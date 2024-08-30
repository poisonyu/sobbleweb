
var vditor = null;
window.onload = function() {
    vditor = new Vditor(document.getElementById('vditor'), {
        placeholder: "placeholder", // 默认内容
        height: window.innerHeight - 40,
        lang: "zh_CN",
        // value: "{{ .article.Content }}",
        cache: {
            enable: false
        },
        "mode": "sv",
        "preview": {
            "mode": "both"
        }
    });

}
// 测试数据填充
function setContent() {
    vditor.setValue("## 测试 \n ### 二级标题 ");
}

function saveContent() {
    var htmlContent = vditor.getHTML();

    let title = document.getElementById("title").value 
    if (!title) {
        var re = new RegExp("<h1>(.*?)<");
        // todo 没有匹配到结果的情况
        let result = htmlContent.match(re)
        if (!result) {
            alert("请填写标题或在markdown中填写一级标题")
            return 
        }
        title = result[1]

    }
    let tag = document.getElementById("tag").value 
    fetch("/article/add", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "nickname":"",
            "title": title,
            "type":tag,
            "mdcontent": vditor.getValue(),
            "htmlcontent": htmlContent,
            // "ishtml": true,
        })
    })
    .then(response => response.json())
    .then(response => callback(response))
    // .then(response => response.json())
    // .then(data => console.log(data.message))
}

function callback(response) {
    console.log(response)
    if (response.code == 1) {
        alert("发布成功");
        window.location.href = response.data.redirect
    } else {
        alert("发生错误")
    }
}
