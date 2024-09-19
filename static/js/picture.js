let downloadbtn = document.getElementById("testimg");

let imgThumb = document.getElementById("imgThumb");
let imgModal = document.getElementById("imgModal");
let imgDetail = document.getElementById("imgDetail");


imgThumb.addEventListener("click", function(e) {
    e.preventDefault();
    imgDetail.src = e.target.href 
    // imgModal.aria-hidden = "false"

})

downloadbtn.addEventListener("click", function(e) {
    e.preventDefault();
    // e.target.href

    fetch(e.target.href, {
        method: "GET",
        mode: "no-cors",
        referrerPolicy: "no-referrer",
    })
    .then(res => res.blob().
    then(blob => {
    var a = document.createElement('a');
    var url = window.URL.createObjectURL(blob);
    a.href = url;
    // var filename = 'myfile.zip';
    // a.download = filename;
    a.click();
    window.URL.revokeObjectURL(url);
    })
    )
})

// fetch('http://somehost/somefile.zip').then(res => res.blob().then(blob => {
//     var a = document.createElement('a');
//     var url = window.URL.createObjectURL(blob);
//     var filename = 'myfile.zip';
//     a.href = url;
//     a.download = filename;
//     a.click();
//     window.URL.revokeObjectURL(url);
// }))