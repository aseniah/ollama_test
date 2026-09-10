function isPrime(num: number): boolean {
  if (num < 2) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return false;
  for (let i = 3; i <= Math.sqrt(num); i += 2) {
    if (num % i === 0) return false;
  }
  return true;
}

const N = parseInt(process.argv[2], 10);
if (isNaN(N)) {
  process.exit(1);
}

for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    console.log(i);
  }
}