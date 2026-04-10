using System;

// Check if an argument was provided
if (Args.Count > 0)
{
    // Attempt to parse the first argument as a long to handle large integer values
    if (long.TryParse(Args[0], out long n) && n >= 1)
    {
        // Initialize the first two numbers of the Fibonacci sequence as specified: 1, 1, ...
        long current = 1;
        long next = 1;

        // Continue generating Fibonacci numbers as long as the current number does not exceed N
        while (current <= n)
        {
            // Print the current Fibonacci number
            Console.WriteLine(current);

            // Calculate the next number in the sequence
            long temp = current + next;
            
            // Shift the window: the old 'next' becomes the new 'current'
            current = next;
            
            // The 'temp' sum becomes the new 'next'
            next = temp;
        }
    }
}