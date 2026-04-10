// The discount calculation function
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases with expected values
const testCases = [
    { description: 'single item (quantity 9)', price: 10, quantity: 9, expected: 90 },
    { description: 'first discount tier (quantity 10)', price: 10, quantity: 10, expected: 90 },
    { description: 'second discount tier (quantity 49)', price: 10, quantity: 49, expected: 441 },
    { description: 'first bulk discount tier (quantity 50)', price: 10, quantity: 50, expected: 375 }
];

// Run tests
for (const testCase of testCases) {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    const passed = result === testCase.expected;
    
    if (passed) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
    }
}

// Read optional arguments from process.argv (for custom test cases)
const customTests = process.argv.slice(2);
if (customTests.length > 0) {
    console.log('\nUsing custom test cases from arguments:');
    for (const arg of customTests) {
        if (!isNaN(parseInt(arg))) {
            const price = 10;
            const quantity = parseInt(arg);
            const result = calculateDiscount(price, quantity);
            const passed = result === 10 * quantity;
            console.log(passed ? `PASS: custom quantity ${quantity}` : `FAIL: custom quantity ${quantity} (expected: ${10 * quantity}, got: ${result})`);
        }
    }
}