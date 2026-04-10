function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases = [
    { price: 10, quantity: 9, expected: 90 },
    { price: 10, quantity: 10, expected: 100 },
    { price: 10, quantity: 49, expected: 441 },
    { price: 10, quantity: 50, expected: 375 }
];

// Run test cases
testCases.forEach(testCase => {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    if (result === testCase.expected) {
        console.log(`PASS: calculateDiscount(${testCase.price}, ${testCase.quantity}) = ${result}`);
    } else {
        console.log(`FAIL: calculateDiscount(${testCase.price}, ${testCase.quantity}) = ${result} (expected: ${testCase.expected})`);
    }
});