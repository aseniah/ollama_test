int n = int.Parse(Args[0]);

if (n < 1) {
    return; // Print nothing
}

// Generate Fibonacci numbers and print them until the next one would exceed n
int a = 1;
int b = 1;

bool firstPrinted = false;

while (true) {
    if (!firstPrinted) {
        Console.WriteLine(a);
        firstPrinted = true;
    }
    
    // If the current number 'a' is already > n, stop (though logic below handles this before printing next)
    // Actually, we need to ensure we print 'a' only if a <= n.
    // Let's restructure: 
    // 1. Print 'a' if a <= n.
    // 2. Calculate next Fibonacci number.
    // 3. If next number > n (or current logic fails), stop.
    
    int next = a + b;
    
    if (next > n) {
        break;
    }
    
    a = b;
    b = next;
}

return;