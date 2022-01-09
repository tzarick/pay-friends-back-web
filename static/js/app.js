console.log("Running Pay Friends Back...");

const resultDivId = "results";

function evenUp() {
  console.log("Evening up...");

  const nameElements = document.getElementsByClassName("friendName");
  const friends = Array.from(nameElements).map((item) => item.value);
  console.log("friends: "  + friends);

  const amountSpentElements = document.getElementsByClassName("amountSpent");
  const amountsSpent = Array.from(amountSpentElements).map((item) => item.value).map((item, i) => {
    // default to 0 if the value was left blank
    if (item === "") {
      amountSpentElements[i].value = 0;
      item = "0"
    }
    return item;
  });
  console.log("amounts spent: " + amountsSpent);
  
  axios.post('/evenUp', {
    data: {
      friends: friends,
      amounts: amountsSpent
     }
  }).then((response) => {
    // expected tx format: "[name1] pays [name2] $[x]"
    const resultsDiv = document.getElementById(resultDivId);
    
    // check response.ok to see if the server successfully processed our request.
    if (response.data.ok) {
      resultsDiv.innerHTML = response.data.transactions.join("\n")
        // .map(item => { // if we want to make the names show up as bold
        //   const firstNameBold = "<strong>" + item.slice(0, item.indexOf("pays")) + "</strong>";
        //   const secondNameBold = "<strong>" + item.slice(item.indexOf("pays") + 4, item.indexOf("$")) + "</strong>";
        //   return firstNameBold + "pays" + secondNameBold + item.slice(item.indexOf("$"));
        // }) 
        // .map(item => "<li>" + item + "</li>") // if we want a bulleted list
    } else {
      resultsDiv.innerHTML = response.data.errorMsg;
    }

    // flash the result div to signal that something happened (useful when the result has not changed)
    changeDivColor(resultDivId, "#4a4646"); 
    setTimeout(() => { 
        changeDivColor(resultDivId, "#272727"); // then change it back
    }, 100);
  })
}

// add a new input row at the bottom
function addInputRow() {
  const inputRow = document.getElementsByClassName("txInput")[0];

  const newInputRow = inputRow.cloneNode(true);
  const newInputElements = newInputRow.getElementsByClassName("txInputElement");
  
  // zero out the input values
  for (const node of newInputElements) {
    node.value = "";
  }

  document.getElementById("inputRows").append(newInputRow);
  newInputRow.getElementsByClassName("friendName")[0].focus(); // focus on the friend name as this is probably where the user would click next
}

// remove the last input row
function removeInputRow() {
  const rows = document.getElementsByClassName("txInput");
  if (rows.length > 1) { // don't do anything if there's only one left
    rows[rows.length - 1].remove();
  }
}

function changeDivColor(id, color) {
  document.getElementById(id).style.backgroundColor = color;
}