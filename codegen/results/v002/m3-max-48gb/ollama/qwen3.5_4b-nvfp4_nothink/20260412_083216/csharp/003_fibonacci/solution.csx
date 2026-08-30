using System;
using System.Collections.Generic;

// Parse command-line argument
int N = int.Parse(Args[0]);

// Handle edge case: if N < 1, print nothing
if (N < 1) {
    return;
}

// Start sequence with first two numbers: 1 and 1
List<int> fibonacci = new List<int> { 1, 1 };
int prev = 1;
int current = 1;

// Generate Fibonacci numbers up to N
while (true) {
    int next = prev + current;
    
    // If next exceeds N, we are done with generation
    if (next > N) {
        break;
    }
    
    // Add the new number to our list
    fibonacci.Add(next);
    
    // Update variables for next iteration
    prev = current;
    current = next;
}

// Print all generated Fibonacci numbers, one per line
foreach (int num in fibonacci) {
    Console.WriteLine(num);
}