

    
let newReservaButton = document.querySelector('[id="new-reservation-button"]');

let reportButton = document.querySelector('[id="generate-report-button"]');

BeginPage();



newReservaButton.addEventListener("click",()=>{

    Swal.fire({
        title: "New Reservation",
        text: "Do you want to make a new reservation?",
        icon: "question",
        showCancelButton: true,
        cancelButtonText: "Cancel",


    }).then((result)=>{

        if (result.isConfirmed) {

            window.location.assign('reservations.html');

        }

    })
})
