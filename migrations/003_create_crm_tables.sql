-- Create customers table
CREATE TABLE IF NOT EXISTS customers (
    id VARCHAR(36) PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    contact_name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(50),
    address TEXT,
    industry VARCHAR(100),
    status VARCHAR(20) NOT NULL CHECK (status IN ('lead', 'prospect', 'active', 'inactive')),
    assigned_to VARCHAR(36) REFERENCES users(id),
    created_by VARCHAR(36) REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create customer_interactions table
CREATE TABLE IF NOT EXISTS customer_interactions (
    id VARCHAR(36) PRIMARY KEY,
    customer_id VARCHAR(36) NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    user_id VARCHAR(36) NOT NULL REFERENCES users(id),
    interaction_type VARCHAR(50) NOT NULL CHECK (interaction_type IN ('call', 'email', 'meeting', 'note')),
    subject VARCHAR(255),
    notes TEXT,
    interaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_customers_status ON customers(status);
CREATE INDEX idx_customers_assigned_to ON customers(assigned_to);
CREATE INDEX idx_customer_interactions_customer_id ON customer_interactions(customer_id);
CREATE INDEX idx_customer_interactions_user_id ON customer_interactions(user_id);
CREATE INDEX idx_customer_interactions_date ON customer_interactions(interaction_date);

-- Create triggers
CREATE TRIGGER update_customers_updated_at BEFORE UPDATE ON customers
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
