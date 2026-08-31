const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 2) {
  process.exit(0);
}

const isPrime = new Array<boolean>(N + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

for (let i = 2; i * i <= N; i++) {
  if (isPrime[i]) {
    for (let j = i * i; j <= N; j += i) {
      isPrime[j] = false;
    }
  }
}

const results: string[] = [];
for (let i = 2; i <= N; i++) {
  if (isPrime[i]) {
    results.push(i.toString());
  }
}

if (results.length > 0) {
  console.log(results.join("\n"));
}