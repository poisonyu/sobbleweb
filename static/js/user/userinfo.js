let infoform = document.getElementById("userinfo");
let changepasswordform = document.getElementById("changepassword");
let verificationbtn = document.getElementById("verification");
// let captchaid = null

infoform.addEventListener("submit", function(e) {
    e.preventDefault();

    // let id = infoform.getAttribute("userid")
    
    let nickname = document.getElementById("Inputnickname").value 
    let email = document.getElementById("InputEmail").value 
    let phone = document.getElementById("Inputphonenumber").value
    
    fetch("/user/editinfo", {
        method: "POST", 
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            // "id": parseInt(id),
            "nickname": nickname,
            "email": email,
            "phone": phone,
        })
    })
    .then(response => response.json())
    .then(response => callback(response))
})

// changepasswordform.addEventListener("submit", function(e) {
//     e.preventDefault();
//     let id = e.target.getAttribute("userid");
//     fetch("/user/changepassword", {
//         method: "POST", 
//         headers: {
//             "Content-Type": "application/json",
//         },
//         body: JSON.stringify({
//             "id": parseInt(id),
//         })
//     })
//     .then(response => response.json())
//     // .then(response => )
// })


function callback(response) {
    console.log(response)
    if (response.code == 1) {
        window.location.reload()
    } else {
        alert("保存失败")
    }
}

verificationbtn.addEventListener("click", function(e){
    e.preventDefault()
    // let id = infoform.getAttribute("userid")
    fetch("/user/verification")
    .then(response => response.json())
    .then(response => verificationCallback(response))
})

function verificationCallback(data) {
    if (data.code == 1) {
        verificationbtn.classList.add("disabled")
        // captchaid = data.data.id
    }
    alert(data.message)
}

changepasswordform.addEventListener("submit", function(e) {

    e.preventDefault();
    let captcha = document.getElementById("emailvalidation").value
    let password = document.getElementById("InputPassword").value
    let newpassword = getValueById("InputNewPassword")
    if (captcha && password && newpassword) {
        fetch("/user/changepassword", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                "password": password,
                "newpassword": newpassword, 
                // "captchaid": captchaid,
                "verificationcode": captcha,
            })
        })
        .then(response => response.json())
        .then(response => changePasswordCallback(response))
        
    } else {
        console.log(captcha, captchaid, password, newpassword)
    }

})

function changePasswordCallback(response) {
    if (response.code == 1) {
        deletecookie()
    } else {
        alert(response.message)
    }

}
function getValueById(elementid) {
    return document.getElementById(elementid).value
}

function deletecookie() {
    let token = getCookie("jwt-token");
    let d = new Date();
    d.setTime(d.getTime() - 1);
    let expireTime = d.toGMTString()
    document.cookie = "jwt-token=" + token + ";expires=" + expireTime + ";path=/"
    window.location.href = "/user/signin"
}

function getCookie(name) {
    let cookieArr = document.cookie.split(";");

    for (let i = 0; i < cookieArr.length; i++) {
        let cookiePair = cookieArr[i].split("=");
        if (name == cookiePair[0].trim()) {
            return cookiePair[1];
        }
    }
    return null;
}