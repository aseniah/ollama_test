function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Read command-line arguments to potentially set the price (default 100)
const price = process.argv[2] ? parseFloat(process.argv[2]) : 100;

// Test cases covering boundary conditions (9, 10, 49, 50)
const testCases = [
    { quantity: 9, expected: 900 },
    { quantity: 10, expected: 900 },
    { quantity: 49, expected: 4410 },
    { quantity: 50, expected: 3750 }
];

for (const testCase of testCases) {
    const result = calculateDiscount(price, testCase.quantity);
    const pass = Math.abs(result - testCase.expected) < 1; // Check for floating point tolerance
    
    if (pass) {
        console.log(`PASS: quantity: ${testCase.quantity}, price: ${price}`);
    } else {
        console.log(`FAIL: quantity: ${testCase.quantity}, price: ${price} (expected: ${testCase.expected}, got: ${result})`);
    }
}