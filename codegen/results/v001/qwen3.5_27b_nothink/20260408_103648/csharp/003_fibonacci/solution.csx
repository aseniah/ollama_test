#r "System.Text.Json"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: dotnet script fib.csx <integer N>");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Error: Argument must be a valid integer.");
    return;
}

if (n < 1)
{
    return;
}

int a = 1;
int b = 1;

// Print the first number if it's within range
if (a <= n)
{
    Console.WriteLine(a);
}

// If N is 1, we are done because the next number is also 1, 
// but the sequence is 1, 1, 2... so we need to handle the second 1 separately
// or just loop while the current number is <= n.

// Let's restart logic to be cleaner:
// Sequence: 1, 1, 2, 3, 5...
// We print 1, then loop generating next.

// Actually, let's just generate iteratively.
// f1 = 1, f2 = 1
// while f1 <= n: print f1; temp = f1 + f2; f1 = f2; f2 = temp;

int f1 = 1;
int f2 = 1;

while (f1 <= n)
{
    Console.WriteLine(f1);
    int next = f1 + f2;
    f1 = f2;
    f2 = next;
}