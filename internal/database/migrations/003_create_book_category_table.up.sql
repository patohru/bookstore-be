CREATE TABLE IF NOT EXISTS book_category(
    book_id uuid REFERENCES books(id),
    category_id uuid REFERENCES categories(id),
    PRIMARY KEY (book_id, category_id)
);
