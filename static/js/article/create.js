
let vditor = null;
let savebtn = document.getElementById("save");

window.onload = function() {
    vditor = new Vditor(document.getElementById('vditor'), {
        placeholder: "请使用Markdown编辑", // 默认内容
        height: window.innerHeight - 40,
        lang: "zh_CN",
        // value: "{{ .article.Content }}",
        cache: {
            enable: false
        },
        mode: "sv",
        preview: {
            "mode": "both"
        },
        // upload: {
        //     accept: "image/jpg, image/jpeg, image/png",
        //     url: "", 
        //     multiple: false,

        // },
    });

}

savebtn.addEventListener("click", function(e) {
    e.preventDefault();
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
    let formData = new FormData();
    let formFile = document.getElementById("formFile")
    let file = null
    if (formFile.files[0]) {
        file = formFile.files[0]
        formData.append("cover", file, file.name)
    }
    formData.append("title", title)
    formData.append("tag", tag)
    formData.append("mdContent", vditor.getValue())
    formData.append("htmlContent", htmlContent)
    fetch("/article/create", {
        method: 'POST',
        body: formData,
    })
    .then(response => response.json())
    .then(response => callback(response))
})


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

// 测试数据填充
// function setContent() {
//     vditor.setValue("## 测试 \n ### 二级标题 ");
// }