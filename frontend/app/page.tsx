// 'use client'/
import { getFrameMetadata } from 'frog/web'
import type { Metadata } from 'next'
import Image from 'next/image'
import MermaidDiagram from './components/Mermaid'
import TransactionBackground from './components/TransactionBackground'

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

  console.log('searchParams', searchParams.hash?.trim());

  //searchParams?.hash ? searchParams?.hash : 

  const diagramBase64 = searchParams?.hash ? searchParams.hash?.trim() : `Z3JhcGggTFIKICAgIDB4ODYxMS4uLjc3ODQtLT58My4yMCBtRVRIfDB4MjUzNS4uLjMwM2IKICAgIDB4ODYxMS4uLjc3ODQtLT58MjAuMDAgbUVUSHwweDAwMDAuLi5iRTU5CiAgICAweERGZDUuLi45NjNkLS0+fDIuMzggbUVUSHwweDg2MTEuLi43Nzg0CiAgICAweDg2MTEuLi43Nzg0LS0+fDEwLjAwIG1FVEh8MHg3Y0NELi4uNURGQQogICAgMHgyOEM2Li4uMWQ2MC0tPnw0Mi43MyBtRVRIfDB4ODYxMS4uLjc3ODQ=`;
  const decodedHash = decodeBase64(diagramBase64)

  return (
    <main className={styles.main}>
      <TransactionBackground>
        <MermaidDiagram diagram={decodedHash} />
      </TransactionBackground>
    </main>
  )
}


function decodeBase64(base64String: string): string {
  if (!base64String) return ''
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


