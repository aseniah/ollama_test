function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTests(): void {
    const testCases = [
        { description: "quantity 9 (less than 10)", price: 10, quantity: 9, expected: 90 },
        { description: "quantity 10 (first discount tier)", price: 10, quantity: 10, expected: 90 },
        { description: "quantity 49 (second discount tier)", price: 10, quantity: 49, expected: 441 },
        { description: "quantity 50 (third discount tier)", price: 10, quantity: 50, expected: 375 },
        { description: "quantity 0 (edge case)", price: 10, quantity: 0, expected: 0 },
        { description: "quantity 100 (third discount tier)", price: 5, quantity: 100, expected: 375 }
    ];

    testCases.forEach(testCase => {
        const result = calculateDiscount(testCase.price, testCase.quantity);
        const pass = Math.abs(result - testCase.expected) < 0.0001; // Using small tolerance for floating point comparison
        
        if (pass) {
            console.log(`PASS: ${testCase.description}`);
        } else {
            console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
        }
    });
}

runTests();