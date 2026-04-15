// Parse N from command line
if (Args.Count == 0 || !int.TryParse(Args[0], out int n))
{
    // If no argument or invalid integer, do nothing (or you could handle it differently)
    // Based on requirements: "Accept a single integer argument N"
    // If missing/invalid, we can't proceed, so print nothing.
}
else
{
    if (n < 1)
    {
        // Print nothing
    }
    else
    {
        int a = 1;
        int b = 1;

        // Print first 1 if n >= 1
        if (a <= n)
        {
            Console.WriteLine(a);
        }

        // Print second 1 if n >= 1
        if (b <= n && b != a) // Avoid duplicate print if n=1 and we already printed 1
        {
            // But since sequence is 1,1, we should print both 1s up to n
            // Actually, the sequence is: 1, 1, 2, 3, 5...
            // So if n >= 1, we print 1.
            // If n >= 1, we print the second 1 as well? 
            // The problem says: "Start the sequence: 1, 1, 2, 3, 5, ..."
            // So we print both 1s if they are <= n.
            
            // Let's reset logic to be clear:
        }
    }
}

// Reset and write clear logic:
if (Args.Count > 0 && int.TryParse(Args[0], out int n) && n >= 1)
{
    int a = 1;
    int b = 1;

    // Print first 1
    Console.WriteLine(a);

    // If n >= 1, we also print the second 1
    if (b <= n)
    {
        Console.WriteLine(b);
    }

    while (true)
    {
        int next = a + b;
        if (next > n) break;
        Console.WriteLine(next);
        a = b;
        b = next;
    }
}