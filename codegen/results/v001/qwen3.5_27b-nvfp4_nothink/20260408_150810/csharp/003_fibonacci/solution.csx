#r "System.Runtime"
#r "System.Text.Json" // Included as per instructions, though not strictly needed for this logic if only printing

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count != 1)
{
    Console.WriteLine("Usage: Please provide exactly one integer argument N.");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.WriteLine("Error: The provided argument is not a valid integer.");
    return;
}

if (n < 1)
{
    // Print nothing
    return;
}

long prev = 0;
long current = 1;

while (current <= n)
{
    Console.WriteLine(current);
    long next = prev + current;
    prev = current;
    current = next;
}