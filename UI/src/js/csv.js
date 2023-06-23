function generateCSV() {
    csvFormula = '';
    formula.forEach((formulaItem, index) => {
        if (formulaItem.winePercent == 0 || formulaItem.winePercent == '') {
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
    formulaTotal = 0;
    formula.forEach((formulaItem, index) => {
        if (formulaItem.winePercent != 0 || formulaItem.winePercent != '') {
            formulaTotal += parseFloat(formulaItem.winePercent);
        }
    }
    );
    if (formulaTotal != 100) {
        swal({
            title: "Error!",
            text: "The formula must be equal to 100%! now it's " + Math.round(formulaTotal * 100) / 100 + "%",
            icon: "error",
            button: "OK",
        });
        return;
    } else {
        swal({
            title: "Success!",
            text: "Your input file will be downloaded! Please wait...",
            icon: "success",
            button: "OK",
        });
        setTimeout(function () {
            swal.close();
            generateCSV();
            var csvFile = new Blob([csvFormula], { type: "text/csv" });
            var downloadLink = document.createElement("a");
            downloadLink.download = "formula.csv";
            downloadLink.href = window.URL.createObjectURL(csvFile);
            downloadLink.style.display = "none";
            document.body.appendChild(downloadLink);
            downloadLink.click();
        }, 2000);
    }
}