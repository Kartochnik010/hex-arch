CREATE TABLE IF NOT EXISTS arith_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date DATE NOT NULL,
    answer INTEGER NOT NULL,
    operation VARCHAR(255) NOT NULL
)