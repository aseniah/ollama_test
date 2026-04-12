// Copy the function
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases including boundary conditions
const testCases = [
    { price: 10, quantity: 9, expected: 90, description: "Quantity: 9 (no discount)" },
    { price: 10, quantity: 10, expected: 90, description: "Quantity: 10 (10% discount)" },
    { price: 10, quantity: 49, expected: 441, description: "Quantity: 49 (10% discount)" },
    { price: 10, quantity: 50, expected: 375, description: "Quantity: 50 (25% discount)" }
];

let allPassed = true;

for (const test of testCases) {
    const result = calculateDiscount(test.price, test.quantity);
    const passed = Math.abs(result - test.expected) < 0.0001;

    if (passed) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        allPassed = false;
    }
}

if (allPassed) {
    console.log("All tests passed!");
}