// 'use client'/
import { getFrameMetadata } from 'frog/web'
import type { Metadata } from 'next'
import Image from 'next/image'
import  MermaidDiagram  from './components/Mermaid'

import styles from './page.module.css'

export async function generateMetadata(): Promise<Metadata> {
  const frameTags = await getFrameMetadata(
    `${process.env.VERCEL_URL || 'http://localhost:3000'}/api`,
  )
  return {
    other: frameTags,
  }
}

type HomeProps = {
  searchParams: { [key: string]: string | undefined };
};

export default function Home({ searchParams }: HomeProps) {
  const diagramBase64 = searchParams.hash;
  console.error(diagramBase64);
  const decodedHash = decodeMermaidDiagram(diagramBase64)

  console.log('encodeValue', decodedHash);

  return (
    <main className={styles.main}>
      <MermaidDiagram diagram={decodedHash} />
    </main>
  )
}


function decodeBase64(base64String: string): string {
  // First, decode the Base64 string to a Uint8Array
  const binaryString = atob(base64String);
  const bytes = new Uint8Array(binaryString.length);

  for (let i = 0; i < binaryString.length; i++) {
    bytes[i] = binaryString.charCodeAt(i);
  }

  // Convert the Uint8Array to a string using TextDecoder
  // Go uses UTF-8 encoding by default
  const decoder = new TextDecoder('utf-8');
  return decoder.decode(bytes);
}

function decodeMermaidDiagram(encodedString) {
  // Use built-in decodeURIComponent to decode the URL-encoded string
  return decodeURIComponent(encodedString);
}
