let infoform = document.getElementById("userinfo");
let changepasswordform = document.getElementById("changepassword");
let verificationbtn = document.getElementById("verification");

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
    }
    alert(data.message)
}