const arg = process.argv[2];
const N = parseInt(arg, 10);

if (isNaN(N) || N < 2) {
  process.exit(0);
}

const isPrime = new Array(N + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

for (let p = 2; p * p <= N; p++) {
  if (isPrime[p]) {
    for (let i = p * p; i <= N; i += p) {
      isPrime[i] = false;
    }
  }
}

for (let i = 2; i <= N; i++) {
  if (isPrime[i]) {
    console.log(i);
  }
}