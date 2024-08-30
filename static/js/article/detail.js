window.onload = function() {
    let content = document.getElementById('content');
    let text = content.textContent;
    content.innerHTML = text;
    content.style.display = "";
}

let deleteButton = document.getElementById("delete")

deleteButton.addEventListener("click", function(e) {
    console.log(e);
    console.log(e.target.attributes);
    fetch("/article/delete", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "id": e.target.getAttribute("articleid")
        })
    })
    .then(response => response.json())
    .then(jsonresponse => callback(jsonresponse))
})

function callback(jsonresponse) {
    console.log(jsonresponse)
    if (jsonresponse.code == 1) {
      alert("删除成功")
      window.location.href = jsonresponse.data.redirect
    } else {
        alert("删除失败\n" + jsonresponse.message)
    }

}

// function editArticle() {
//     window.location = "http://localhost:8888/editor"
// }   

// function deleteArticle() {
//     console.log(this);
//     fetch("/article/delete", {
//         method: "POST",
//         headers: {
//             "Content-Type": "application/json",
//         },
//         body: JSON.stringify({
//             // GO HTML模板渲染 貌似不能放在JavaScript文件中
//             "id": this.getAttribute("articleid")
//         })
//     })
//     .then(response => response.json())
//     .then(jsonresponse => callback(jsonresponse))

// }

// function getCookie(name) {
//     let cookieArr = document.cookie.split(";");

//     for (let i = 0; i < cookieArr.length; i++) {
//         let cookiePair = cookieArr[i].split("=");
//         if (name == cookiePair[0].trim()) {
//             return cookiePair[1];
//         }
//     }
//     return null;
// }

// function logout() {
//     let token = getCookie("jwt-token");
//     let d = new Date();
//     d.setTime(d.getTime() - 1);
//     let expireTime = d.toGMTString()
//     document.cookie = "jwt-token=" + token + ";expires=" + expireTime + ";path=/"
// }