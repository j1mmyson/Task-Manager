
doneForm = document.getElementById("done_form");
inprogressForm = document.getElementById("inprogress_form");
todoForm = document.getElementById("todo_form");

doneAdd = document.getElementById("doneAddCard");
inprogressAdd = document.getElementById("inProgressAddCard");
todoAdd = document.getElementById("toDoAddCard");


function handleForm(event) {
    console.log("handleForm()!")
    switch (event.target.id){
        case "doneAddCard":
            console.log("doneform!")
            doneForm.classList.toggle('invisible');
            event.target.classList.toggle('invisible');
            break
        case "inProgressAddCard":
            inprogressForm.classList.toggle('invisible');
            event.target.classList.toggle('invisible');
            break
        case "toDoAddCard":
            todoForm.classList.toggle('invisible');
            event.target.classList.toggle('invisible');
            break
    }
}

doneAdd.addEventListener('click', handleForm);
inprogressAdd.addEventListener('click', handleForm);
todoAdd.addEventListener('click', handleForm);
