-- Create employees table (extends users)
CREATE TABLE IF NOT EXISTS employees (
    id VARCHAR(36) PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    employee_code VARCHAR(50) UNIQUE NOT NULL,
    department VARCHAR(100),
    designation VARCHAR(100),
    date_of_joining DATE NOT NULL,
    salary_amount DECIMAL(12, 2) NOT NULL,
    salary_currency VARCHAR(3) DEFAULT 'USD',
    bank_account VARCHAR(100),
    tax_id VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create payroll_records table
CREATE TABLE IF NOT EXISTS payroll_records (
    id VARCHAR(36) PRIMARY KEY,
    employee_id VARCHAR(36) NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    gross_salary DECIMAL(12, 2) NOT NULL,
    deductions DECIMAL(12, 2) DEFAULT 0,
    net_salary DECIMAL(12, 2) NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('draft', 'approved', 'paid')),
    paid_on TIMESTAMP,
    payment_reference VARCHAR(100),
    created_by VARCHAR(36) REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create attendance table
CREATE TABLE IF NOT EXISTS attendance (
    id VARCHAR(36) PRIMARY KEY,
    employee_id VARCHAR(36) NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    check_in TIMESTAMP,
    check_out TIMESTAMP,
    status VARCHAR(20) NOT NULL CHECK (status IN ('present', 'absent', 'leave', 'holiday')),
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(employee_id, date)
);

-- Create indexes
CREATE INDEX idx_employees_department ON employees(department);
CREATE INDEX idx_employees_date_of_joining ON employees(date_of_joining);
CREATE INDEX idx_payroll_records_employee_id ON payroll_records(employee_id);
CREATE INDEX idx_payroll_records_period ON payroll_records(period_start, period_end);
CREATE INDEX idx_payroll_records_status ON payroll_records(status);
CREATE INDEX idx_attendance_employee_id ON attendance(employee_id);
CREATE INDEX idx_attendance_date ON attendance(date);

-- Create triggers
CREATE TRIGGER update_employees_updated_at BEFORE UPDATE ON employees
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_payroll_records_updated_at BEFORE UPDATE ON payroll_records
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
