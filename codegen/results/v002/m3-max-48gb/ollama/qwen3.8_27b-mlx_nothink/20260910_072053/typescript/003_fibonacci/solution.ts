function main(): void {
  const N = parseInt(process.argv[2], 10);
  
  if (isNaN(N) || N < 1) {
    return;
  }
  
  let a = 1;
  let b = 1;
  
  // Print the first 1
  console.log(a);
  
  if (N === 1) {
    return;
  }
  
  // Print the second 1 and continue
  console.log(b);
  
  if (N === 1) {
    return;
  }
  
  while (true) {
    const next = a + b;
    if (next > N) {
      break;
    }
    console.log(next);
    a = b;
    b = next;
  }
}

main();