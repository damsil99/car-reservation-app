document.getElementById('searchForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const criteria = document.getElementById('criteria').value;

    const response = await fetch(`/cars?criteria=${criteria}`);
    const cars = await response.json();

    const carsList = document.getElementById('carsList');
    carsList.innerHTML = '';
    cars.forEach(car => {
        const carDiv = document.createElement('div');
        carDiv.innerHTML = `
            <h3>${car.make} ${car.model}</h3>
            <p>${car.year} - ${car.color}</p>
            <button onclick="reserveCar(${car.id})">Reserve</button>
        `;
        carsList.appendChild(carDiv);
    });
});

async function reserveCar(carId) {
    const token = localStorage.getItem('token');
    const response = await fetch('/reservations', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({ carId, extras: 'none', totalPrice: 100 })
    });

    if (response.ok) {
        alert('Car reserved successfully');
        window.location.reload();
    } else {
        alert('Reservation failed');
    }
}

async function loadReservations() {
    const token = localStorage.getItem('token');
    const response = await fetch('/reservations/me', {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    });

    const reservations = await response.json();
    const reservationsList = document.getElementById('reservationsList');
    reservationsList.innerHTML = '';
    reservations.forEach(reservation => {
        const reservationDiv = document.createElement('div');
        reservationDiv.innerHTML = `
            <h3>Reservation for car ID ${reservation.carId}</h3>
            <p>Extras: ${reservation.extras}</p>
            <p>Total Price: ${reservation.totalPrice}</p>
        `;
        reservationsList.appendChild(reservationDiv);
    });
}

loadReservations();
