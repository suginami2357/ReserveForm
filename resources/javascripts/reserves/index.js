window.addEventListener('DOMContentLoaded', ()=> {
  var buttons = document.getElementsByName('confirm_button');
  for (let button of buttons) {
    button.addEventListener('click', (e) =>{
      if (confirm('予約を取り消しますか？')) {
        var form = document.getElementById("confirm_form");
        form.action = e.target.id;
        form.submit();
      }}
    )
  }
});
