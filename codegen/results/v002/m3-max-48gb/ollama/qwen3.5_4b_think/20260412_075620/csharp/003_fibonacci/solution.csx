int n = int.Parse(Args[0]);

if (n < 1)
{
    // print nothing
}
else
{
    long a = 1L;
    long b = 1L;
    
    // Print first two Fibonacci numbers (1, 1)
    Console.WriteLine(a);
    Console.WriteLine(b);
    
    // Generate remaining Fibonacci numbers
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
}