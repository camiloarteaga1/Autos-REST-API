let homeButton = document.querySelector('[id="home-icon"]');

let payButton = document.querySelector('[id="pay-button"]');


BeginReportPage();

payButton.addEventListener("click",()=>{

    Swal.fire({
        title: "Pago",
        text: `Deseas proceder con el pago?`,
        icon: "question",
        showCancelButton: true,
        cancelButtonText: "Cancelar",

    }).then((result)=>{

        if (result.isConfirmed) {

            Swal.fire("¡Gracias!", "Pago correctamente realizado.", "success").then((result2)=>{

                window.location.assign('home.html');
            })

        }

    })
})

homeButton.addEventListener("click",()=>{

    window.location.assign('home.html');

})

async function GetReport(){

    const reporte = await UserReportQuery();

    if(reporte.subTotal === 0){

        Swal.fire({
            title: "Comienza a reservar YA!",
            text: "Por ahora NO tienes reservas...",
            icon: "info",
            timer: 3200,
            timerProgressBar: true,
            showConfirmButton: false

        })

        payButton.style.display = 'none';

    }else{

        DisplayReportData(reporte);
    }
}

async function UserReportQuery(){

    reportOBJCT = {
        email: sessionStorage.getItem('email')
    };

    try {

        const response = await axios.post('http://localhost:8080/report', reportOBJCT);

        return response.data;

    } catch (error) {

        console.error(`API Error: ${error}`);

    }
}


async function BeginReportPage(){

    await GetReport();

}


function DisplayReportData(reportOBJCT){

    let cardContainer = document.querySelector('[id="details-container"]');

    let totalPrice = document.querySelector('[id="precio-total"]');

    totalPrice.innerHTML = reportOBJCT.precioTotal;

    for( carElement of reportOBJCT.details ){

        const reservaResumenCard = document.createElement("div");

        reservaResumenCard.setAttribute("class", "detail-reserva");

        reservaResumenCard.innerHTML = `<div class="detail-reserv-img" style='background-image: url("${carElement.img}");width:190px;height:100px;background-position:center;background-size:cover;'></div>
                                        <div class="detail-textBox">
                                        <div class="detail-textContent">
                                            <p class="detail-car-name">${carElement.brand} ${carElement.carType} (${carElement.modelYear})</p>
                                        </div>
                                        <p class="detail-p">${carElement.transmission} | ${carElement.price} $</p>`;

        cardContainer.appendChild(reservaResumenCard);

    }

}