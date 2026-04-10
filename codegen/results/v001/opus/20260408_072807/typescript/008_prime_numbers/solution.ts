const N = parseInt(process.argv[2], 10);

if (!isNaN(N) && N >= 2) {
  const sieve = new Uint8Array(N + 1).fill(1);
  sieve[0] = 0;
  sieve[1] = 0;

  for (let i = 2; i * i <= N; i++) {
    if (sieve[i]) {
      for (let j = i * i; j <= N; j += i) {
        sieve[j] = 0;
      }
    }
  }

  for (let i = 2; i <= N; i++) {
    if (sieve[i]) console.log(i);
  }
}
