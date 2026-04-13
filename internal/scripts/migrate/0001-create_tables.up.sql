CREATE TABLE products (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    user_name UUID NOT NULL,
    name TEXT NOT NULL,
    value NUMERIC(10,2) NOT NULL,
    image TEXT NOT NULL,
    stock INT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);
CREATE TABLE purchases (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    user_id UUID NOT NULL,
    value NUMERIC(10,2) NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE ratings (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    user_name TEXT NOT NULL,
    purchase_id UUID NOT NULL,
    product_id UUID NOT NULL,
    rating NUMERIC(2,1) NOT NULL CHECK (rating >= 0 AND rating <= 5),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_purchase FOREIGN KEY (purchase_id) REFERENCES purchases(id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT unique_purchase_rating UNIQUE (purchase_id)
);