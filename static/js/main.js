$(function(){
  Messenger.options = {
    extraClasses: "messenger-fixed messenger-on-bottom messenger-on-left",
    theme: "air"
  };
  var steps = function() {
      var msg = Messenger().post({
        message: '您有新的客户啦!',
        type: 'info',
      });
      //设置标签播放

      setTimeout(
        function(){
          msg.hide();
          //清除播放
        }, 
        6000
      );
    };
    steps();
});
document.getElementById("save").onclick=function(){
  //发送Ajax查询请求并处理
  var request = new XMLHttpRequest();
  request.open("GET", "/admin/message?time=");
  request.send();
  request.onreadystatechange = function(){
      if(request.readyState === 4){
          if(request.status === 200){
              document.getElementById("").innerHTML=request.responseText;
          }else{
              alert("发生错误"+request.status);
          }
      }
  }
}