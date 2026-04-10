function isPrime(n: number): boolean {
  if (n < 2) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return n === 2;

  const limit = Math.sqrt(n);
  for (let i = 3; i <= limit; i += 2) {
    if (n % i === 0) return false;
  }
  return true;
}

function main() {
  const n = parseInt(process.argv[2], 10);

  if (isNaN(n)) {
    console.error('Invalid input: please provide an integer argument');
    process.exit(1);
  } else if (n < 2) {
    return;
  }

  const primes: number[] = [];
  for (let i = 2; i <= n; i++) {
    if (isPrime(i)) {
      primes.push(i);
    }
  }

  for (const prime of primes) {
    console.log(prime);
  }
}

main();