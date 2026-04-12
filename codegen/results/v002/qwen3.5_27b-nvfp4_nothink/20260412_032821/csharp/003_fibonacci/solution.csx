if (Args.Count == 0 || !int.TryParse(Args[0], out int n))
{
    return; // Exit if no argument or invalid integer
}

if (n < 1)
{
    return; // Print nothing if N < 1
}

long a = 1;
long b = 1;

// Print the first number immediately since the sequence starts with 1
Console.WriteLine(a);

// Loop while the next Fibonacci number does not exceed n
while (b <= n)
{
    Console.WriteLine(b);
    long temp = a + b;
    a = b;
    b = temp;
}