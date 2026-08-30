// Accept N from command line arguments
if (Args.Count == 0) return;

if (!int.TryParse(Args[0], out int n)) return;

if (n < 1) return;

int a = 1;
int b = 1;

// Print the first 1
if (a <= n)
{
    Console.WriteLine(a);
}

// Print the second 1 if it's within range (and distinct from the first logic)
// Since the sequence is 1, 1, 2..., we need to handle the second 1 if n >= 1
// The first 1 is already printed above. If n >= 1, we print the second 1.
// However, if n < 1, we printed nothing.
// If n >= 1, we print the first 1. Then we check the second 1.
// Since both are 1, we just need to ensure we print it once per step in the loop logic or handle it explicitly.
// Let's restart the logic to be cleaner:
// Sequence: 1, 1, 2, 3, ...

// Reset
a = 1;
b = 1;

// Handle the first '1'
if (a <= n)
{
    Console.WriteLine(a);
}

// If n >= 1, we also need to print the second '1' (unless we treat the sequence as starting 1, 2... 
// but the prompt says 1, 1, 2, 3...)
// So if n >= 1, we print the second 1 as well.
if (b <= n)
{
    Console.WriteLine(b);
}

// Now generate the rest
while (true)
{
    int next = a + b;
    if (next > n) break;
    
    Console.WriteLine(next);
    a = b;
    b = next;
}