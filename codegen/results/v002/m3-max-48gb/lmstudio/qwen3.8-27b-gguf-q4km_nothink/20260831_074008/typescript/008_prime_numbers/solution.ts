function isPrime(n: number): boolean {
  if (n < 2) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return false;
  for (let i = 3; i * i <= n; i += 2) {
    if (n % i === 0) return false;
  }
  return true;
}

function main() {
  const N = parseInt(process.argv[2], 10);
  if (isNaN(N)) {
    return;
  }
  
  const primes: number[] = [];
  for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
      primes.push(i);
    }
  }
  
  if (primes.length > 0) {
    console.log(primes.join('\n'));
  }
}

main();