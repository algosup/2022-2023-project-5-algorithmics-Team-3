function downloadPreset() {
    const preset = {
        name: 'My preset',
        description: 'My preset description',
        version: 1,
        tanks: localStorage.getItem('tanks') ? JSON.parse(localStorage.getItem('tanks')) : [],
        formula: localStorage.getItem('formula') ? JSON.parse(localStorage.getItem('formula')) : [],
        wines: localStorage.getItem('wines') ? JSON.parse(localStorage.getItem('wines')) : [],
    };
    const dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(preset));
    const downloadAnchorNode = document.createElement('a');
    downloadAnchorNode.setAttribute("href", dataStr);
    downloadAnchorNode.setAttribute("download", "preset.json");
    document.body.appendChild(downloadAnchorNode); // required for firefox
    downloadAnchorNode.click();
    downloadAnchorNode.remove();
}

const presetFileInput = document.getElementById('presetFileInput');
presetFileInput.addEventListener('change', uploadPreset, false);

function uploadPreset() {
    const file = this.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = function (event) {
            const preset = JSON.parse(event.target.result);
            localStorage.setItem('tanks', JSON.stringify(preset.tanks));
            localStorage.setItem('formula', JSON.stringify(preset.formula));
            localStorage.setItem('wines', JSON.stringify(preset.wines));
            location.reload();
        };
        reader.readAsText(file);
    }
}