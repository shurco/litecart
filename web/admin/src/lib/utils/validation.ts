export interface ValidationRule {
  field: string;
  validator: (value: any) => boolean;
  message: string;
}

export function validateFields(
  data: Record<string, any>,
  rules: ValidationRule[],
): Record<string, string> {
  const errors: Record<string, string> = {};

  for (const rule of rules) {
    const value = data[rule.field];
    if (!rule.validator(value)) {
      errors[rule.field] = rule.message;
    }
  }

  return errors;
}

export const validators = {
  required: (message = "This field is required") => ({
    validator: (value: any) =>
      value !== null && value !== undefined && value !== "",
    message,
  }),
  minLength: (min: number, message?: string) => ({
    validator: (value: string) => !value || value.length >= min,
    message: message || `Must be at least ${min} characters`,
  }),
  email: (message = "Invalid email format") => ({
    validator: (value: string) =>
      !value || /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value),
    message,
  }),
  positiveNumber: (message = "Must be a positive number") => ({
    validator: (value: number | string) => {
      const num = typeof value === "string" ? parseFloat(value) : value;
      return !num || (!isNaN(num) && num >= 0);
    },
    message,
  }),
};
