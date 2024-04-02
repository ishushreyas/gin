let form = document.querySelector('form');
let imageUploader = document.querySelector('#imageUploader');
let output = document.querySelector('.output');

form.onsubmit = async (ev) => {
    ev.preventDefault()
    output.textContent = "Generating...";

    var data = new FormData(form);
    var request = new XMLHttpRequest();
    request.open("POST", "/api/generate");
    request.onload = function () {
        // Read the response and interpret the output as markdown.
        let md = window.markdownit();
        output.innerHTML = md.render(request.responseText);
    };
    request.send(data);
    return false;
}

imageUploader.onsubmit = async (ev) => {
    ev.preventDefault()
    output.textContent = "Uploading...";

    var data = new FormData(imageUploader);
    var request = new XMLHttpRequest();
    request.open("POST", "/upload");
    request.onload = function () {
        output.innerHTML = request.responseText;
    };
    request.send(data);
    return false;
}