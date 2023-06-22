function root() {
    getLastPosition();

    getWinesFromLocalStorage();
    displayWinesInTable();
    displayWinesGraphically();

    generateSelectOptions();
    
    getTanksFromLocalStorage();
    displayTanksInTable();
    displayTanksGraphically();

    generateWineFormulaOptions();
    getFormulaFromLocalStorage();
    displayFormulaInBarChart();

    // logWines();
    // logTanks();
}

window.addEventListener('load', root);