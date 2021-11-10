window.addEventListener('DOMContentLoaded', ()=> {
  let buttons = document.getElementsByName('button');
  for(let button of buttons){
    button.addEventListener('click', (e) =>{
      if (!confirm('予約を取り消しますか？')) {
        return
      }
      form.action = "/reserves/" + e.target.id + "/delete";
      form.submit();
    })
  }
});
