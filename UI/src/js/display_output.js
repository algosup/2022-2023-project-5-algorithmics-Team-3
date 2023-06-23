fileInput = document.getElementById("csvFileInput");
fileInput.addEventListener("change", function (event) {
    var reader = new FileReader();
    reader.onload = function (event) {
        var contents = event.target.result;
        cache = formatInput(contents);
        generateHTML(cache);
    };
    reader.onerror = function (event) {
        console.error("File could not be read! Code " + event.target.error.code);
    };
    reader.readAsText(event.target.files[0]);
}, false);

function formatInput(input) {
    const preprocessedInput = input
        .replace(/\{/g, '[')
        .replace(/\}/g, ']')
        .replace(/\]\[/g, '],[')
        .replace(/(\d+)\s(\d+)\s(\d+)/g, '{"startTank": $1, "destTank": $2, "volume": $3}');

    try {
        const parsedInput = JSON.parse(preprocessedInput);
        const stepsData = [];
        parsedInput.forEach((substeps) => {
            const step = [];
            substeps.forEach((substep) => {
                step.push(substep);
            });
            stepsData.push(step);
        });
        cache = stepsData;
        return stepsData;
    } catch (error) {
        console.error("Error parsing input:", error);
        return null;
    }
}

function generateHTML(cache) {
    let html = `<h2>Here are the instructions you need to follow to complete your mix:</h2>
                <div class="steps_container">`;
    cache.forEach((step, index) => {
        html += `<div class="step_container">
                  <h2>Step ${index + 1}:</h2>`;
        step.forEach((substep, subIndex) => {
            html += `<div class="substep">
                      <p>${index+1}.${subIndex+1}: Empty ${substep[0].volume}L from tank number ${substep[0].startTank} into tank number ${substep[0].destTank}.</p>
                  </div>`;
        });
        html += '</div>';
    });
    html += '</div>';

    document.getElementsByClassName("instructions_container")[0].innerHTML = html;
    return html;
}