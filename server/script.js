console.log("JS Chargé")

const api_password = "http://localhost:8080/api/generate"
var app = document.querySelector("#pass_text")
const cooldown = document.getElementById("t");
var slider = document.getElementById("myRange");
var output = document.getElementById("value");





function generatePassword() {

  fetch(api_password)
    .then(response => {
      if (!response.ok) {
        throw Error("ERROR");
      }
      return response.json();
    })
    .then(data => {

      const html = `${data.password}`
      app.innerHTML = html
      output.innerHTML = data.lenght
      slider.value = data.lenght


    })
    .catch(error => {
      console.log(error);
      app.innerHTML = "erreor"
    });

  cooldown.classList.remove("round-time-bar");
  cooldown.offsetWidth;
  cooldown.classList.add("round-time-bar");

}

function copy(that) {
  var inp = document.createElement('input');
  document.body.appendChild(inp)
  inp.value = that.textContent
  inp.select();
  document.execCommand('copy', false);
  inp.remove();
  var temp = that.textContent
  that.classList.toggle('fade')
  that.textContent = "Copied!"
  setTimeout(() => {
    that.textContent = temp;
    that.classList.toggle('fade');
    generatePassword();
  }, 800);



}

slider.oninput = function () {
  output.innerHTML = this.value;

  var dataObject = {
    "lenght": this.value,
  }

  fetch('http://localhost:8080/api/parameter', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(dataObject)
  }).then(response => {
    return response.json()
  }).then(data =>
    // this is the data we get after putting our data,
    app.innerHTML = `${data.password}`
  );
}



generatePassword()


console.log(cooldown)

cooldown.addEventListener("animationend", generatePassword);
app.addEventListener("click", copyClipboard);