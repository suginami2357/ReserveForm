window.addEventListener('DOMContentLoaded', ()=> {
  let d = new Date();
  let yyyy = d.getFullYear();
  let mm = ("0" + (d.getMonth() + 1)).slice(-2);
  let dd = ("0" + d.getDate()).slice(-2);
  form.date.value = yyyy + '-' + mm + '-' + dd;
  
  form.button.addEventListener('click', function() { 
    if(form.PlaceID.value == ""){
        alert("内容を選択してください。")
        return;
      }

    if (!confirm('予約を確定しますか？')) {
      return
    } 
    form.submit();
  });
});

