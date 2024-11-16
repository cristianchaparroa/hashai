// 'use client'/
import { getFrameMetadata } from 'frog/web'
import type { Metadata } from 'next'
import Image from 'next/image'
import { Mermaid } from './components/Mermaid'

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
  const encodeValue = searchParams.hash; 

  console.log('encodeValue', encodeValue);
  
  return (
    <main className={styles.main}>
      <Mermaid encodeValue={encodeValue} />
    </main>
  )
}
