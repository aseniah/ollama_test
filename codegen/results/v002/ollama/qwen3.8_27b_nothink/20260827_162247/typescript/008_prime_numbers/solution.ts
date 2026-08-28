function isPrime(num: number): boolean {
  if (num < 2) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return false;
  for (let i = 3; i * i <= num; i += 2) {
    if (num % i === 0) return false;
  }
  return true;
}

function main() {
  const arg = process.argv[2];
  if (!arg) {
    return;
  }
  const n = parseInt(arg, 10);
  if (isNaN(n)) {
    return;
  }
  for (let i = 2; i <= n; i++) {
    if (isPrime(i)) {
      console.log(i);
    }
  }
}

main();