import fs from 'fs';
import path from 'path';
import { argv } from 'process';

if (argv.length !== 2) {
  console.log('Usage: node fibonacci.js <N>');
  process.exit(1);
}

const N = Number(argv[1]);

if (N < 1) {
  console.log('N must be a positive integer.');
  process.exit(1);
}

const fibonacci = [1, 1];

while (true) {
  const nextFib = fibonacci[fibonacci.length - 1] + fibonacci[fibonacci.length - 2];
  if (nextFib > N) {
    break;
  }
  fibonacci.push(nextFib);
}

console.log(...fibonacci);