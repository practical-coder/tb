function CpuTemp() {
  if ("WebSocket" in window) {
    alert("Websockets are supported by your Browser")
    var conn = new WebSocket("ws://localhost:12345/cputemp");
    conn.onopen = function(){
      alert("Websocket connection is opened ...")
    }
    conn.onmessage = function(evt) {
      var message = evt.data;
      document.getElementById("cputemp").innerHTML = "<pre>" + message + "</pre>"
      conn.send("Message received")
    }
    conn.onclose = function(){
      alert("Websocket connection is closed ...")
    }
  }
}