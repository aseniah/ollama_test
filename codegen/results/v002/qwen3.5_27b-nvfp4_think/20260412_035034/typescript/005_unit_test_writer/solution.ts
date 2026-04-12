// Function definition provided in the task
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner setup
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

const testCases: TestCase[] = [
    // Boundary Condition: Just below first discount threshold (No Discount)
    { 
        description: "Quantity 9 (No Discount)", 
        price: 10, 
        quantity: 9, 
        expected: 90 
    },
    // Boundary Condition: Start of first discount tier (10% off)
    { 
        description: "Quantity 10 (10% Discount Threshold)", 
        price: 10, 
        quantity: 10, 
        expected: 90 
    },
    // Boundary Condition: Just below second discount threshold (10% off)
    { 
        description: "Quantity 49 (10% Discount Max)", 
        price: 10, 
        quantity: 49, 
        expected: 441 
    },
    // Boundary Condition: Start of second discount tier (25% off)
    { 
        description: "Quantity 50 (25% Discount Threshold)", 
        price: 10, 
        quantity: 50, 
        expected: 375 
    }
];

// Execute tests and print results to stdout
function runTests(): void {
    testCases.forEach((test) => {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Use a small epsilon for floating point comparison safety
        const epsilon = 0.0001;
        const passed = Math.abs(result - test.expected) < epsilon;

        if (passed) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        }
    });
}

// Run the test suite
runTests();