let ButtonLogin = document.querySelector('[id="logeate-a"]');
let ButtonRegister = document.querySelector('[id="registrate-a"]');

let registerForm = document.querySelector('[id="Register-form"]');
let loginForm = document.querySelector('[id="login-form"]');

let badCreds = document.querySelector('[id="bad-creds"]');

let usrRegistered = document.querySelector('[id="user-registered"]');


let logeame = document.querySelector('[id="loginButton"]');

let registrame = document.querySelector('[id="RegisterButton"]');


let registro_nombreUsuario = document.querySelector('[id="Register-form"] [class="input-name"]');

let registro_emailUsuario = document.querySelector('[id="Register-form"] [class="input-email"]');

let registro_passwordUsuario = document.querySelector('[id="Register-form"] [class="input-password"]');


let login_emailUsuario = document.querySelector('[id="login-form"] [class="input-email"]');

let login_passwordUsuario = document.querySelector('[id="login-form"] [class="input-password"]');



registerForm.style.display = 'none'; 

badCreds.style.display = 'none'; 

usrRegistered.style.display = 'none';


ButtonLogin.addEventListener("click",()=>{

   loginForm.style.display = 'block';
   registerForm.style.display = 'none'

})

ButtonRegister.addEventListener("click",()=>{

    registerForm.style.display = 'block';
    loginForm.style.display = 'none';
    badCreds.style.display = 'none';

})



logeame.addEventListener("click",()=>{

    
    loginData = {

        mail : login_emailUsuario.value,
        password : login_passwordUsuario.value

    }

    
    LogeoUsuario(loginData);

})

registrame.addEventListener("click",()=>{

    registroData = {

        mail : registro_emailUsuario.value,
        username: registro_nombreUsuario.value,
        password : registro_passwordUsuario.value

    }



    RegistroUsuario(registroData);
    
})



async function LoginQuery( loginOBJCT ){


    try {

        const response = await axios.post('http://localhost:8080/login', loginOBJCT);

        return response.data;
        
    } catch (error) {
        
        console.error(`API Error: ${error}`);

    }

};


async function LogeoUsuario(loginOBJCT) {

    const loginAttempt =  await LoginQuery(loginOBJCT);

    if(!loginAttempt){

        badCreds.style.display = 'block';


    }else{

        badCreds.style.display = 'none';


        sessionStorage.setItem('mail',loginAttempt.mail);
        sessionStorage.setItem('username',loginAttempt.username);
        sessionStorage.setItem('reservation',JSON.stringify(loginAttempt.reservation));
        sessionStorage.setItem('usuarioID',loginAttempt.ID);
        sessionStorage.setItem('password',loginAttempt.password);

        Swal.fire({
            title: "Bienvenido,",
            text: `${loginAttempt.username}!`,
            icon: "success"
        }).then((result)=>{

            if (result.isConfirmed) {

                window.location.assign('home.html');

              }

        })

        

        
    }
    
};



async function RegisterQuery( registerOBJCT ){


    try {

        const response = await axios.post('http://localhost:8080/register', registerOBJCT);

        return response.data;
        
    } catch (error) {
        
        console.error(`API Error: ${error}`);

    }

};


async function RegistroUsuario(registerOBJCT) {

    const registerAttempt =  await RegisterQuery(registroData);

    if(registerAttempt){

        usrRegistered.style.display = 'block';  

        Swal.fire({
            title: "Registered Successfully",
            text: `Proceed to log in, ${registerAttempt.username}.`,
            icon: "Success"
        });
    }
    
};


