CREATE TABLE IF NOT EXISTS "contents" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    excerpt VARCHAR(250) NOT NULL,
    description text NOT NULL,
    image text NULL,
    status VARCHAR(25) NOT NULL DEFAULT 'PUBLISH',
    tags VARCHAR(250) UNIQUE NOT NULL,
    created_by_id INT REFERENCES users(id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contents_created_by_id ON contents(created_by_id);
CREATE INDEX idx_categories_category_id ON contents(category_id);