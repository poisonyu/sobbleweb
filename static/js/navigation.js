
let logout = document.getElementById("logout")

if (logout) {
    logout.addEventListener("click", function (e) {
        e.preventDefault();
        let token = getCookie("jwt-token");
        let d = new Date();
        d.setTime(d.getTime() - 1);
        let expireTime = d.toGMTString()
        document.cookie = "jwt-token=" + token + ";expires=" + expireTime + ";path=/"
        window.location.href = "/"
    })
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