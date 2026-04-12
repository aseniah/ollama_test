function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    name: string;
    price: number;
    quantity: number;
    expected: number;
}

function runTests(): void {
    const testCases: TestCase[] = [
        {
            name: 'Single item (quantity 9)',
            price: 10,
            quantity: 9,
            expected: 90 // 10 * 9
        },
        {
            name: 'Just under threshold (quantity 10)',
            price: 20,
            quantity: 10,
            expected: 180 // 20 * 10 * 0.9
        },
        {
            name: 'Just under 50 (quantity 49)',
            price: 30,
            quantity: 49,
            expected: 1323 // 30 * 49 * 0.9
        },
        {
            name: 'At threshold (quantity 50)',
            price: 40,
            quantity: 50,
            expected: 1500 // 40 * 50 * 0.75
        }
    ];

    testCases.forEach(testCase => {
        const result = calculateDiscount(testCase.price, testCase.quantity);
        const status = result === testCase.expected ? 'PASS' : 'FAIL';
        const details = `Expected: ${testCase.expected}, Got: ${result}`;
        console.log(`${status}: ${test case.name} ${details}`);
    });

    // Additional tests from command-line args if provided
    if (process.argv.length > 2) {
        const args = process.argv.slice(2);
        args.forEach((arg, index) => {
            const [priceArg, quantityArg] = arg.split(',');
            const price = parseFloat(priceArg);
            const quantity = parseFloat(quantityArg);
            
            if (!isNaN(price) && !isNaN(quantity)) {
                const result = calculateDiscount(price, quantity);
                const status = result === 0 ? 'PASS' : 'FAIL';
                console.log(`${status}: Price=${price}, Quantity=${quantity}, Result=${result}`);
            }
        });
    }

    // Summary
    console.log('');
    console.log('Testing complete.');
}

runTests();