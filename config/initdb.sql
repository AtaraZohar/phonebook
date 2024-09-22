CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20),
    address TEXT
);

INSERT INTO contacts (first_name, last_name, phone_number, address) 
VALUES 
    ('John', 'Doe', '123-456-7890', '123 Elm Street'),
    ('Jane', 'Smith', '987-654-3210', '456 Oak Avenue'),
    ('Harry', 'Potter', '123-456-7890', '4 Privet Drive'),
    ('Hermione', 'Granger', '987-654-3210', '5th Avenue'),
    ('Ron', 'Weasley', '123-456-7890', '123 Oxford Street');