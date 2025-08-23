CREATE TABLE IF NOT EXISTS book_category(
    book_id uuid REFERENCES books(book_id),
    category_id uuid REFERENCES categories(category_id),
    PRIMARY KEY (book_id, category_id)
);
