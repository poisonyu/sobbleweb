window.onload = function() {
    let content = document.getElementById('content');
    let text = content.textContent;
    content.innerHTML = text;
    content.style.display = "";
}
function editArticle() {
    window.location = "http://localhost:8888/editor"
}   

function deleteArticle() {
    fetch("/article/delete", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        "id": "{{ .article.ID }}",
      })
    })
    .then(response => callback(response))

}
function callback(response) {
    console.log(response)
    if (response.redirected) {
      alert("删除成功")
      window.location.href = response.url
    }

}
