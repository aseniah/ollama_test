function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases: { price: number; quantity: number; expected: number; description: string }[] = [
    { price: 10, quantity: 9, expected: 90, description: "quantity 9 (below 10, no discount)" },
    { price: 10, quantity: 10, expected: 90, description: "quantity 10 (10% discount)" },
    { price: 10, quantity: 49, expected: 441, description: "quantity 49 (10% discount)" },
    { price: 10, quantity: 50, expected: 375, description: "quantity 50 (25% discount)" },
    { price: 5, quantity: 0, expected: 0, description: "quantity 0 (below 10, no discount)" },
    { price: 20, quantity: 100, expected: 1500, description: "quantity 100 (25% discount)" },
];

let allPassed = true;

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    // Use a small epsilon for floating point comparison
    const epsilon = 1e-9;
    if (Math.abs(result - tc.expected) < epsilon) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
        allPassed = false;
    }
}

if (allPassed) {
    console.log("All tests passed!");
}