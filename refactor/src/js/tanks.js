// DOM Variables
const tankCapacityInput = document.getElementById('capacityField');
const tankContentInput = document.getElementById('numberField');
const tankAddBtn = document.getElementById('inputAdd');
const tankResetBtn = document.getElementById('clearBtn');
const tankDownloadBtn = document.getElementById('downloadBtn');
const tankTableBody = document.querySelector('#tankTable tbody');
const tankDisplayContainer = document.getElementsByClassName('tanks_display')[0];

// Global Variables
var tanks = [];

// Set-up tankContentInput using the wines in localStorage
function generateSelectOptions() {
    var select = document.getElementById('numberField');
    select.innerHTML = '';
    var wines = JSON.parse(localStorage.getItem('wines'));
    var el = document.createElement('option');
    el.textContent = "Empty";
    el.value = "0";
    select.appendChild(el);
    for (var i = 0; i < wines.length; i++) {
        var name = wines[i].wineName;
        var id = wines[i].wineId;
        var el = document.createElement('option');
        el.textContent = name;
        el.value = id;
        select.appendChild(el);
    }
}

// function to add a tank to the array
function addTank() {
    //get the values from the wine type
    var capacity = tankCapacityInput.value;
    var content = tankContentInput.options[tankContentInput.selectedIndex].text;
    var contentId = tankContentInput.value;
    const wines = JSON.parse(localStorage.getItem('wines'));
    var contentLabelColor = '';
    for (var i = 0; i < wines.length; i++) {
        if (contentId == 0) {
            contentLabelColor = '#c2c2c2'
        } else if (contentId == wines[i].wineId) {
            contentLabelColor = wines[i].wineColor;
        }
    }
    tanks.push({
        capacity: capacity,
        content: content,
        contentId: contentId,
        contentLabelColor: contentLabelColor,
    });
    storeTanksIntoLocalStorage();
    displayTanksInTable();
    displayTanksGraphically();
    logTanks();
};

// function to store the tanks array in localStorage
function storeTanksIntoLocalStorage() {
    localStorage.setItem('tanks', JSON.stringify(tanks));
}

// function to get the tanks array from localStorage
function getTanksFromLocalStorage() {
    const tanksRef = localStorage.getItem('tanks');
    if (tanksRef) {
        tanks = JSON.parse(tanksRef);
    }
}

// function to log the tanks array and the localStorage tanks array
function logTanks() {
    console.log(tanks);
    console.log(localStorage.getItem('tanks'));
}

// function to display the tanks in the table
function displayTanksInTable() {
    tankTableBody.innerHTML = '';
    tanks.forEach((tank, index) => {
        const newRow = document.createElement('tr');
        newRow.innerHTML = `
            <td>${tank.capacity}</td>
            <td>${tank.content}</td>
            <td>
                <button onclick="deleteTank(${index})">Delete</button>
                <button onclick="duplicateTank(${index})">Duplicate</button>
            </td>
        `;
        tankTableBody.appendChild(newRow);
    });
}

// function to display the tanks graphically
function displayTanksGraphically() {
    tankDisplayContainer.innerHTML = '';
    tanks.forEach((tank, index) => {
        const newTank = document.createElement('div');
        const tankLabel = document.createElement('p');
        newTank.classList.add('tank');
        newTank.style.height = `${tank.capacity * 2}px`;
        newTank.style.maxHeight = `500px`;
        newTank.style.backgroundColor = `${tank.contentLabelColor}`;
        tankLabel.className = 'tankLabel';
        tankLabel.innerHTML = `${tank.capacity}hL<br>${tank.content}`;
        newTank.appendChild(tankLabel);
        tankDisplayContainer.appendChild(newTank);
    });
}

//function to clear the tanks array and the localStorage tanks array
function resetTanks() {
    tanks = [];
    localStorage.removeItem('tanks');
    displayTanksInTable();
    displayTanksGraphically();
    logTanks();
}

// function to delete a tank from the array
function deleteTank(index) {
    tanks.splice(index, 1);
    storeTanksIntoLocalStorage();
    displayTanksInTable();
    displayTanksGraphically();
    logTanks();
}

// function to duplicate a tank in the array
function duplicateTank(index) {
    var tank = tanks[index];
    var duplicate = {
        capacity: tank.capacity,
        content: tank.content,
        contentId: tank.contentId,
        contentLabelColor: tank.contentLabelColor,
    };
    tanks.splice(index + 1, 0, duplicate);
    storeTanksIntoLocalStorage();
    displayTanksInTable();
    displayTanksGraphically();
    logTanks();
}

// buttons event listeners
tankAddBtn.addEventListener('click', function (event) {
    event.preventDefault();
    addTank();
});

tankResetBtn.addEventListener('click', function (event) {
    swal({
        title: "Are you sure?",
        text: "This will clear all the tanks !",
        icon: "warning",
        buttons: true,
        dangerMode: true,
    }).then((willReset) => {
        if (willReset) {
            resetTanks();
            swal("All tanks and wines have been cleared!", {
                icon: "success",
            });
        } else {
            swal("Your tanks and wines are safe!");
        }
    });
});