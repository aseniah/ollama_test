if (Args.Count != 1)
{
    Console.WriteLine("Usage: script.csx <N>");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.WriteLine("Invalid integer provided.");
    return;
}

if (n < 1)
{
    return;
}

long a = 1;
long b = 1;

// Print the first number
Console.WriteLine(a);

// Check if N is large enough to include the second '1' explicitly in the loop logic below,
// or handle it directly. Since sequence starts 1, 1, we print first '1' above.
// If n < 1, we already returned. So n >= 1.

while (true)
{
    long next = a + b;
    
    if (next > n)
    {
        break;
    }

    Console.WriteLine(next);
    a = b;
    b = next;
}