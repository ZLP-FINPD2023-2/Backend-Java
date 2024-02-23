CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL CHECK (LENGTH(password) >= 8),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    patronymic VARCHAR(255),
    gender VARCHAR(40),
    birthday DATE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS goals(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL ,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_goals_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS budgets(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL ,
    title VARCHAR(255) NOT NULL,
    goal_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at timestamp,
    CONSTRAINT fk_budgets_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE,
    CONSTRAINT fk_budgets_goal_id
    FOREIGN KEY (goal_id)
    REFERENCES goals(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Transactions(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title varchar(255),
    date TIMESTAMP NOT NULL,
    amount DECIMAL NOT NULL,
    budget_from INT NOT NULL,
    budget_to INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_transactions_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE,
    CONSTRAINT fk_transactions_budget_from
    FOREIGN KEY (budget_from)
    REFERENCES budgets (id)
    ON DELETE CASCADE,
    CONSTRAINT fk_transactions_budget_to
    FOREIGN KEY (budget_to)
    REFERENCES budgets (id)
    ON DELETE CASCADE
);