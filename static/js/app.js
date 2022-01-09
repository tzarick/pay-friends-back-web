console.log("Running Pay Friends Back...");

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
    console.log(JSON.stringify(response.data));
    console.log(response.data.transactions);
    const resultsDiv = document.getElementById("results");
    
    // check response.ok to see if the server successfully processed our request.
    // if yes, add the txList data to the output display
    // if no, add the errorMsg to the output display
    if (response.data.ok) {
      resultsDiv.innerHTML = response.data.transactions
        // .map(item => { // make the names show up as bold
        //   const firstNameBold = "<strong>" + item.slice(0, item.indexOf("pays")) + "</strong>";
        //   const secondNameBold = "<strong>" + item.slice(item.indexOf("pays") + 4, item.indexOf("$")) + "</strong>";
        //   return firstNameBold + "pays" + secondNameBold + item.slice(item.indexOf("$"));
        // }) 
        .join("\n");
    } else {
      resultsDiv.innerHTML = response.data.errorMsg;
    }
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
}

// remove the last input row
function removeInputRow() {
  const rows = document.getElementsByClassName("txInput");
  if (rows.length > 1) { // don't do anything if there's only one left
    rows[rows.length - 1].remove();
  }
}