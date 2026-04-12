const N = parseInt(process.argv[2]);

if (N < 2) {
  process.exit(0);
}

function isPrime(n: number): boolean {
  if (n <= 1) return false;
  if (n === 2 || n === 3) return true;
  if (n % 2 === 0) return false;
  
  for (let i = 3; i * i <= n; i += 2) {
    if (n % i === 0) return false;
  }
  return true;
}

for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    console.log(i);
  }
}