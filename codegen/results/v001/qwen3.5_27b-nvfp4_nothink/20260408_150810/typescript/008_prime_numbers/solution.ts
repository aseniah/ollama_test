import * as fs from "fs";
import * as path from "path";

function main() {
  const args = process.argv.slice(2);

  if (args.length !== 1) {
    console.error("Usage: tsx script.ts <integer N>");
    process.exit(1);
  }

  const input = args[0].trim();
  const n = parseInt(input, 10);

  if (isNaN(n)) {
    console.error("Invalid number provided.");
    process.exit(1);
  }

  if (n < 2) {
    // No primes up to N, print nothing.
    return;
  }

  const primes: number[] = [];
  const isPrime: boolean[] = new Array(n + 1).fill(true);

  // 0 and 1 are not prime
  if (n >= 0) isPrime[0] = false;
  if (n >= 1) isPrime[1] = false;

  for (let i = 2; i * i <= n; i++) {
    if (isPrime[i]) {
      for (let j = i * i; j <= n; j += i) {
        isPrime[j] = false;
      }
    }
  }

  for (let i = 2; i <= n; i++) {
    if (isPrime[i]) {
      console.log(i);
    }
  }
}

main();