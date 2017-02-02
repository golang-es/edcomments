/**
 * Created by alexys on 1/02/17.
 */

let formLogin = $('form-login'),
    email = $('email'),
    password = $('password'),
    // btnLogin = $('btnLogin'),
    mensajeLogin = $('mensaje-login');

formLogin.addEventListener('submit', e => {
    e.preventDefault();
    let obj = {
        email: email.value,
        password: password.value
    };

    peticionAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
        .then(respuesta => {
            if (respuesta.status === 200) {
                mensajeLogin.textContent ='Ingresaste';
                sessionStorage.setItem('token', respuesta.response.token);
                console.log(respuesta.response);
            } else {
                mensajeLogin.textContent = respuesta.response.message;
                console.log(respuesta.response);
            }
        })
        .catch(error => {
            console.log(error);
        });
});



