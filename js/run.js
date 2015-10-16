// Get context with jQuery - using jQuery's .get() method.
var ctx = $("#myChart").get(0).getContext("2d");


var data = {
  labels: ["January", "February", "March", "April", "May", "June", "July"],
  datasets: [{
    label: "My First dataset",
    fillColor: "rgba(220,220,220,0.5)",
    strokeColor: "rgba(220,220,220,0.8)",
    highlightFill: "rgba(220,220,220,0.75)",
    highlightStroke: "rgba(220,220,220,1)",
    data: [65, 59, 80, 81, 56, 55, 40]
  }, {
    label: "My Second dataset",
    fillColor: "rgba(151,187,205,0.5)",
    strokeColor: "rgba(151,187,205,0.8)",
    highlightFill: "rgba(151,187,205,0.75)",
    highlightStroke: "rgba(151,187,205,1)",
    data: [28, 48, 40, 19, 86, 27, 90]
  }]
};

var options = {};
var myBarChart = new Chart(ctx).Bar(data, options);

document.getElementById('txtFileUpload').addEventListener('change', upload, false);

// Method that checks that the browser supports the HTML5 File API
function browserSupportFileUpload() {
  var isCompatible = false;
  if (window.File && window.FileReader && window.FileList && window.Blob) {
    isCompatible = true;
  }
  return isCompatible;
}

// Method that reads and processes the selected file
function upload(evt) {
  if (!browserSupportFileUpload()) {
    alert('The File APIs are not fully supported in this browser!');
  } else {
    var data = null;
    var file = evt.target.files[0];
    var reader = new FileReader();
    reader.readAsText(file);
    reader.onload = function(event) {
      var csvData = event.target.result;
      var jsonData = CSV2JSON(csvData);
      console.log(jsonData);
    };
    reader.onerror = function() {
      alert('Unable to read ' + file.fileName);
    };
  }
}

//var csv is the CSV file with headers
function CSVToArray(strData, strDelimiter) {
    // Check to see if the delimiter is defined. If not,
    // then default to comma.
    strDelimiter = (strDelimiter || ",");
    // Create a regular expression to parse the CSV values.
    var objPattern = new RegExp((
    // Delimiters.
    "(\\" + strDelimiter + "|\\r?\\n|\\r|^)" +
    // Quoted fields.
    "(?:\"([^\"]*(?:\"\"[^\"]*)*)\"|" +
    // Standard fields.
    "([^\"\\" + strDelimiter + "\\r\\n]*))"), "gi");
    // Create an array to hold our data. Give the array
    // a default empty first row.
    var arrData = [[]];
    // Create an array to hold our individual pattern
    // matching groups.
    var arrMatches = null;
    // Keep looping over the regular expression matches
    // until we can no longer find a match.
    while (arrMatches = objPattern.exec(strData)) {
        // Get the delimiter that was found.
        var strMatchedDelimiter = arrMatches[1];
        // Check to see if the given delimiter has a length
        // (is not the start of string) and if it matches
        // field delimiter. If id does not, then we know
        // that this delimiter is a row delimiter.
        if (strMatchedDelimiter.length && (strMatchedDelimiter != strDelimiter)) {
            // Since we have reached a new row of data,
            // add an empty row to our data array.
            arrData.push([]);
        }
        // Now that we have our delimiter out of the way,
        // let's check to see which kind of value we
        // captured (quoted or unquoted).
        if (arrMatches[2]) {
            // We found a quoted value. When we capture
            // this value, unescape any double quotes.
            var strMatchedValue = arrMatches[2].replace(
            new RegExp("\"\"", "g"), "\"");
        } else {
            // We found a non-quoted value.
            var strMatchedValue = arrMatches[3];
        }
        // Now that we have our value string, let's add
        // it to the data array.
        arrData[arrData.length - 1].push(strMatchedValue);
    }
    // Return the parsed data.
    return (arrData);
}

function CSV2JSON(csv) {
    var array = CSVToArray(csv);
    var objArray = [];
    for (var i = 1; i < array.length; i++) {
        objArray[i - 1] = {};
        for (var k = 0; k < array[0].length && k < array[i].length; k++) {
            var key = array[0][k];
            objArray[i - 1][key] = array[i][k]
        }
    }

    var json = JSON.stringify(objArray);
    var str = json.replace(/},/g, "},\r\n");

    return str;
}
