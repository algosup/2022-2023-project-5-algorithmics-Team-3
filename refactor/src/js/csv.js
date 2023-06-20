

function generateCSV() {
    isFormulaIsEqual100();
    csvFormula = '';
    formula.forEach((formulaItem, index) => {
        if (formulaItem.winePercent == 0 || formulaItem.winePercent == ''){
            return;
        } else {
            csvFormula += `${formulaItem.winePercent},`;
        }
    });
    csvFormula = csvFormula.slice(0, -1);
    csvFormula += '\n';

    tanks.forEach((tank, index) => {
        csvFormula += `${tank.capacity},${tank.contentId}\n`;
    });
    console.log(csvFormula);
}

function downloadCSV() {
    generateCSV();
    var csvFile = new Blob([csvFormula], {type: "text/csv"});
    var downloadLink = document.createElement("a");
    downloadLink.download = "formula.csv";
    downloadLink.href = window.URL.createObjectURL(csvFile);
    downloadLink.style.display = "none";
    document.body.appendChild(downloadLink);
    downloadLink.click();
}

function isFormulaIsEqual100() {
    formulaTotal = 0;
    formula.forEach((formulaItem, index) => {
        if (formulaItem.winePercent == 0 || formulaItem.winePercent == ''){
            return;
        } else {
            formulaTotal += parseInt(formulaItem.winePercent);
        }
    }
    );
    if (formulaTotal != 100) {
        swal({
            title: "Error!",
            text: "The formula must be equal to 100%! now it's " + formulaTotal + "%" ,
            icon: "error",
            button: "OK",
        });
        return false;
    } else {
        return true;
    }
}