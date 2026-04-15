// Check if there is exactly one argument
if (Args.Count != 1)
{
    // No output required if argument count is invalid, but the prompt implies a single integer argument is expected.
    // If the argument is missing or invalid, we simply exit without printing anything based on the "If N < 1" logic.
    // However, if the user provides a non-integer, we should handle it gracefully.
    return;
}

// Parse the integer argument
if (!int.TryParse(Args[0], out int n))
{
    // If parsing fails, treat as no valid input, print nothing.
    return;
}

// If N < 1, print nothing
if (n < 1)
{
    return;
}

int a = 1;
int b = 1;

// Print the first 1 if it's within range
if (a <= n)
{
    Console.WriteLine(a);
}

// Generate subsequent numbers
while (true)
{
    int next = a + b;
    if (next > n)
    {
        break;
    }
    Console.WriteLine(next);
    
    // Move forward
    a = b;
    b = next;
}