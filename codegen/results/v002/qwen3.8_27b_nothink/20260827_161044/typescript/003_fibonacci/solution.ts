const args = process.argv.slice(2);

if (args.length < 1) {
  // No argument provided
  return;
}

const N = parseInt(args[0], 10);

if (isNaN(N) || N < 1) {
  // N < 1 or invalid, print nothing
  return;
}

const fibs: number[] = [1, 1];
let a = 1, b = 1;
while (true) {
  const next = a + b;
  if (next > N) break;
  fibs.push(next);
  a = b;
  b = next;
}

// Print all fib numbers
for (const f of fibs) {
  console.log(f);
}