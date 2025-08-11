import { describe, it, expect } from 'vitest';

describe('flaky test', () => {
  it('should pass sometimes', () => {
    // This is a flaky test that sometimes passes and sometimes fails
    // The flakiness comes from using Math.random() which is non-deterministic
    const randomValue = Math.random();
    const shouldPass = randomValue > 0.5;
    
    // This assertion will fail when shouldPass is true (which happens ~50% of the time)
    // The error message matches: "expected true to be false"
    expect(shouldPass).toBe(false);
  });
});