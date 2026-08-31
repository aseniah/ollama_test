using System;
using System.Collections.Generic;

int N = Args.Count > 0 ? int.Parse(Args[0]) : 0;

if (N < 1)
{
    return;
}

List<int> fibonacci = new List<int>();
int a = 1, b = 1;

// Add first two numbers
fibonacci.Add(a);
if (b <= N)
{
    fibonacci.Add(b);
}

// Generate rest of Fibonacci sequence
while (true)
{
    int next = a + b;
    if (next > N)
        break;
    
    fibonacci.Add(next);
    a = b;
    b = next;
}

// Print each Fibonacci number
foreach (int num in fibonacci)
{
    Console.WriteLine(num);
}