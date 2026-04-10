import { readline } from 'readline';
import fs from 'fs';
import path from 'path';
import { resolve } from 'path';

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

const args = process.argv.slice(2);
const N = parseInt(args[0], 10);

function isPrime(num: number) {
  if (num <= 1) return false;
  if (num <= 3) return true;
  if (num % 2 === 0 || num % 3 === 0) return false;
  let i = 5;
  while (i * i <= num) {
    if (num % i === 0 || num % (i + 2) === 0) return false;
    i += 6;
  }
  return true;
}

for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    rl.writeLine(i);
  }
}

rl.close();