function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases = [
    { price: 100, quantity: 5, expected: 500, description: "quantity 5 (less than 10)" },
    { price: 100, quantity: 9, expected: 900, description: "quantity 9 (boundary case)" },
    { price: 100, quantity: 10, expected: 900, description: "quantity 10 (boundary case)" },
    { price: 100, quantity: 49, expected: 4410, description: "quantity 49 (boundary case)" },
    { price: 100, quantity: 50, expected: 3750, description: "quantity 50 (boundary case)" },
    { price: 50, quantity: 25, expected: 1125, description: "quantity 25 (middle range)" },
    { price: 200, quantity: 75, expected: 11250, description: "quantity 75 (highest discount)" }
];

let passedTests = 0;
let totalTests = testCases.length;

for (const testCase of testCases) {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    if (result === testCase.expected) {
        console.log(`PASS: ${testCase.description}`);
        passedTests++;
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
    }
}

console.log(`\nTest Results: ${passedTests}/${totalTests} tests passed`);