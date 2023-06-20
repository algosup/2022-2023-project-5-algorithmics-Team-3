// DOM Variables
const wineNameInput = document.getElementById('wineInput');
const wineColorInput = document.getElementById('wineColorLabel');
const wineAddBtn = document.getElementById('addWineBtn');
const wineResetBtn = document.getElementById('clearWineBtn');
const wineTableBody = document.querySelector('#wineTable tbody');
const wineBottlesContainer = document.getElementsByClassName('winesBottlesView')[0];

var wines = [];

// function to add a wine to the array
function addWine() {
    wines.push({
        wineName: wineNameInput.value,
        wineColor: wineColorInput.value
    });
    for (var i = 0; i < wines.length; i++) {
        wines[i].wineId = i + 1;
    }
    storeWinesIntoLocalStorage();
    displayWinesInTable();
    displayWinesGraphically();

};

// function to store the wines array in localStorage
function storeWinesIntoLocalStorage() {
    localStorage.setItem('wines', JSON.stringify(wines));
}

// function to get the wines array from localStorage
function getWinesFromLocalStorage() {
    const winesRef = localStorage.getItem('wines');
    if (winesRef) {
        wines = JSON.parse(winesRef);
    }
}

// function to log the wines array and the localStorage wines array
function logWines() {
    console.log(wines);
    console.log(localStorage.getItem('wines'));
}

// function to display the wines in the table
function displayWinesInTable() {
    wineTableBody.innerHTML = '';
    wines.forEach((wine, index) => {
        const newRow = document.createElement('tr');
        newRow.innerHTML = `
            <td>${wine.wineId}</td>
            <td>${wine.wineName}</td>
            <td><span style="color: ${wine.wineColor}; background-color: ${wine.wineColor}; width: 50%; height: 100%; display: inline-block; border-radius: 5px;">${wine.wineColor}</span></td>
            <td><button class="btn btn-danger" onclick="deleteWine(${wine.wineId})">Delete</button></td>
        `;
        wineTableBody.appendChild(newRow);
    });
}

