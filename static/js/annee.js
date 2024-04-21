const value1 = document.querySelector("#value")
const input1 = document.querySelector("#pi_input")
value1.textContent = input1.value1
input1.addEventListener("input", (event) => {
  value1.textContent = event.target.value
})
const value = document.querySelector("#valueAlbum")
const input = document.querySelector("#pi_input2")
value.textContent = input.value
input.addEventListener("input", (event) => {
  value.textContent = event.target.value
})

