function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases = [
    { price: 100, quantity: 5, description: "quantity 5 (less than 10)" },
    { price: 100, quantity: 9, description: "quantity 9 (boundary condition)" },
    { price: 100, quantity: 10, description: "quantity 10 (boundary condition)" },
    { price: 100, quantity: 49, description: "quantity 49 (boundary condition)" },
    { price: 100, quantity: 50, description: "quantity 50 (boundary condition)" },
    { price: 50, quantity: 25, description: "quantity 25 (between 10 and 49)" },
    { price: 20, quantity: 75, description: "quantity 75 (greater than 50)" }
];

let passedTests = 0;
let totalTests = testCases.length;

for (const testCase of testCases) {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    let expected: number;
    
    if (testCase.quantity < 10) {
        expected = testCase.price * testCase.quantity;
    } else if (testCase.quantity < 50) {
        expected = testCase.price * testCase.quantity * 0.9;
    } else {
        expected = testCase.price * testCase.quantity * 0.75;
    }
    
    if (result === expected) {
        console.log(`PASS: ${testCase.description}`);
        passedTests++;
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${expected}, got: ${result})`);
    }
}

console.log(`\nTest Results: ${passedTests}/${totalTests} tests passed`);