const N = parseInt(process.argv[2], 10);

if (!Number.isNaN(N) && N >= 2) {
  const isPrime = (x: number): boolean => {
    if (x < 2) return false;
    if (x === 2) return true;
    if (x % 2 === 0) return false;
    for (let i = 3; i * i <= x; i += 2) {
      if (x % i === 0) return false;
    }
    return true;
  };

  for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
      console.log(i);
    }
  }
}