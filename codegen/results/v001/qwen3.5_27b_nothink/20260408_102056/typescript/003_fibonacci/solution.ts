const main = () => {
  const args = process.argv.slice(2);
  
  if (args.length === 0) {
    return;
  }

  const n = parseInt(args[0], 10);

  if (isNaN(n) || n < 1) {
    return;
  }

  let a = 1;
  let b = 1;

  // Print the first number if it's within range
  if (a <= n) {
    console.log(a);
  } else {
    return;
  }

  // Continue generating Fibonacci numbers
  while (b <= n) {
    console.log(b);
    const next = a + b;
    a = b;
    b = next;
  }
};

main();