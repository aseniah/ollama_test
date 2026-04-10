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
testCases.forEach((testCase, index) => {
    const { price, quantity, expected } = testCase;
    const result = calculateDiscount(price, quantity);
    if (result === expected) {
        console.log(`PASS: Test case ${index + 1}: Price ${price}, Quantity ${quantity}`);
    } else {
        console.log(`FAIL: Test case ${index + 1}: Price ${price}, Quantity ${quantity} (expected: ${expected}, got: ${result})`);
    }
});