var tableRows = [];

function renderTable() {
    var tableBody = document.getElementById("tableBody");
    tableBody.innerHTML = "";

    var fragment = document.createDocumentFragment();

    for (var i = 0; i < tableRows.length; i++) {
        var row = tableRows[i];

        var newRow = document.createElement("tr");

        var capacityData = document.createElement("td");
        capacityData.textContent = row.capacity;
        newRow.appendChild(capacityData);

        var numberData = document.createElement("td");
        numberData.textContent = row.number;
        newRow.appendChild(numberData);

        var actionsData = document.createElement("td");

        var duplicateButton = document.createElement("button");
        duplicateButton.innerHTML = '<i class="fa-solid fa-copy"></i>';
        duplicateButton.addEventListener("click", (function (index) {
            return function () {
                duplicateRow(index);
            };
        })(i));
        actionsData.appendChild(duplicateButton);

        var deleteButton = document.createElement("button");
        deleteButton.innerHTML = '<i class="fa-solid fa-trash"></i>';
        deleteButton.addEventListener("click", (function (index) {
            return function () {
                deleteRow(index);
            };
        })(i));
        actionsData.appendChild(deleteButton);

        newRow.appendChild(actionsData);
        fragment.appendChild(newRow);
    }

    tableBody.appendChild(fragment);

    generateCsv();
    renderTanks();

    
}
clearButton = document.getElementById("clearButton");
clearButton.addEventListener("click", clearTable);

function clearTable() {
    tableRows = [];
    renderTable();
}

function duplicateRow(index) {
    var newRow = Object.assign({}, tableRows[index]);
    tableRows.push(newRow);
    renderTable();
}

function deleteRow(index) {
    tableRows.splice(index, 1);
    renderTable();
}

document.getElementById("inputAdd").addEventListener("click", function (event) {
    event.preventDefault();

    var capacityField = document.getElementById("capacityField");
    var numberField = document.getElementById("numberField");

    var capacityValue = capacityField.value;
    var numberValue = numberField.value;

    if (capacityValue && numberValue) {
        var newRow = {
            capacity: capacityValue,
            number: numberValue
        };

        tableRows.push(newRow);
        renderTable();
    }

    capacityField.value = '';
    numberField.value = '';
});

function generateCsv() {
    const downloadBtn = document.getElementById('downloadBtn');
    downloadBtn.innerHTML = '';

    const csv = tableRows.map(row => `${row.capacity},${row.number}`).join('\n');

    const downloadLink = document.createElement('a');
    downloadLink.href = `data:text/csv;charset=utf-8,${encodeURI(csv)}`;
    downloadLink.download = 'data.csv';
    downloadLink.innerHTML = 'Download CSV <i class="fa-solid fa-download"></i>';
    downloadBtn.appendChild(downloadLink);
}

function renderTanks() {
    const container = document.getElementById('tanks_display');
    container.innerHTML = '';

    const wineColors = {
        '0': '#c2c2c2',
        '1': '#ff5e5e',
        '2': '#ffb15e',
        '3': '#ffeb5e',
        '4': '#5eff7a',
        '5': '#5effe0',
        '6': '#5e9eff'
    };

    const fragment = document.createDocumentFragment();

    tableRows.forEach(element => {
        const tankDiv = document.createElement('div');
        tankDiv.className = 'tank';
        tankDiv.style.height = element.capacity * 3 + 'px';
        tankDiv.style.backgroundColor = wineColors[element.number];

        const tankLabel = document.createElement('p');
        tankLabel.className = 'tankLabel';
        tankLabel.innerHTML = `${element.capacity}hL<br>Contains: ${element.number}`;

        tankDiv.appendChild(tankLabel);
        fragment.appendChild(tankDiv);
    });

    container.appendChild(fragment);
}

var tableView = document.getElementById('tableView');
var tanksView = document.getElementById('tanksView');
var tableViewElements = document.getElementsByClassName('tableView');
var tanksViewElements = document.getElementsByClassName('tanksView');

function handleClick(activeElement, inactiveElement, activeElements, inactiveElements) {
    activeElement.className = 'active';
    inactiveElement.className = '';
    activeElements[0].style.display = 'block';
    inactiveElements[0].style.display = 'none';
}

tableView.addEventListener('click', function () {
    handleClick(tableView, tanksView, tableViewElements, tanksViewElements);
});

tanksView.addEventListener('click', function () {
    handleClick(tanksView, tableView, tanksViewElements, tableViewElements);
});