document.getElementById('pedidoForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const nombre = document.getElementById('nombre').value;
    const establecimiento = document.getElementById('establecimiento').value;
    const productosElements = document.querySelectorAll('.producto');
    const productos = [];

    productosElements.forEach(productoElement => {
        const idProducto = productoElement.querySelector('.producto_id').value;
        const cantidad = productoElement.querySelector('.cantidad').value;
        productos.push({
            ID_Producto: parseInt(idProducto),
            Cantidad: parseInt(cantidad)
        });
    });

    axios.post('http://localhost:8000/pedido', {
        Nombre: nombre,
        ID_Establecimiento: parseInt(establecimiento),
        productos: productos
    })
    .then(response => {
        alert('Pedido creado exitosamente con ID: ' + response.data.ID_Pedido);
    })
    .catch(error => {
        console.error(error);
        alert('Error al crear el pedido');
    });
});

document.getElementById('addProduct').addEventListener('click', function() {
    const productDiv = document.createElement('div');
    productDiv.classList.add('producto');
    productDiv.innerHTML = `
        <label for="producto_id">ID del Producto:</label>
        <input type="number" class="producto_id" name="producto_id" required>
        <label for="cantidad">Cantidad:</label>
        <input type="number" class="cantidad" name="cantidad" required>
    `;
    document.getElementById('productos').appendChild(productDiv);
});

document.getElementById('fetchRanking').addEventListener('click', function() {
    axios.get('http://localhost:8000/ranking')
    .then(response => {
        const rankingList = document.getElementById('rankingList');
        rankingList.innerHTML = '';
        response.data.forEach(establecimiento => {
            const listItem = document.createElement('li');
            listItem.textContent = `${establecimiento.RazonSocial}: ${establecimiento.Total_Vendido}`;
            rankingList.appendChild(listItem);
        });
    })
    .catch(error => {
        console.error(error);
        alert('Error al obtener el ranking');
    });
});

document.getElementById('ubicacionForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const ubicacionId = document.getElementById('ubicacion_id').value;

     // Enviar solicitud GET al backend para obtener la ubicación
     fetch(`http://localhost:8000/ubicacion/${ubicacionId}`)
     .then(response => {
         if (!response.ok) {
             throw new Error('Error al obtener la ubicación');
         }
         return response.json();
     })
     .then(data => {
         // Mostrar la ubicación en el párrafo ubicacionResult
         document.getElementById('ubicacionResult').textContent = `Ubicación: ${data.Ubicacion_Establecimiento}`;
     })
     .catch(error => {
         console.error(error);
         alert('Error al obtener la ubicación');
    });
});

