const [_, nStr] = process.argv;
const N = parseInt(nStr, 10);

if (isNaN(N) || N < 2) {
  process.exit(0);
}

function isPrime(num: number): boolean {
  if (num < 2) return false;
  if (num === 2 || num === 3) return true;
  if (num % 2 === 0 || num % 3 === 0) return false;
  
  let i = 5;
  while (i * i <= num) {
    if (num % i === 0 || num % (i + 2) === 0) return false;
    i += 6;
  }
  return true;
}

const primes: number[] = [];
const upperLimit = Math.floor(Math.sqrt(N));

for (let num = 2; num <= N; num++) {
  if (isPrime(num)) {
    primes.push(num);
  }
}

if (primes.length > 0) {
  for (const prime of primes) {
    console.log(prime);
  }
}

process.exit(0);