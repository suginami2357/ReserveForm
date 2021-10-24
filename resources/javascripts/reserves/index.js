window.addEventListener('DOMContentLoaded', ()=> {
  var buttons = document.getElementsByName('confirm-button');
  for(let button of buttons){
    button.addEventListener('click', (e) =>{
      if (confirm('予約を取り消しますか？')) {
        var form = document.getElementById("confirm-form");
        form.action = "/reserves/" + e.target.id + "/delete";
        form.submit();
      }}
    )
  }
});
