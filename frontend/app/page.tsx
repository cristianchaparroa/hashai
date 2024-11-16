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
  // console.log('hash', hash);
  const flowchartExample = `
graph TD
    A[Start] --> B{Is it?}
    B -- Yes --> C[OK]
    C --> D[Rethink]
    D --> B
    B -- No ----> E[End]
  `;

  console.log('encodeValue', flowchartExample);

  return (
    <main className={styles.main}>
      <MermaidDiagram diagram={flowchartExample} />
    </main>
  )
}
