function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Read price from command-line argument, default to 100
const price = Number(process.argv[2] ?? 100);

// Define test cases
const tests = [
    { 
        description: "quantity 9 (Full Price)", 
        quantity: 9, 
        expected: price * 9 
    },
    { 
        description: "quantity 10 (90% Discount)", 
        quantity: 10, 
        expected: price * 10 * 0.9 
    },
    { 
        description: "quantity 49 (90% Discount)", 
        quantity: 49, 
        expected: price * 49 * 0.9 
    },
    { 
        description: "quantity 50 (75% Discount)", 
        quantity: 50, 
        expected: price * 50 * 0.75 
    },
];

// Run tests and output results
for (const test of tests) {
    const result = calculateDiscount(price, test.quantity);
    // Use strict equality check with a tiny tolerance for floating point safety
    if (Math.abs(result - test.expected) < 0.0001) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
}