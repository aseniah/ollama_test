using System;
using System.Collections.Generic;

List<string> lines = new List<string>();
if (N >= 1) {
    int prev = 0;
    int current = 1;
    
    // Add the first Fibonacci number
    if (current <= N) {
        lines.Add(current.ToString());
    }
    
    // Generate subsequent numbers
    while (true) {
        int next = prev + current;
        
        if (next > N && current <= N) {
            // Break loop since the previous one was valid, but check if 'next' needs to be added? 
            // Wait, the problem says "up to and including". We already added 'current'. 
            // If 'next' > N, we stop. But what if prev < N <= next? That's fine, current is the last.
            // However, we must ensure we add all numbers up to the largest one <= N.
            // The sequence is: 1, 1, 2, 3...
            // My logic above added 1. Then 1+1=2. 
            // Let's restructure the loop to be clearer.
        }
        
        if (next > N) {
            break;
        }
        
        lines.Add(next.ToString());
        prev = current;
        current = next;
    }
}

string output = string.Join("\n", lines);
Console.WriteLine(output);