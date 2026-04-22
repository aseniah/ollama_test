function main() {
  const arg = process.argv[2];
  if (arg === undefined) return;

  const n = parseInt(arg, 10);
  if (isNaN(n)) return;

  for (let i = 2; i <= n; i++) {
    let isPrime = true;
    for (let j = 2; j * j <= i; j++) {
      if (i % j === 0) {
        isPrime = false;
        break;
      }
    }
    if (isPrime) {
      console.log(i);
    }
  }
}

main();