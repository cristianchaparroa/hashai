# Frontend Hash AI

## Overview

This is a Next.js React project set up for rendering and displaying diagrams using the `mermaid` library, alongside other supporting tools.

## Installation

To set up the project locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-folder>```


2. Install the required dependencies:

```bash
npm install
```


3. Start the development server:

```bash
npm run dev
```


4. Open your browser and navigate to:

```bash
http://localhost:3000/
```


## Available Scripts
```bash
npm run dev
```
Starts the development server on http://localhost:3000.
```bash
npm run build
```
Builds the application for production. The output will be generated in the .next folder.
```bash
npm run start
```
Starts the production server.
```bash
npm run deploy
```
Deploys the application to Vercel.

## Dependencies
Here are the key dependencies used in this project:

- Next.js (14.1.0): Framework for server-rendered React applications.
- React (18): Library for building user interfaces.
- Mermaid (11.4.0): Library for creating mermaid diagrams.

## Development Setup
This project uses the following development tools:

- Typescript (5): Strongly typed JavaScript.
- Vercel: For deployment and hosting.
- Diagram Rendering
The MermaidDiagram component enables the rendering of Mermaid diagrams within your application. Below is an example of how to use it:

```bash
'use client';

import React from 'react';
import MermaidDiagram from './MermaidDiagram';

const App = () => {
  const sampleDiagram = `graph TD
    A[Start] --> B{Is it working?}
    B -- Yes --> C[Great!]
    B -- No --> D[Fix it]
    D --> A`;

  return (
    <div>
      <h1>Mermaid Diagram Example</h1>
      <MermaidDiagram diagram={sampleDiagram} />
    </div>
  );
};

export default App;
```

If you want to try the Diagram Rendering, you can use it with code or pass an encoded URL like this:

```bash
https://hashtracker.vercel.app/?hash=graph+LR%0A0x3154...2C35--%3E%7C3.36+mETH%7C0xf418...EEEE%0A0x3154...2C35--%3E%7C613.41+%C2%B5ETH%7C0x6632...eeeE%0A0x2535...303b--%3E%7C46.15+mETH%7C0x3154...2C35%0A0x3154...2C35--%3E%7C5.00+mETH%7C0x82E0...Ec8A%0A0x7D1A...eBB0--%3E%7C10.12+mETH%7C0x3154...2C35%0A0x3154...2C35--%3E%7C4.00+mETH%7C0x9Ae7...0a29%0A0xEc56...df66--%3E%7C10.34+mETH%7C0x3154...2C35%0A0x3154...2C35--%3E%7C40.00+mETH%7C0xdAC1...1ec7
```

This will display a Mermaid diagram.

Notes:
Ensure the diagram syntax is correct to avoid rendering issues.
Use the diagram prop to provide your Mermaid diagram string.
Customize the className prop as needed for additional styling.

License
This project is licensed under the MIT License. See the LICENSE file for more details.

