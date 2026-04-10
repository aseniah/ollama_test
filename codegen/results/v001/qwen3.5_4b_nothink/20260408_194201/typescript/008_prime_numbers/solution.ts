const parseArgs = (args: string[]): number | undefined => {
  for (let i = 0; i < args.length; i++) {
    const arg = args[i];
    if (!arg) {
      return undefined;
    }
    const parsed = parseInt(arg, 10);
    if (isNaN(parsed)) {
      return undefined;
    }
    if (parsed === Number.MAX_SAFE_INTEGER) {
      return undefined;
    }
    if (parsed === Number.MIN_SAFE_INTEGER) {
      return undefined;
    }
    return parsed;
  }
  return undefined;
};

const primes = (N: number): number[] => {
  const sieve = new Uint8Array(N + 1);
  sieve[2] = 1;
  const limit = Math.floor(Math.sqrt(N));
  for (let i = 3; i <= limit; i += 2) {
    sieve[i] = 1;
  }
  for (let i = 3; i <= limit; i += 2) {
    if (sieve[i]) {
      const p = i;
      for (let j = p * p; j <= N; j += 2 * p) {
        sieve[j] = 0;
      }
    }
  }
  const result: number[] = [];
  for (let i = 2; i <= N; i++) {
    if (sieve[i]) {
      result.push(i);
    }
  }
  return result;
};

const N = parseArgs(process.argv.slice(1));
if (typeof N === 'undefined') {
  process.exit(1);
}
const p = primes(N);
for (const i of p) {
  console.log(i);
}