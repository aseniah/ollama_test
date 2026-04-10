function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function testCalculateDiscount() {
    const testCases = [
        { description: "quantity < 10", price: 10, quantity: 9, expected: 90 },
        { description: "quantity = 10", price: 10, quantity: 10, expected: 100 },
        { description: "10 < quantity < 50", price: 10, quantity: 49, expected: 441 },
        { description: "quantity >= 50", price: 10, quantity: 50, expected: 375 }
    ];

    testCases.forEach(testCase => {
        const result = calculateDiscount(testCase.price, testCase.quantity);
        if (result === testCase.expected) {
            console.log(`PASS: ${testCase.description}`);
        } else {
            console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
        }
    });
}

testCalculateDiscount();