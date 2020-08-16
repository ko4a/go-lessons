window.socket = new WebSocket('ws://' + location.host + '/ws');
setUpSocket()

function sendMessage(msg) {
    let len = '' + msg.length
    while (len.length < 5) len += ' '
    socket.send(len + msg)
}

function handleSubmit() {
    let el = document.getElementById('chat-input')
    sendMessage(el.value)
    el.value=''
    return false
}

function setUpSocket() {

    socket.onopen = function(e) {
        console.log("Connected");
    };

    socket.onmessage = function (event) {
        displayMessage(event.data)
    }

    socket.onclose = function(event) {
        if (event.wasClean) {
            console.log(`Connection closed`);
        } else {
            console.log('ERROR: Connection reset');
            console.log('Code: ' + event.code + ' reason: ' + event.reason);
        }
    };

    socket.onerror = function(error) {
        console.log(`ERROR: ${error.message}`);
    };

}

function displayMessage(msg) {
    let container = document.getElementById('container')
    let div = document.createElement('div')
    let textNode = document.createTextNode(msg)
    div.className='message'

    div.appendChild(textNode)
    container.appendChild(div)
}