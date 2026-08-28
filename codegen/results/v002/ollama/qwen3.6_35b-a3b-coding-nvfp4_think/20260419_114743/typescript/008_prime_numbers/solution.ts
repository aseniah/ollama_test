const input = process.argv[2];
if (!input) process.exit(0);

const N = parseInt(input, 10);
if (isNaN(N)) process.exit(0);

if (N < 2) process.exit(0);

const isPrime = new Array(N + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

for (let i = 2; i * i <= N; i++) {
  if (isPrime[i]) {
    for (let j = i * i; j <= N; j += i) {
      isPrime[j] = false;
    }
  }
}

for (let i = 2; i <= N; i++) {
  if (isPrime[i]) {
    console.log(i);
  }
}