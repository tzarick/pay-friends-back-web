console.log("running Pay Friends Back...");

function evenUp() {
  console.log("evening up...");

  const nameElements = document.getElementsByClassName("friendName");
  const friends = Array.from(nameElements).map((item) => item.value);
  console.log("friends: "  + friends);

  const amountSpentElements = document.getElementsByClassName("amountSpent");
  const amountsSpent = Array.from(amountSpentElements).map((item) => item.value);
  console.log("amounts spent: " + amountsSpent);
  
  axios.post('/evenUp', {
    data: {
      friends: friends,
      amounts: amountsSpent
     }
  }).then((response) => {
    console.log(response);
  })


}

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