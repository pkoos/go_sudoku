function loadJavaScript() {
    htmxConfigRequest();
    hidePlaceholder();
    console.log("JavaScript should be loaded....")
}

function hidePlaceholder() {
    var s = document.getElementById("sel_placeholder");
    s.style.display = "none";

}

function htmxConfigRequest() {
    document.body.addEventListener("htmx:configRequest", function (event) {
        let pathWithParameters = event.detail.path.replace(/:([A-Za-z0-9_]+)/g, function (_match, parameterName) {
          let parameterValue = event.detail.parameters[parameterName]
          delete event.detail.parameters[parameterName]
    
          return parameterValue
        })
    
        event.detail.path = pathWithParameters
    })
}



document.addEventListener("DOMContentLoaded", loadJavaScript);
