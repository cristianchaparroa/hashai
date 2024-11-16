'use client'
import React, { useState, useEffect, useCallback } from 'react'
import mermaid from 'mermaid'
import { useParams } from 'next/navigation'
import styles from '../page.module.css'

export const Mermaid = ({ encodeValue }: any) => {
  // const [mermaidInput, setMermaidInput] = useState(`graph TD
  //   A[Client] --> B[Load Balancer]
  //   B --> C[Server01]
  //   B --> D[Server02]`)
  const [mermaidInput, setMermaidInput] = useState(`
  graph LR
    0xDFd5...963d-->|2.38 mETH|0x8611...7784
    0x8611...7784-->|10.00 mETH|0x7cCD...5DFA
    0x28C6...1d60-->|42.73 mETH|0x8611...7784
    0x8611...7784-->|3.20 mETH|0x2535...303b
    0x8611...7784-->|20.00 mETH|0x0000...bE59
  `)
  const [svg, setSvg] = useState('')
  const [error, setError] = useState('')
  const [decodedOutput, setDecodedOutput] = useState('')

  function encodeUnicodeToBase64(str: string) {
    try {
      const utf8Encoder = new TextEncoder();
      const utf8Array: any = utf8Encoder.encode(str);

      console.log('utf8Array', utf8Array);

      const base64String = btoa(String.fromCharCode(...utf8Array));
      console.log('base64String', base64String);
    } catch (error) {
      console.log('error', error);

      setDecodedOutput('Invalid Base64 string');
    }

    // return base64String;
  }

  const handleDecode = (encodeValue: string) => {
    console.log('encodeValue', encodeValue);
    try {
      let encodedValue: any = encodeValue
      encodedValue = encodedValue.replace(/\s+/g, '');

      // // Handle URL-safe encoding
      encodedValue = encodedValue.replace(/-/g, '+').replace(/_/g, '/');

      // // Add padding
      if (encodedValue.length % 4 !== 0) {
        encodedValue += '='.repeat(4 - (encodedValue.length % 4));
      }

      // Decode safely
      // /try {
      let decodedValue = atob(encodedValue);
      setDecodedOutput(decodedValue);
      console.log('decodedValue', decodedValue);
      // } catch (error: any) {
      //   console.error("Error decoding Base64 string:", error.message);
      // }

      // console.log('decodedStringAtoB', decodedStringAtoB);
    } catch (error) {
      console.log('error', error);

      setDecodedOutput('Invalid Base64 string');
    }
  };

  useEffect(() => {
    if (encodeValue) {
      encodeUnicodeToBase64(encodeValue)
    }
  }, [encodeValue])

  useEffect(() => {
    if (typeof window !== 'undefined' && typeof document !== 'undefined') {
      mermaid.initialize({ startOnLoad: true });
    }
  }, []);

  const renderMermaid = useCallback(async () => {
    try {
      const { svg } = await mermaid.render('mermaid-diagram', mermaidInput)
      setSvg(svg)
      setError('')
    } catch (err) {
      setError('Invalid Mermaid syntax. Please check your input.')
      setSvg('')
    }
  }, [mermaidInput])

  useEffect(() => {
    renderMermaid()
  }, [renderMermaid])

  return (
    <div className={styles.mermaid}>
      {/* <button onClick={renderMermaid}>Render Diagram</button> */}
      {error && <div className="text-red-500">{error}</div>}
      {svg && (
        <div
          className="mermaid bg-white p-4 rounded-lg shadow w-full"
          dangerouslySetInnerHTML={{ __html: svg }}
        />
      )}
      {/* decodedOutput ? decodedOutput : */}
    </div >
  )
}