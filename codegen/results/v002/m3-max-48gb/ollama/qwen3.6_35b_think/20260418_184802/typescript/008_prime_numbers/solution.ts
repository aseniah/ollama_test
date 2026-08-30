function main() {
  const nStr = process.argv[2];
  if (!nStr) return;

  const n = parseInt(nStr, 10);
  if (isNaN(n) || n < 2) return;

  // Sieve of Eratosthenes
  const isPrime = new Array(n + 1).fill(true);
  isPrime[0] = false;
  isPrime[1] = false;

  for (let p = 2; p * p <= n; p++) {
    if (isPrime[p]) {
      for (let i = p * p; i <= n; i += p) {
        isPrime[i] = false;
      }
    }
  }

  const primes: number[] = [];
  for (let p = 2; p <= n; p++) {
    if (isPrime[p]) {
      primes.push(p);
    }
  }

  if (primes.length > 0) {
    console.log(primes.join('\n'));
  }
}

main();