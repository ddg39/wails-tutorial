import { useState } from 'react'

async function generateQR(text: string): Promise<string | null> {
  if (!text) {
    alert('Please enter some text')
    return null
  }

  try {
    const qrSrc = `/qrservice?text=${encodeURIComponent(text)}`
    return qrSrc
  } catch (err) {
    console.error('Failed to generate QR code:', err)
  }

  return null
}

function App() {
  const [inputText, setInputText] = useState("")
  const [qr, setQR] = useState("")

  function generate() {
    generateQR(inputText).then((qrSrc) => {
      if (!qrSrc) return
      setQR(qrSrc)
    })
  }
  return (
    <div>
      <img src={ qr } />
      <input type="text" placeholder="Enter text" value={inputText} onChange={(e) => setInputText(e.target.value)} />
      <button onClick={() => generate()}>Generate QR Code</button>
    </div>
  )
}

export default App
