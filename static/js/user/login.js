

function validateThisField (field) {
	if (field.required && field.value === '') {
		field.classList.add('required');
		formIsValid = false;
	}
	if (field.pattern
			&& !(new RegExp(field.pattern).exec(field.value) !== null) ) {
		field.classList.add('invalid');
        //console.log(field)
		formIsValid = false;
	}

}

let captchaid = null
var form = document.getElementById('loginform');
var fields = form.querySelectorAll('input');
var formIsValid = true;

form.addEventListener('submit', function (e) {
	e.preventDefault();
	Array.prototype.forEach.call(fields, validateThisField);
  	// also have a global state on the form
	if (!formIsValid) {
    form.classList.remove('errors');
        setTimeout(function() {form.classList.add('errors');}, 0);
    }
    let username = document.getElementById("username").value;
    let password = document.getElementById("password").value;
    let captchav = document.getElementById("captcha").value
    if (!username || !password || !captchav) {
        return 
    }
    fetch("/user/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "username": username,
            "password": password,
            "captcha": captchav,
            "captchaid": captchaid,
        })
    })
    .then(response => response.json())
    .then(data => callback(data))
});

function callback(response) {
    console.log(response)
    if (response.code == 1) {
        // let referer = form.getAttribute("redirect")
        // if (referer != "") {
        //     window.location.href = referer
        // } else {
        //     window.location.href = "/"
        // }
        let referer = document.referer
        console.log("referer", referer)
        if (referer) {
            console.log(referer)
            window.location.href = referer
        } else {
            window.location.href = "/"
        }
        // window.location.href = document.referrer
    } else {
        alert("登录失败\n" + response.message);
        generate_captcha()
    }
}
/*
 form.addEventListener('focus', function (e) {
	e.target.classList.remove('required');
	e.target.classList.remove('invalid');
}, true); */

form.addEventListener('blur', function (e) {
  e.target.classList.remove('required');
	e.target.classList.remove('invalid');
	validateThisField(e.target);
}, true);


function generate_captcha() {
    fetch("/captcha")
    .then(response => response.json())
    .then(data => captcha(data))
}
function captcha(data) {
    console.log(data)
    if (data.code == 1) {
        let img = document.querySelector(".captcha-img")
        img.src = data.data.b64s 
        captchaid = data.data.id 
    }
}
generate_captcha()