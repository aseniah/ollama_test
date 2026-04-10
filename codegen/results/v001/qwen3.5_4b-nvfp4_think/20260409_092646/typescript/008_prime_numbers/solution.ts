const n = parseInt(process.argv[2]);

if (isNaN(n) || Number.isNaN(n)) {
  process.exit(0);
}

function isPrime(num: number): boolean {
  if (num < 2) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return num !== 2;

  for (let i = 3; i <= Math.sqrt(num); i += 2) {
    if (num % i === 0) return false;
  }
  return true;
}

for (let i = 2; i <= n; i++) {
  if (isPrime(i)) {
    console.log(i.toString());
  }
}