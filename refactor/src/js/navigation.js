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

var wineTableView = document.getElementById('winesTableView');
var wineBottlesView = document.getElementById('winesBottlesView');
var wineTableViewElements = document.getElementsByClassName('winesTableView');
var wineBottlesViewElements = document.getElementsByClassName('winesBottlesView');

function handleClickWines(activeElement, inactiveElement, activeElements, inactiveElements) {
    activeElement.className = 'active';
    inactiveElement.className = '';
    activeElements[0].style.display = 'flex';
    inactiveElements[0].style.display = 'none';
}

wineTableView.addEventListener('click', function () {
    handleClickWines(wineTableView, wineBottlesView, wineTableViewElements, wineBottlesViewElements);
});

wineBottlesView.addEventListener('click', function () {
    handleClickWines(wineBottlesView, wineTableView, wineBottlesViewElements, wineTableViewElements);
});

var navFormula = document.getElementById('nav_formula');
var navTanks = document.getElementById('nav_tanks');
var navWines = document.getElementById('nav_wines');
var navActions = document.getElementById('nav_actions');

var sidebarFormula = document.getElementsByClassName('formula_sidebar')[0];
var sidebarTanks = document.getElementsByClassName('tanks_sidebar')[0];
var sidebarWines = document.getElementsByClassName('wines_sidebar')[0];
var sidebarActions = document.getElementsByClassName('actions_sidebar')[0];

var containerFormula = document.getElementsByClassName('global_container_formula')[0];
var containerTanks = document.getElementsByClassName('global_container_tanks')[0];
var containerWines = document.getElementsByClassName('global_container_wines')[0];
var containerActions = document.getElementsByClassName('global_container_actions')[0];

var lastPosition = 0;

document.getElementById('nav_formula').addEventListener('click', function () {
    location.reload();
    formulaScreen();
});

function formulaScreen() {
    navFormula.classList.add = 'active';
    navTanks.classList.remove = 'active';
    navWines.classList.remove = 'active';
    navActions.classList.remove = 'active';

    sidebarFormula.style.display = 'block';
    sidebarTanks.style.display = 'none';
    sidebarWines.style.display = 'none';
    sidebarActions.style.display = 'none';

    containerFormula.style.display = 'block';
    containerTanks.style.display = 'none';
    containerWines.style.display = 'none';
    containerActions.style.display = 'none';

    lastPosition = 3;
    saveLastPosition();
}

document.getElementById('nav_tanks').addEventListener('click', function () {
    tanksScreen();
});

function tanksScreen(){
    navFormula.classList.remove = 'active';
    navTanks.classList.add = 'active';
    navWines.classList.remove = 'active';
    navActions.classList.remove = 'active';

    sidebarFormula.style.display = 'none';
    sidebarTanks.style.display = 'block';
    sidebarWines.style.display = 'none';
    sidebarActions.style.display = 'none';

    containerFormula.style.display = 'none';
    containerTanks.style.display = 'block';
    containerWines.style.display = 'none';
    containerActions.style.display = 'none';

    generateSelectOptions();

    lastPosition = 2;
    saveLastPosition();
}

document.getElementById('nav_wines').addEventListener('click', function () {
    winesScreen();
});

function winesScreen(){
    navFormula.classList.remove = 'active';
    navTanks.classList.remove = 'active';
    navWines.classList.add = 'active';
    navActions.classList.remove = 'active';

    sidebarFormula.style.display = 'none';
    sidebarTanks.style.display = 'none';
    sidebarWines.style.display = 'block';
    sidebarActions.style.display = 'none';

    containerFormula.style.display = 'none';
    containerTanks.style.display = 'none';
    containerWines.style.display = 'block';
    containerActions.style.display = 'none';

    lastPosition = 1;
    saveLastPosition();
}

document.getElementById('nav_actions').addEventListener('click', function () {
    actionsScreen();
});

function actionsScreen(){
    navFormula.classList.remove = 'active';
    navTanks.classList.remove = 'active';
    navWines.classList.remove = 'active';
    navActions.classList.add = 'active';

    sidebarFormula.style.display = 'none';
    sidebarTanks.style.display = 'none';
    sidebarWines.style.display = 'none';
    sidebarActions.style.display = 'block';

    containerFormula.style.display = 'none';
    containerTanks.style.display = 'none';
    containerWines.style.display = 'none';
    containerActions.style.display = 'block';

    lastPosition = 0;
    saveLastPosition();
}

// When Var LastPosition = 0, it means that the user is on the Actions page
// When Var LastPosition = 1, it means that the user is on the Wines page
// When Var LastPosition = 2, it means that the user is on the Tanks page
// When Var LastPosition = 3, it means that the user is on the Formula page

function saveLastPosition() {
    switch (lastPosition) {
        case 0:
            localStorage.setItem('lastPosition', 'actions');
            break;
        case 1:
            localStorage.setItem('lastPosition', 'wines');
            break;
        case 2:
            localStorage.setItem('lastPosition', 'tanks');
            break;
        case 3:
            localStorage.setItem('lastPosition', 'formula');
            break;
    }
}

function getLastPosition() {
    var lastPosition = localStorage.getItem('lastPosition');
    switch (lastPosition) {
        case 'actions':
            actionsScreen();
            break;
        case 'wines':
            winesScreen();
            break;
        case 'tanks':
            tanksScreen();
            break;
        case 'formula':
            formulaScreen();
            break;
    }
}