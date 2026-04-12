import { readdirSync, statSync, join } from 'fs';
import path from 'path';
import {
  isFile,
  stat,
} from 'fs';

async function fibonacciSequence(maxLimit: number): Promise<void> {
  const fibonacciSequence: number[] = [1, 1];
  
  while (fibonacciSequence[fibonacciSequence.length - 1] <= maxLimit) {
    fibonacciSequence.push(fibonacciSequence[fibonacciSequence.length - 2] + fibonacciSequence[fibonacciSequence.length - 1]);
  }

  console.log(...fibonacciSequence);
}

async function main(): Promise<void> {
  const args = process.argv;
  
  if (args.length < 2) {
    console.error('Usage: node script.tsx <N>');
    return;
  }

  const N = parseInt(args[1]);

  if (N < 1) {
    console.log('N must be at least 1.');
    return;
  }

  await fibonacciSequence(N);
}

main().catch((error) => console.error('Error:', error));