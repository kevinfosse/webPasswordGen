console.log("JS Chargé")

const api_password = "http://localhost:8080/api/generate"
var app = document.querySelector("#pass_text")
const cooldown = document.getElementById("t");


document.getElementById("newPassword").addEventListener("click", generatePassword);

  function generatePassword() {

    fetch(api_password)
  .then(response => {
      if (!response.ok) {
          throw Error("ERROR");
      }
      return response.json();
  })
  .then(data => {
  
      const html = `<p>${data.password}</p>`
      app.innerHTML = html
      
  })
  .catch(error => {
      console.log(error);
      app.innerHTML ="erreor"
  });
   
  cooldown.classList.remove("round-time-bar");
  cooldown.offsetWidth;
  cooldown.classList.add("round-time-bar");

}

function copyClipboard() {
    var copyText = app

    /* Select the text field */
    copyText.select();
    copyText.setSelectionRange(0, 99999); /* For mobile devices */
  
     /* Copy the text inside the text field */
    navigator.clipboard.writeText(copyText.value);
  
    /* Alert the copied text */
    alert("Copied the text: " + copyText.value);
}

generatePassword()


console.log(cooldown)

cooldown.addEventListener("animationend", generatePassword);
app.addEventListener("click", copyClipboard);