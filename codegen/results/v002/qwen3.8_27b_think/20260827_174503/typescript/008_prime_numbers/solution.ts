function isPrime(n: number): boolean {
  if (n < 2) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return false;
  for (let i = 3; i <= Math.sqrt(n); i += 2) {
    if (n % i === 0) return false;
  }
  return true;
}

const N = parseInt(process.argv[2], 10);

for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    console.log(i);
  }
}