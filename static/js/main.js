$(document).ready(function(){
  //每过5分钟向服务器发送请求
  var t
  var message = 0
  setInterval(function(){
    var id = $("table tbody tr th").first()[0].childNodes[0].data;
    //console.log(id);
    $.ajax({
      type: "get",
      url:"/admin/message?id="+id,
      async:true,
      success:function(data){
        message = data;
        if (message != 0) {
          $("#message").html(message);//message.responseText
          $("#message").css("background-color", "#2b90d9");
          var audio = $("#audio")[0];
          audio.play();     
        }else{
          $("#message").html("0");
          $("#message").css("background-color", "#ccc");
        }
      }
    });
  }, 300000)
  $("#message").click(function(){
    $("#message").css("background-color","#F8F8F8");
  });
});