// function to display the wines graphically
function displayWinesGraphically() {
    wineBottlesContainer.innerHTML = '';
    wines.forEach((wine, index) => {
        const newBottle = document.createElement('div');
        const newBottleName = document.createElement('h3');
        newBottle.classList.add('wine-container');
        newBottle.innerHTML = `
        <svg fill="#000000" height="200px" width="200px" version="1.1" id="Capa_1"xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="-51.1 -51.1 613.20 613.20" xml:space="preserve" stroke="#000000" stroke-width="0.00511" transform="matrix(1, 0, 0, 1, 0, 0)rotate(0)"><g id="SVGRepo_bgCarrier" stroke-width="0" transform="translate(0,0), scale(1)"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round" stroke="${wine.wineColor}"stroke-width="90"><g>
        <path
        d="M311.715,157.677l-15.196-15.196c-6.138-6.139-9.519-14.3-9.519-22.981V41.141c4.899-4.31,8-10.619,8-17.641 C295,10.542,284.458,0,271.5,0h-32C226.542,0,216,10.542,216,23.5c0,7.023,3.101,13.332,8,17.641V119.5 c0,8.681-3.381,16.842-9.519,22.981l-15.196,15.196C184.27,172.692,176,192.656,176,213.892V487.5c0,12.958,10.542,23.5,23.5,23.5 h112c12.958,0,23.5-10.542,23.5-23.5V213.892C335,192.656,326.73,172.692,311.715,157.677z M239.5,15h32c4.687,0,8.5,3.813,8.5,8.5 s-3.813,8.5-8.5,8.5h-32c-4.687,0-8.5-3.813-8.5-8.5S234.813,15,239.5,15z M191,327h16.5c4.142,0,7.5-3.358,7.5-7.5 s-3.358-7.5-7.5-7.5H191v-17h32.5c4.142,0,7.5-3.358,7.5-7.5s-3.358-7.5-7.5-7.5H191v-25h48.5c4.687,0,8.5,3.813,8.5,8.5v128 c0,4.687-3.813,8.5-8.5,8.5H191V327z M320,487.5c0,4.687-3.813,8.5-8.5,8.5h-112c-4.687,0-8.5-3.813-8.5-8.5V415h48.5 c12.958,0,23.5-10.542,23.5-23.5v-128c0-12.958-10.542-23.5-23.5-23.5H191v-26.108c0-17.229,6.709-33.426,18.892-45.608 l15.196-15.196c8.972-8.972,13.913-20.9,13.913-33.587V46.987c0.167,0.004,0.332,0.013,0.5,0.013h32c0.168,0,0.333-0.009,0.5-0.013 V119.5c0,12.688,4.941,24.616,13.913,33.587l15.196,15.196C313.291,180.466,320,196.663,320,213.892V487.5z">
        </path>
        <path
        d="M295.5,206.392c-4.142,0-7.5,3.358-7.5,7.5V287.5c0,4.142,3.358,7.5,7.5,7.5s7.5-3.358,7.5-7.5v-73.608 C303,209.75,299.642,206.392,295.5,206.392z">
        </path>
        <path
        d="M278.481,190.911c1.464,1.465,3.384,2.197,5.303,2.197c1.919,0,3.839-0.732,5.303-2.197c2.929-2.929,2.929-7.677,0-10.606 l-15.196-15.196c-2.928-2.929-7.677-2.929-10.606,0c-2.929,2.929-2.929,7.677,0,10.606L278.481,190.911z">
        </path>
        <path
        d="M231.5,464H223v-24.5c0-4.142-3.358-7.5-7.5-7.5s-7.5,3.358-7.5,7.5v32c0,4.142,3.358,7.5,7.5,7.5h16 c4.142,0,7.5-3.358,7.5-7.5S235.642,464,231.5,464z">
        </path>
        <path
        d="M279.5,464h-16c-4.142,0-7.5,3.358-7.5,7.5s3.358,7.5,7.5,7.5h16c4.142,0,7.5-3.358,7.5-7.5S283.642,464,279.5,464z">
        </path>
        </g>
        </g>
        <g id="SVGRepo_iconCarrier">
        <g>
        <path
        d="M311.715,157.677l-15.196-15.196c-6.138-6.139-9.519-14.3-9.519-22.981V41.141c4.899-4.31,8-10.619,8-17.641 C295,10.542,284.458,0,271.5,0h-32C226.542,0,216,10.542,216,23.5c0,7.023,3.101,13.332,8,17.641V119.5 c0,8.681-3.381,16.842-9.519,22.981l-15.196,15.196C184.27,172.692,176,192.656,176,213.892V487.5c0,12.958,10.542,23.5,23.5,23.5 h112c12.958,0,23.5-10.542,23.5-23.5V213.892C335,192.656,326.73,172.692,311.715,157.677z M239.5,15h32c4.687,0,8.5,3.813,8.5,8.5 s-3.813,8.5-8.5,8.5h-32c-4.687,0-8.5-3.813-8.5-8.5S234.813,15,239.5,15z M191,327h16.5c4.142,0,7.5-3.358,7.5-7.5 s-3.358-7.5-7.5-7.5H191v-17h32.5c4.142,0,7.5-3.358,7.5-7.5s-3.358-7.5-7.5-7.5H191v-25h48.5c4.687,0,8.5,3.813,8.5,8.5v128 c0,4.687-3.813,8.5-8.5,8.5H191V327z M320,487.5c0,4.687-3.813,8.5-8.5,8.5h-112c-4.687,0-8.5-3.813-8.5-8.5V415h48.5 c12.958,0,23.5-10.542,23.5-23.5v-128c0-12.958-10.542-23.5-23.5-23.5H191v-26.108c0-17.229,6.709-33.426,18.892-45.608 l15.196-15.196c8.972-8.972,13.913-20.9,13.913-33.587V46.987c0.167,0.004,0.332,0.013,0.5,0.013h32c0.168,0,0.333-0.009,0.5-0.013 V119.5c0,12.688,4.941,24.616,13.913,33.587l15.196,15.196C313.291,180.466,320,196.663,320,213.892V487.5z">
        </path>
        <path
        d="M295.5,206.392c-4.142,0-7.5,3.358-7.5,7.5V287.5c0,4.142,3.358,7.5,7.5,7.5s7.5-3.358,7.5-7.5v-73.608 C303,209.75,299.642,206.392,295.5,206.392z">
        </path>
        <path
        d="M278.481,190.911c1.464,1.465,3.384,2.197,5.303,2.197c1.919,0,3.839-0.732,5.303-2.197c2.929-2.929,2.929-7.677,0-10.606 l-15.196-15.196c-2.928-2.929-7.677-2.929-10.606,0c-2.929,2.929-2.929,7.677,0,10.606L278.481,190.911z">
        </path>
        <path
        d="M231.5,464H223v-24.5c0-4.142-3.358-7.5-7.5-7.5s-7.5,3.358-7.5,7.5v32c0,4.142,3.358,7.5,7.5,7.5h16 c4.142,0,7.5-3.358,7.5-7.5S235.642,464,231.5,464z">
        </path>
        <path
        d="M279.5,464h-16c-4.142,0-7.5,3.358-7.5,7.5s3.358,7.5,7.5,7.5h16c4.142,0,7.5-3.358,7.5-7.5S283.642,464,279.5,464z">
        </path>
        </g>
        </g>
        </svg>`
        newBottleName.innerHTML = wine.wineName;
        wineBottlesContainer.appendChild(newBottle);
        newBottle.appendChild(newBottleName);
    });
}

//function to clear the wine array and the localStorage wine array
function resetWines() {
    wines = [];
    localStorage.setItem("wines", JSON.stringify(wines));
    displayWinesInTable();
    displayWinesGraphically();
    logWines();
}

function deleteWine(wineId) {
    wines = wines.filter((wine) => wine.wineId !== wineId);
    for (var i = 0; i < wines.length; i++) {
        wines[i].wineId = i + 1;
    }
    storeWinesIntoLocalStorage();
    displayWinesInTable();
    displayWinesGraphically();
}

//Button event listeners
wineAddBtn.addEventListener('click', function (event) {
    event.preventDefault();
    addWine();
});

wineResetBtn.addEventListener('click', function (event) {
    swal({
        title: "Are you sure?",
        text: "This will clear all the wines and all the tanks!",
        icon: "warning",
        buttons: true,
        dangerMode: true,
    }).then((willReset) => {
            if (willReset) {
                resetWines();
                resetTanks();
                swal("All wines and tanks have been reset!", {
                    icon: "success",
                });
            } else {
                swal("No wines or tanks have been reset!");
            }
        }
    );
});