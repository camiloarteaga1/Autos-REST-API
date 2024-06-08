


let newReservaButton = document.querySelector('[id="new-reservation-button"]');

let reportButton = document.querySelector('[id="generate-report-button"]');

BeginPage();



newReservaButton.addEventListener("click",()=>{

    Swal.fire({
        title: "Nueva Reserva",
        text: "Deseas hacer una nueva reserva?",
        icon: "question",
        showCancelButton: true,
        cancelButtonText: "Cancelar",


    }).then((result)=>{

        if (result.isConfirmed) {

            window.location.assign('reservations.html');

        }

    })
})


reportButton.addEventListener("click",()=>{

    Swal.fire({
        title: "Reporte",
        text: "Deseas generar tu reporte?",
        icon: "question",
        showCancelButton: true,
        cancelButtonText: "Cancelar",


    }).then((result)=>{

        if (result.isConfirmed) {

            window.location.assign('report.html');

        }

    })
})


async function DeleteReservation(id){


    try {

        await axios({
            method: 'delete',
            url: 'http://localhost:8080/reservations/delete',
            data:{
                reservaID: id
            }
        })

        return true;

    } catch (error) {

        console.error(`API Error: ${error}`);

    }

}

async function UpdateUserDataQuery(){

    loginOBJCT = {

        email : sessionStorage.getItem("email"),
        password : sessionStorage.getItem("password")

    }

    try {

        const response = await axios.post('http://localhost:8080/login', loginOBJCT);

        return response.data;

    } catch (error) {

        console.error(`API Error: ${error}`);

    }
}

async function UpdateUserData(){

    const loginAttempt =  await UpdateUserDataQuery();

    sessionStorage.setItem('email',loginAttempt.email);
    sessionStorage.setItem('nombreUsuario',loginAttempt.username);
    sessionStorage.setItem('reservas',JSON.stringify(loginAttempt.reservas));
    sessionStorage.setItem('usuarioID',loginAttempt.ID);
    sessionStorage.setItem('password',loginAttempt.password);
}

function DispalyReservas(reservasOBJCT) {

    if(reservasOBJCT[2] == null){

        Swal.fire({
            title: "Comienza a reservar YA!",
            text: "Por ahora NO tienes reservas...",
            icon: "info",
            timer: 3200,
            timerProgressBar: true,
            showConfirmButton: false

        })

    }else{


        let reservasJSONArray = JSON.parse(reservasOBJCT);

        let mainContainerReservas = document.querySelector('[id="resumen-container"]');

        for( reserv of reservasJSONArray){

            const reservaCard = document.createElement("div");

            reservaCard.setAttribute("class", "card-reserva");



            const resumenReserv = document.createElement("p");

            resumenReserv.setAttribute("class", "resumen-name");

            const deleteIcon = document.createElement("button");

            deleteIcon.setAttribute("class", "delete-reserva");

            deleteIcon.setAttribute("id", reserv.IDr);

            deleteIcon.innerHTML = `<i class="fa-solid fa-trash" id="${reserv.IDr}"></i>`;

            resumenReserv.innerHTML = `<i class="fa-solid fa-car"></i>  ${reserv.summary}`;

            reservaCard.appendChild(resumenReserv);
            reservaCard.appendChild(deleteIcon);

            mainContainerReservas.appendChild(reservaCard);

        }


        //Proceso de eliminación de la card

        let reservaClicked = document.querySelectorAll('button[class="delete-reserva"]');

        for(elementoReserva of reservaClicked){


            elementoReserva.addEventListener("click",(clickedButton)=>{

                Swal.fire({
                    title: "Eliminar Reserva",
                    text: "Deseas eliminar esta reserva?",
                    icon: "warning",
                    showCancelButton: true,
                    cancelButtonText: "Cancelar",


                }).then((result)=>{

                    if (result.isConfirmed) {

                        let reserva_id

                        let card2beDeleted

                        if(clickedButton.target.className == 'delete-reserva'){

                            reserva_id = clickedButton.target.id;

                            card2beDeleted = clickedButton.target.parentNode;

                        }else{

                            reserva_id = clickedButton.target.parentNode.id;

                            card2beDeleted = clickedButton.target.parentNode.parentNode;
                        }

                        DeleteReservation(parseInt(reserva_id));
                        card2beDeleted.remove();
                        window.location.reload();


                    }

                })


            })

        }



    }
}

async function BeginPage(){


    await UpdateUserData();


    let reservationsArray = sessionStorage.getItem("reservas");


    DispalyReservas(reservationsArray);
}