const N = parseInt(process.argv[2], 10);

if (N >= 1) {
  let a = 1, b = 1;
  while (a <= N) {
    console.log(a);
    [a, b] = [b, a + b];
  }
}
