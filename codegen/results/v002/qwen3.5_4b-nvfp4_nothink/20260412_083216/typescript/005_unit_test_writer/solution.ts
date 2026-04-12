function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases covering boundary conditions:
// 1. quantity = 9 (should use 0.0 multiplier)
// 2. quantity = 10 (should use 0.9 multiplier)
// 3. quantity = 49 (should use 0.9 multiplier)
// 4. quantity = 50 (should use 0.75 multiplier)

const testCases = [
    { name: 'quantity = 9', price: 10, expected: 90 },
    { name: 'quantity = 10', price: 10, expected: 90 },
    { name: 'quantity = 49', price: 10, expected: 441 },
    { name: 'quantity = 50', price: 10, expected: 375 },
];

let passedCount = 0;
let failedCount = 0;

for (const testCase of testCases) {
    const actualResult = calculateDiscount(testCase.price, testCase.quantity);
    const description = `price=${testCase.price}, quantity=${testCase.quantity}`;
    
    if (actualResult === testCase.expected) {
        console.log(`PASS: ${description}`);
        passedCount++;
    } else {
        console.log(`FAIL: ${description} (expected: ${testCase.expected}, got: ${actualResult})`);
        failedCount++;
    }
}

console.log(`\nResults: ${passedCount} passed, ${failedCount} failed.`);