import { argv } from 'process';

const parseArgument = (args: string[]) => {
  for (let i = 2; i < args.length; i++) {
    const num = parseInt(args[i], 10);
    if (Number.isNaN(num)) {
      process.exit(1);
    }
    return num;
  }
  process.exit(1);
};

const sievePrimes = (n: number): number[] => {
  if (n < 2) {
    return [];
  }
  
  const isPrime = new Array<boolean>(n + 1).fill(true);
  isPrime[0] = false;
  isPrime[1] = false;
  
  for (let i = 2; i <= Math.floor(Math.sqrt(n)); i++) {
    if (isPrime[i]) {
      for (let j = i * i; j <= n; j += i) {
        isPrime[j] = false;
      }
    }
  }
  
  const primes: number[] = [];
  for (let i = 2; i <= n; i++) {
    if (isPrime[i]) {
      primes.push(i);
    }
  }
  
  return primes;
};

const printPrimes = (primes: number[]) => {
  for (const prime of primes) {
    console.log(prime);
  }
};

const main = () => {
  const N = parseArgument(argv);
  const primes = sievePrimes(N);
  printPrimes(primes);
};

main();