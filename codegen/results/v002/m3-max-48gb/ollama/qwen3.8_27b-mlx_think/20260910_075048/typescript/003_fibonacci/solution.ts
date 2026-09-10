const N = BigInt(process.argv[2]);

if (N < 1n) {
  // N < 1, print nothing
} else {
  let a: bigint = 1n;
  let b: bigint = 1n;
  while (a <= N) {
    console.log(a.toString());
    const next = a + b;
    a = b;
    b = next;
  }
}