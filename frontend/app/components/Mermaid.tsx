'use client';

import React, { useEffect, useRef, useState } from 'react';
import mermaid from 'mermaid';

interface MermaidProps {
  diagram: string;
  className?: string;
}

const MermaidDiagram: React.FC<MermaidProps> = ({ diagram, className = '' }) => {
  const elementRef = useRef<HTMLDivElement>(null);
  const [isMounted, setIsMounted] = useState(false);

  useEffect(() => {
    setIsMounted(true);
  }, []);

  useEffect(() => {
    if (!isMounted) return;

    const renderDiagram = async () => {
      if (!elementRef.current) return;

      try {
        // Reset mermaid to ensure clean rendering
        mermaid.initialize({
          startOnLoad: false,
          theme: 'dark',
          securityLevel: 'loose',
        });

        // Clear previous content
        elementRef.current.innerHTML = `<div class="mermaid">${diagram}</div>`;

        // Render new diagram
        await mermaid.run({
          querySelector: '.mermaid',
        });

        const svg = elementRef.current.querySelector('svg');
        if (svg) {
          svg.style.display = 'block';
          svg.style.margin = '0 auto';
          // Ensure SVG maintains aspect ratio and is responsive
          svg.setAttribute('preserveAspectRatio', 'xMidYMid meet');
        }

      } catch (err) {
        console.error('Failed to render mermaid diagram:', err);
        if (elementRef.current) {
          elementRef.current.innerHTML = '<div class="text-red-500">Failed to render diagram. Please check your syntax.</div>';
        }
      }
    };

    renderDiagram();
  }, [diagram, isMounted]);

  // Don't render anything on the server
  if (!isMounted) {
    return <div className={`mermaid-wrapper ${className}`}>Loading...</div>;
  }

  return (
      <div className={`mermaid-wrapper ${className}`}>
        <div ref={elementRef} />
      </div>
  );
};

export default MermaidDiagram;
