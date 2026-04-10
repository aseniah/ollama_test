const arg = process.argv[2];
const n = parseInt(arg, 10);

if (isNaN(n)) {
  process.exit(1);
}

if (n < 2) {
  // No primes below 2, print nothing
  process.exit(0);
}

const isPrime = new Array<boolean>(n + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

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