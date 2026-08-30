import { writeFileSync } from 'node:fs';

const args = process.argv.slice(1);
if (args.length !== 1) {
  console.error('Usage: tsx script.ts <N>');
  process.exit(1);
}

const n = parseInt(args[0], 10);

if (isNaN(n) || n < 1) {
  // N must be >= 1 to print anything. If N < 1, print nothing.
  process.exit(0);
}

let a = 1;
let b = 1;

while (true) {
  if (a <= n) {
    console.log(a);
  } else {
    break;
  }

  const nextA = a + b;
  a = b;
  b = nextA;

  // Prevent overflow in case N is very large and b exceeds Number.MAX_SAFE_INTEGER
  if (!Number.isFinite(b)) {
    break;
  }
}

// If the loop broke because b became too large, we might miss one more value.
// However, since we only print when a <= n, and 'a' is always the previous Fibonacci number,
// this logic should be correct. The initial sequence starts: 1, 1, 2, 3...

// Let's trace carefully:
// Initial: a=1, b=1
// Iteration 1: print 1 (if 1<=n), nextA=2, a=1, b=2 -> Wait, logic above is slightly flawed for the first two 1s.
// The standard definition often starts with F0=0, F1=1. But the prompt says "Start the sequence: 1, 1, 2...".
// So we treat the first '1' as F1 and the second '1' as F2? Or just simply generate them in order.

// Let's refactor to match the sequence explicitly: 1, 1, 2, 3, 5...
let f1 = 1;
let f2 = 1;

const results: string[] = [];

if (f1 <= n) results.push(f1.toString());

// Generate next numbers until we exceed N
while (true) {
  let nextF = f1 + f2;
  
  if (nextF > n) break;
  results.push(nextF.toString());
  
  // Shift for the next iteration: new f1 is old f2, new f2 is sum
  // Wait, standard Fibonacci shift: F(n) = F(n-1) + F(n-2).
  // Here we have f1 (current) and f2 (next_to_check)? No.
  
  // Let's trace the prompt sequence: 1, 1, 2, 3, 5...
  // Step 0: print 1 (First term) -> store as current?
  // Actually, simpler approach:
  // Current = 1, Next = 1.
  // Loop:
  //   If Current <= N, print it.
  //   Calculate next_val = Current + Next.
  //   Shift: Next = Current; Current = next_val.
  
  // Trace again with this logic:
  // Init: cur=1, nxt=1
  // Check 1 <= n? Yes (n>=1). Print 1.
  // Calc next_val = 2.
  // Shift: nxt=1, cur=2.
  // Loop: Check 2 <= n? Yes. Print 2.
  // Calc next_val = 3.
  // Shift: nxt=2, cur=3.
  // Loop: Check 3 <= n? Yes. Print 3.
  // Calc next_val = 5.
  // Shift: nxt=3, cur=5.
  // Loop: Check 5 <= n? Yes. Print 5.
  
  // This works perfectly. The initial '1, 1' are handled by the first iteration printing the start value, 
  // then updating for the second '1'.
  
  // Wait, my trace shifted wrong.
  // Init: cur=1 (first 1), nxt=1 (second 1).
  // Iteration 1: Print cur (1). calc next_val = cur(1) + nxt(1) = 2. Update: cur becomes 2? No, that skips the second 1.
  
  // Correct Logic for "1, 1, 2, 3...":
  // We need two variables to represent the last two Fibonacci numbers printed so far? 
  // Or just maintain current and next in sequence order.
  
  // Let's restart with a standard generator for 1, 1, 2, 3...
  let fibA = 0; // placeholder for F-1 effectively
  let fibB = 1; // First actual term is 1
  
  if (fibB <= n) results.push(fibB.toString());
  
  // We need to generate the sequence explicitly.
  // Let's use a while loop that generates terms one by one.
  // Term 0: 0? No, prompt says start with 1, 1.
  
  // Let's try: 
  // current = 1, next = 1.
  // Loop: if current <= n: print current; new_next = current + next; current = next; next = new_next;
  // Trace:
  // Start: cur=1, nxt=1
  // Iter 1: 1<=n? Print 1. new=1+1=2. cur=1, nxt=2. -> This prints FIRST 1. But we need SECOND 1.
  
  // The issue is that the second '1' is equal to the first '1'. 
  // So if I print cur immediately, I print one 1. Then I update cur to be the NEXT number (which is 1). 
  // So the logic `cur = nxt` happens AFTER printing? 
  // Let's re-trace:
  // Start: a=0, b=1 (conceptually F0=0, F1=1) -> But prompt says start sequence 1, 1.
  // Let's just hardcode the start of the loop to handle the "1, 1" specifically if needed, 
  // or realize that standard Fibonacci 1, 2, 3... is missing one 1 compared to the prompt?
  // Standard: 0, 1, 1, 2, 3...
  // Prompt: 1, 1, 2, 3... (omits 0).
  
  // So we can generate standard Fibonacci starting from F1=1, F2=1.
  
  let f_prev = 1; // Represents the term before the current one? No.
  // Let's define: 
  // a = 1 (first term)
  // b = 1 (second term)
  // Then next is a+b = 2.
  // But we need to print both 'a' and 'b'.
  
  // Refined Algorithm:
  // Let a=1, b=1.
  // Loop while true:
  //   if (a <= n) print(a);
  //   else break;
  //   // Advance: we have printed 'a'. The next number to check is 'b'? 
  //   // Wait, if a=1 and b=1, we printed a. What do we set for the NEXT iteration's 'a'? 
  //   // It should be 'b' (which is 1). And the NEXT 'b' should be a+b (2).
  
  let f0 = 1;
  let f1 = 1;
  
  if (f0 <= n) {
    console.log(f0);
  } else {
    process.exit(0);
  }

  while (true) {
    // If the current 'b' is still valid and greater than or equal to the last printed?
    // Actually, just check if f1 <= n.
    // But wait, the sequence is 1, 1, 2, 3...
    // So we print f0 (1). Then we set new_f0 = f1 (1), new_f1 = f0+f1 (2).
    // Check if new_f0 <= n. Print it?
    
    let nextVal = f0 + f1;
    
    if (nextVal > n) {
      break;
    }
    
    f0 = f1;
    f1 = nextVal;
    
    if (f0 <= n) { // This should be implicit since we checked nextVal > n, so f0 must be <= n?
       // Wait. 
       // Case 1: N=1.
       // Start: f0=1, f1=1. Print 1.
       // Loop check: nextVal = 2. 2 > 1 (true). Break. Result: only "1". 
       // But prompt says "up to and including the largest... that does not exceed N". 
       // If N=1, max fib <= 1 is 1. 
       // Is the second '1' considered distinct? No, it's the same value. 
       // So output should be "1".
       // My logic: print f0 (1). Break loop. Output: "1". Correct.
       
       // Case 2: N=2.
       // Start: f0=1, f1=1. Print 1.
       // Loop: nextVal = 2. 2 > 2 (false). 
       // f0 becomes 1, f1 becomes 2.
       // Check if f0 <= n? (1 <= 2). Yes? No, the loop continues to the top of `while`.
       // Wait, my code structure inside loop doesn't print again immediately.
       // I need to print f0 at the end of the loop body or begin?
       
       // Let's restructure:
       // We have printed the current 'f0'. We need to calculate the next term and check if it is <= n.
       // If yes, print it (which becomes the new f0).
    }
    
    // Let's rewrite cleanly without confusion.
  }
}