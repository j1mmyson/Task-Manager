
doneForm = document.getElementById("done_form");
inprogressForm = document.getElementById("inprogress_form");
todoForm = document.getElementById("todo_form");

doneAdd = document.getElementById("doneAddCard");
inprogressAdd = document.getElementById("inProgressAddCard");
todoAdd = document.getElementById("toDoAddCard");

doneCancel = document.getElementById("doneCancel");
inprogressCancel = document.getElementById("inprogressCancel");
todoCancel = document.getElementById("todoCancel");


function handleForm(event) {
    switch (event.target.id){
        case "doneAddCard":
        case "doneCancel":
            doneForm.classList.toggle('invisible');
            document.getElementById("doneAddCard").classList.toggle('invisible');
            break
        case "inProgressAddCard":
        case "inprogressCancel":
            inprogressForm.classList.toggle('invisible');
            document.getElementById("inProgressAddCard").classList.toggle('invisible');
            break
        case "toDoAddCard":
        case "todoCancel":
            todoForm.classList.toggle('invisible');
            document.getElementById("toDoAddCard").classList.toggle('invisible');
            break
    }
}

doneAdd.addEventListener('click', handleForm);
inprogressAdd.addEventListener('click', handleForm);
todoAdd.addEventListener('click', handleForm);

doneCancel.addEventListener('click', handleForm);
todoCancel.addEventListener('click', handleForm);
inprogressCancel.addEventListener('click', handleForm);
