CREATE TABLE IF NOT EXISTS reviews(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    rate smallint NOT NULL DEFAULT 0 CHECK (rate BETWEEN 0 AND 5),
    book_id uuid NOT NULL REFERENCES books(book_id),
    user_id uuid NOT NULL REFERENCES users(user_id),
    content text NOT NULL,
    is_edited boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
