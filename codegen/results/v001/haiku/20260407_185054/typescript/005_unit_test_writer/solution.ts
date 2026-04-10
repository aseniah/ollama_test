function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    price: number;
    quantity: number;
    expected: number;
    description: string;
}

const testCases: TestCase[] = [
    {
        price: 10,
        quantity: 9,
        expected: 90,
        description: "quantity < 10: no discount"
    },
    {
        price: 10,
        quantity: 10,
        expected: 90,
        description: "quantity = 10: 10% discount applies"
    },
    {
        price: 10,
        quantity: 49,
        expected: 441,
        description: "quantity = 49: 10% discount (not 25%)"
    },
    {
        price: 10,
        quantity: 50,
        expected: 375,
        description: "quantity = 50: 25% discount applies"
    },
    {
        price: 5,
        quantity: 0,
        expected: 0,
        description: "quantity = 0: no discount"
    },
    {
        price: 100,
        quantity: 100,
        expected: 7500,
        description: "quantity = 100: 25% discount"
    }
];

function runTests(): void {
    let passCount = 0;
    let failCount = 0;

    for (const test of testCases) {
        const result = calculateDiscount(test.price, test.quantity);
        if (result === test.expected) {
            console.log(`PASS: ${test.description}`);
            passCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
            failCount++;
        }
    }

    console.log(`\nTotal: ${passCount} passed, ${failCount} failed`);
}

runTests();
