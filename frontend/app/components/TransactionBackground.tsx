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
        <h1 className={styles.title}>
          HAS# AI
        </h1>
        <div className={styles.innerContainer}>
          {children}
        </div>
        <div className={styles.flex}>
          <div className="leftAligned">
            <Link href="converse://home" className={`${styles.transactionModalButton} ${styles.secondary}`}>
              ← Back
            </Link>
          </div>
          <button
            className={styles.transactionModalButton}
            onClick={() => {
              alert('Transaction reported');
              window.location.href = 'converse://home';
            }}
          >
            Report transaction
          </button>
        </div>

      </div>
    </div>
  )
}
