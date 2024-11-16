'use client'

import React from 'react'
import Link from 'next/link'
import styles from '../page.module.css'

interface TransactionBackgroundProps {
  children: React.ReactNode
}

export default function TransactionBackground({ children }: TransactionBackgroundProps) {
  return (
    <div className={styles.container}>
      <div className={styles.dotPattern} />
      <div className={styles.content}>
        <div className="leftAligned">
          <Link href="converse://home" className={`${styles.transactionModalButton} ${styles.secondary}`}>
            ← Back
          </Link>
        </div>
        <div className={styles.innerContainer}>
          {children}
        </div>
        <button className={styles.transactionModalButton} onClick={() => alert('Transaction reported')}>
          Report transaction
        </button>
      </div>
    </div>
  )
}