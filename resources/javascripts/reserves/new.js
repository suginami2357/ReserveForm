window.addEventListener('DOMContentLoaded', ()=> {
  var date = new Date();
  var yyyy = date.getFullYear();
  var mm = ("0"+(date.getMonth()+1)).slice(-2);
  var dd = ("0"+date.getDate()).slice(-2);
  document.getElementById("today").value = yyyy + '-' + mm + '-' + dd;
  
  document.getElementById("confirm_button").addEventListener('click', () =>{
    if (confirm('予約を確定しますか？')) {
      document.getElementById("confirm_form").submit();
    }}
  )
});

