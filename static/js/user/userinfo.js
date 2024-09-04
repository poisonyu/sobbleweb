let infoform = document.getElementById("userinfo");

infoform.addEventListener("submit", function(e) {
    e.preventDefault();

    let id = infoform.getAttribute("userid")
    
    let nickname = document.getElementById("Inputnickname").value 
    let email = document.getElementById("InputEmail").value 
    let phone = document.getElementById("Inputphonenumber").value
    
    fetch("/user/editinfo", {
        method: "POST", 
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "id": parseInt(id),
            "nickname": nickname,
            "email": email,
            "phone": phone,
        })
    })
    .then(response => response.json())
    .then(response => callback(response))
})

function callback(response) {
    console.log(response)
    if (response.code == 1) {
        window.location.reload()
    } else {
        alert("保存失败")
    }
}