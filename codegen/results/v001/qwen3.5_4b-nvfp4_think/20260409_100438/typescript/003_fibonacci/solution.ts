const N = parseInt(process.argv[2]);

if (N < 1) {
  process.exit(0);
}

let a: number = 1;
let b: number = 1;

const result: number[] = [];

result.push(a);
result.push(b);

while (b <= N) {
  const next = a + b;
  a = b;
  b = next;
  if (b <= N) {
    result.push(b);
  } else {
    break;
  }
}

for (const num of result) {
  console.log(num);
}