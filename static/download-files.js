const ws = new WebSocket('ws://localhost:8080/send-files/ws')

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
    peer.send('hello')
})
peer.on('data', (data) => {
    console.log('recieved', data)
})

ws.addEventListener('message', (e) => {
    console.log(e.data)
    const data = JSON.parse(e.data)
    if (data.type === 'web_rtc_offer') {
        peer.signal(JSON.parse(data.web_rtc_offer))
    }
})
