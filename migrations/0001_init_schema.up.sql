CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
     created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    
    UNIQUE(username)
);

CREATE TABLE suppliers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL,
     created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(email)
);

CREATE TABLE items (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    price BIGINT NOT NULL,
     created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(name)
);

CREATE TABLE purchasing (
    id VARCHAR(50) PRIMARY KEY,
    date  TIMESTAMP NOT NULL,
    supplier_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    grand_total BIGINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_purchasing_supplier
        FOREIGN KEY (supplier_id)
        REFERENCES suppliers(id)
        ON DELETE RESTRICT,
    
    CONSTRAINT fk_purchasing_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
       ON DELETE RESTRICT

);

CREATE TABLE purchasing_details(
    id VARCHAR(50) PRIMARY KEY,
    purchasing_id VARCHAR(50) NOT NULL,
    item_id VARCHAR(50) NOT NULL,
    qty INT NOT NULL,
    sub_total BIGINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_purchasing_details_purchasing
        FOREIGN KEY (purchasing_id)
        REFERENCES purchasing(id)
        ON DELETE RESTRICT,
    
    CONSTRAINT fk_purchasing_details_items
        FOREIGN KEY (item_id)
        REFERENCES items(id)
        ON DELETE RESTRICT
);

-- purchasing
CREATE INDEX idx_purchasing_supplier_id
ON purchasing(supplier_id);

CREATE INDEX idx_purchasing_user_id
ON purchasing(user_id);

-- purchasing_details
CREATE INDEX idx_pd_purchasing_id
ON purchasing_details(purchasing_id);

CREATE INDEX idx_pd_item_id
ON purchasing_details(item_id);
