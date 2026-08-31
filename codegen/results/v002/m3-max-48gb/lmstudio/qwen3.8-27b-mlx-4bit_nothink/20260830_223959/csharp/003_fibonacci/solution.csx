// Read the command line argument
if (Args.Count == 0)
{
    Console.Error.WriteLine("Usage: Provide a single integer argument N");
    return;
}

if (!int.TryParse(Args[0], out int N))
{
    Console.Error.WriteLine("Error: Argument must be an integer");
    return;
}

// If N < 1, print nothing
if (N < 1)
{
    return;
}

// Generate Fibonacci numbers up to N
int a = 1, b = 1;

// Print the first number
Console.WriteLine(a);

// If N >= 1, we need to handle the second 1
if (N >= 1)
{
    Console.WriteLine(b);
}

// Continue generating until we exceed N
while (true)
{
    int next = a + b;
    if (next > N)
        break;
    Console.WriteLine(next);
    a = b;
    b = next;
}