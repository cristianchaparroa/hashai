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

  if (!diagramBase64) {
    return (
        <div>
          <p>Incorrect hash</p>
        </div>
    )
  }
  const decodedHash = decodeMermaidDiagram(diagramBase64)

  console.log('encodeValue', decodedHash);

  return (
    <main className={styles.main}>
      <MermaidDiagram diagram={decodedHash} />
    </main>
  )
}

function decodeMermaidDiagram(encodedString:string) {
  // Use built-in decodeURIComponent to decode the URL-encoded string
  return decodeURIComponent(encodedString);
}
