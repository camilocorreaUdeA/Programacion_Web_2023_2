async function consumeApiWithAxios(url){
    try{
        const response = await axios.get(url);
        console.log(`la petición a la api se completo correctamente con status: ${response.status}`);
        return await response.data;
    } catch(error){
        console.error(`fallo la petición a la api con error: ${error.message}`);
    }    
}

const respuestaPeticion = consumeApiWithAxios('https://reqres.in/api/users?page=2');

procesarDatosRespuesta(respuestaPeticion);

/* Acceder a los datos en la respuesta de la peticion */

async function procesarDatosRespuesta(resp){
    const respApi = await resp;
    const datos = respApi.data;
    const main = document.querySelector('main');
    const contenedor = document.createElement('div');
    contenedor.setAttribute('class', 'container');
    for(dato of datos){
        const user = document.createElement('div');
        user.setAttribute('class', 'user');
        const userImg = document.createElement('div');
        userImg.setAttribute('class', 'user-pic');
        const img = document.createElement('img');
        img.setAttribute('src', dato.avatar);
        const userName = document.createElement('div');
        userName.setAttribute('class', 'user-name');
        const userEmail = document.createElement('div');
        userEmail.setAttribute('class', 'user-mail');
        const nombreCompleto = document.createElement('p');
        nombreCompleto.innerHTML = `<strong>Nombre: </strong>${dato.first_name} ${dato.last_name}`;
        const email = document.createElement('p');
        email.innerHTML = `<strong>Email: </strong>${dato.email}`;
        userName.appendChild(nombreCompleto);
        userEmail.appendChild(email);
        userImg.appendChild(img);
        user.appendChild(userImg);
        user.appendChild(userName);
        user.appendChild(userEmail);
        contenedor.appendChild(user);
    }
    main.appendChild(contenedor);
}
