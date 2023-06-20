const horizontalBarChartContainer = document.querySelector('.horizontal_bar_chart_container');
const winesList = document.getElementById('wines_percent_list_children');
const winesInputs = document.getElementsByClassName('percent');

const winesData = JSON.parse(localStorage.getItem('wines'));
var formula = [];

function addWineToFormula() {
    formula = [];
    for (var i = 0; i < winesInputs.length; i++) {
        formula.push({
            wineId: winesData[i].wineId,
            wineName: winesData[i].wineName,
            wineColor: winesData[i].wineColor,
            winePercent: winesInputs[i].value,
        });
    }

    localStorage.setItem('formula', JSON.stringify(formula));
}

function displayFormulaInBarChart() {
    horizontalBarChartContainer.innerHTML = '';

    for (var i = 0; i < formula.length; i++) {
        if (formula[i].winePercent == 0) {
            continue;
        } 
        else {
            const bar = document.createElement('div');
            bar.style.width = formula[i].winePercent + '%';
            bar.style.backgroundColor = formula[i].wineColor;
            bar.innerHTML = formula[i].winePercent + '%';
            horizontalBarChartContainer.appendChild(bar);
        }
    }
}

function generateWineFormulaOptions() {
    winesList.innerHTML = '';

    winesData.forEach(wine => {
        const listItem = document.createElement('li');

        const colorDiv = document.createElement('div');
        colorDiv.className = 'color';
        colorDiv.style.backgroundColor = wine.wineColor;

        const wineTypeDiv = document.createElement('div');
        wineTypeDiv.className = 'wineType';
        wineTypeDiv.textContent = wine.wineName;

        const percentInput = document.createElement('input');
        percentInput.className = 'percent';
        percentInput.type = 'number';
        percentInput.min = '0';
        percentInput.max = '100';
        percentInput.step = '0.1';
        percentInput.onchange = function () {
            addWineToFormula();
            displayFormulaInBarChart();
            alertWhenFormulaIsUpTo100();
        }

        const percentSymbol = document.createElement('p');
        percentSymbol.className = 'percentSymbol';
        percentSymbol.textContent = '%';

        listItem.appendChild(colorDiv);
        listItem.appendChild(wineTypeDiv);
        listItem.appendChild(percentInput);
        listItem.appendChild(percentSymbol);
        winesList.appendChild(listItem);
    });
}

function getFormulaFromLocalStorage() {
    const formulaRef = localStorage.getItem('formula');
    if (formulaRef) {
        formula = JSON.parse(formulaRef);
    }
    for (var i = 0; i < formula.length; i++) {
        winesInputs[i].value = formula[i].winePercent;
    }
}

function alertWhenFormulaIsUpTo100() {
    var sum = 0.0;
    for (var i = 0; i < winesInputs.length; i++) {
        sum += parseFloat(winesInputs[i].value);
        console.log(sum);
    }
    if (sum > 100) {
        swal({
            title: "Warning!",
            text: "Your formula is up to 100%!",
            icon: "warning",
            button: "OK",
            timer: 1500,
        });
            
    }   
}
