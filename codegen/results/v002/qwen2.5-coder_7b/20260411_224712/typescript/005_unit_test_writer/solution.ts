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

for (const testCase of testCases) {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    if (result === testCase.expected) {
        console.log(`PASS: Expected discount for price ${testCase.price} and quantity ${testCase.quantity} is ${testCase.expected}`);
    } else {
        console.log(`FAIL: Expected discount for price ${testCase.price} and quantity ${testCase.quantity} to be ${testCase.expected}, but got ${result}`);
    }
}