const N = parseInt(process.argv[2]);

if (isNaN(N) || N < 2) {
    process.exit(0);
}

// Sieve of Eratosthenes
const sieve: boolean[] = new Array(N + 1).fill(true);
sieve[0] = false;
sieve[1] = false;

for (let i = 2; i * i <= N; i++) {
    if (sieve[i]) {
        for (let j = i * i; j <= N; j += i) {
            sieve[j] = false;
        }
    }
}

for (let i = 2; i <= N; i++) {
    if (sieve[i]) {
        console.log(i);
    }
}