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

let captchaid = null;
var form = document.getElementById('loginform');
var fields = form.querySelectorAll('input');
var formIsValid = true;
let password = null;
let passwordRepeat = document.getElementById("passwordRepeat")


form.addEventListener('submit', function (e) {
	e.preventDefault();
	Array.prototype.forEach.call(fields, validateThisField);
  	// also have a global state on the form
	if (!formIsValid) {
    form.classList.remove('errors');
        setTimeout(function() {form.classList.add('errors');}, 0);
        return
    }

    // let password = document.getElementById("password").value
    // let passwordRepeat = document.getElementById("passwordRepeat").value

    fetch("/user/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "username": document.getElementById("username").value,
            "password": password.value,
            "passwordrepeat": passwordRepeat.value,
            "nickname": document.getElementById("nickname").value,
            "captcha": document.getElementById("captcha").value,
            "captchaid": captchaid,
        })
    })
    .then(response => response.json())
    .then(data => callback(data))
});

function callback(response) {
    //console.log(response)
    if (response.message == "success") {
        alert("注册成功")
        window.location.href = "/"
    } else {
        alert(response.message);
        generate_captcha()
    }
}

form.addEventListener('blur', function (e) {
    if (e.target.id == "passwordRepeat") {
        return 
    }
    e.target.classList.remove('required');
	e.target.classList.remove('invalid');
	validateThisField(e.target);
}, true);

passwordRepeat.addEventListener('blur', function (e) {
    formIsValid = true
    e.target.classList.remove('required')
    password = document.getElementById("password");
    //console.log("password", password.value)
    if (password.value != passwordRepeat.value) {
        e.target.classList.add('required')
        formIsValid = false
    }
})

function generate_captcha() {
    fetch("/captcha")
    .then(response => response.json())
    .then(data => captcha(data))
}
function captcha(data) {
    //console.log(data)
    if (data.message == "success") {
        let img = document.querySelector(".captcha-img")
        img.src = data.data.b64s 
        captchaid = data.data.id 
    }
}
generate_captcha()