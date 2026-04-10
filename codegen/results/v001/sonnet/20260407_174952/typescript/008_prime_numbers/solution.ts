const N = parseInt(process.argv[2], 10);

for (let n = 2; n <= N; n++) {
  let isPrime = true;
  for (let i = 2; i <= Math.sqrt(n); i++) {
    if (n % i === 0) {
      isPrime = false;
      break;
    }
  }
  if (isPrime) console.log(n);
}
