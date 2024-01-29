const fileInput = document.querySelector('input[type=file]')
const fileName = document.querySelector('#file_name')
const fileSize = document.querySelector('#file_size')
fileInput.addEventListener('change', (e) => {
    const file = e.target.files[0]
    fileName.value = file.name
    fileSize.value = file.size
})

document
    .querySelector('.send-files-result')
    .addEventListener('htmx:afterSwap', handleFileResultHtmx)

function handleFileResultHtmx(e) {
    const sendFilesBtn = document.querySelector('#send-files-btn')
    sendFilesBtn.disabled = true
    sendFilesBtn.innerText = 'Link generated'

    const fileId = document.querySelector('#send-file').value.split('/').pop()
    const ws = new WebSocket('ws://localhost:8080/send-files/ws')
    const peer = new SimplePeer({
        initiator: true,
        trickle: false
    })

    let websocket_open = false
    let peer_open = false
    let data = null

    peer.on('signal', (data) => {
        console.log(JSON.stringify(data))
        data = JSON.stringify({
            type: 'web_rtc_offer',
            client: 'sender',
            file_id: fileId,
            web_rtc_offer: JSON.stringify(data)
        })

        peer_open = true
        if (!websocket_open) return
        ws.send(data)
    })
    peer.on('connect', () => {
        console.log('connected')
    })
    peer.on('data', (data) => {
        console.log('recieved', data)
    })

    ws.addEventListener('open', () => {
        console.log('connected')
        websocket_open = true
        if (!peer_open) return
        ws.send(data)
    })

    ws.addEventListener('message', (msg) => {
        console.log(msg.data)
        const data = JSON.parse(msg.data)
        if (data.type === 'web_rtc_answer') {
            console.log('received answer')
            peer.signal(JSON.parse(data.web_rtc_answer))
        }
    })

    ws.addEventListener('close', () => {
        console.log('disconnected')
    })
}
