function resize(obj) {
    obj.style.height = obj.scrollHeight;
}


function upsize(obj) {
    textarea = obj.querySelector('.contentInput');
    textarea.style.height = textarea.scrollHeight;
}

function downsize(obj) {
    textarea = obj.querySelector('.contentInput');
    textarea.style.height = "auto";
}
