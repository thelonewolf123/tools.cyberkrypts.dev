const ws = new WebSocket('ws://localhost:8080/send-files/ws')
const downloadBtn = document.querySelector('#download-btn')
const fileBuffer = []

ws.addEventListener('open', () => {
    console.log('connected')
    const data = JSON.stringify({
        type: 'get_web_rtc_offer',
        client: 'receiver',
        file_id: location.pathname.split('/').pop()
    })
    ws.send(data)
})

const peer = new SimplePeer({
    initiator: false,
    trickle: false
})
peer.on('signal', (data) => {
    console.log(JSON.stringify(data))
    const message = JSON.stringify({
        type: 'web_rtc_answer',
        client: 'receiver',
        file_id: location.pathname.split('/').pop(),
        web_rtc_answer: JSON.stringify(data)
    })
    ws.send(message)
})
peer.on('connect', () => {
    console.log('connected')
    downloadBtn.disabled = false
    downloadBtn.innerText = 'Download'
})

function downloadFileHandler() {
    const blob = new Blob([fileBuffer], { type: 'application/octet-stream' })
    const url = URL.createObjectURL(blob)
    const fileName = document.querySelector('#file-name').innerText
    const aTag = document.createElement('a')
    aTag.href = url
    aTag.download = fileName
    aTag.click()
}

const fileSizeBytes = parseInt(
    document.querySelector('#file-size-bytes').innerText
)

peer.on('data', (data) => {
    console.log('received', data)
    fileBuffer.push(...data)
    console.log(fileBuffer.length, fileSizeBytes)
    if (fileSizeBytes === fileBuffer.length) {
        downloadFileHandler()
    }
})

ws.addEventListener('message', (e) => {
    console.log(e.data)
    const data = JSON.parse(e.data)
    if (data.type === 'web_rtc_offer') {
        peer.signal(JSON.parse(data.web_rtc_offer))
    }
})

downloadBtn.addEventListener('click', () => {
    ws.send(
        JSON.stringify({
            type: 'start_download',
            client: 'receiver',
            file_id: location.pathname.split('/').pop()
        })
    )
})
