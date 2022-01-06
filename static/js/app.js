console.log("running app.js...")

function evenUp() {
  console.log("evening up...");
  axios.post('/evenUp', {
    data: {
      "Id": "3", 
      "Title": "Newly Created Post", 
      "desc": "The description for my new post", 
      "content": "my articles content" 
     },
    params: {
      "animal": "giraffe"
    },
  }).then((response) => {
    console.log(response);
  })
}