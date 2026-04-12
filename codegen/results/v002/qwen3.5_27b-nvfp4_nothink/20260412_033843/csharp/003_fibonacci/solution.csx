if (Args.Count != 1)
{
    // No input provided, do nothing or handle as needed. 
    // Based on instructions, we assume valid single integer if provided.
}

if (int.TryParse(Args[0], out int n))
{
    if (n >= 1)
    {
        long a = 1;
        long b = 1;

        while (a <= n)
        {
            Console.WriteLine(a);
            
            long next = a + b;
            a = b;
            b = next;
        }
    }
}