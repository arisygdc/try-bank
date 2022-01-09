var api = "http://127.0.0.1:8080/api/v1"

function Login() {
    var form = document.getElementById('login')
    var login = {
      register_number: form.elements['register_number'].value,
      pin: form.elements['pin'].value
    }

    console.log(login)

    $.ajax({
      url: api+"/login",
      contentType : 'application/json',
      type : 'POST',
      data: JSON.stringify(login),
      success: function () {
        alert("Thanks!"); 
      }
    })
}

// function VAPay() {
//     axios.get(API_SERVER + '/todos', { withCredentials: true })
// }