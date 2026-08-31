const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 2) {
    process.exit(0);
}

// Sieve of Eratosthenes
const sieve = new Array(N + 1).fill(true);
sieve[0] = false;
sieve[1] = false;

for (let i = 2; i * i <= N; i++) {
    if (sieve[i]) {
        for (let j = i * i; j <= N; j += i) {
            sieve[j] = false;
        }
    }
}

const results: string[] = [];
for (let i = 2; i <= N; i++) {
    if (sieve[i]) {
        results.push(i.toString());
    }
}

if (results.length > 0) {
    console.log(results.join("\n"));
}