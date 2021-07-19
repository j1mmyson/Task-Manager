
function clickEditButton(obj) {
    card = obj.parentNode.parentNode.parentNode;
    t = card.querySelector('.titleInput');
    c = card.querySelector('.contentInput');

    b1 = card.querySelector('.box1');
    b2 = card.querySelector('.box2');

    t.readOnly = !t.readOnly;
    c.readOnly = !c.readOnly;
    b1.classList.toggle('invisible');
    b2.classList.toggle('invisible');
}

