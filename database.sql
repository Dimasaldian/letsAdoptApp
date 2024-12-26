-- Active: 1732702326143@@127.0.0.1@5432@letsadoptdb
-- Membuat tabel Users

-- Membuat tabel Pets (Tabel untuk hewan peliharaan)
CREATE TABLE Pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('dog', 'cat', 'bird', 'other')),
    breed VARCHAR(100),
    age INT,
    description TEXT,
    vaccinated BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Membuat tabel Adoptions
CREATE TABLE Adoptions (
    id_adopt SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    pet_id INT NOT NULL,
    reason TEXT NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected')),
    notification_sent BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (pet_id) REFERENCES Pets(id) ON DELETE CASCADE
);

-- Membuat tabel Admins
CREATE TABLE Admins (
    id SERIAL PRIMARY KEY,                          -- ID unik untuk setiap admin
    name VARCHAR(100) NOT NULL,                    -- Nama admin
    email VARCHAR(100) UNIQUE NOT NULL,            -- Email unik untuk admin
    password VARCHAR(255) NOT NULL,                -- Password admin (disarankan dalam bentuk hash)
    privileges TEXT,                               -- Hak istimewa admin
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu pembuatan record
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Waktu pembaruan record
);



-- Membuat tabel pet_images
CREATE TABLE pet_images (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES Pets(id) ON DELETE CASCADE
);

-- Membuat index untuk pet_id di tabel pet_images
CREATE INDEX idx_pet_id ON pet_images(pet_id);

-- Insert dummy data into Pets
INSERT INTO Pets (name, type, breed, age, description, vaccinated) VALUES
('Dam', 'cat', 'Golden Retriever', 5, 'Friendly and energetic cat.', TRUE);

-- Insert dummy data into Adoptions
INSERT INTO Adoptions (user_id, pet_id, reason, status) VALUES
(1, 1, 'Looking for a companion.', 'approved'),
(2, 2, 'I want a low-maintenance pet.', 'pending'),
(3, 3, 'I love birds and want to take care of them.', 'approved');

-- Insert dummy data into Admins
INSERT INTO Admins (user_id, privileges) VALUES
(2, 'Manage users, approve adoptions');



-- Insert dummy data into pet_images
INSERT INTO pet_images (pet_id, url) VALUES
(1, 'https://example.com/images/rex.jpg'),
(2, 'https://example.com/images/whiskers.jpg'),
(3, 'https://example.com/images/tweety.jpg');

ALTER TABLE notifications ALTER COLUMN adoption_id TYPE INT USING adoption_id::integer;

ALTER TABLE notifications DROP CONSTRAINT notifications_adoption_id_fkey;
ALTER TABLE notifications ADD CONSTRAINT notifications_adoption_id_fkey FOREIGN KEY (adoption_id) REFERENCES adoptions(id_adopt) ON DELETE CASCADE;

UPDATE pet_images
SET url = REPLACE(url, '/assets/img/', '/public/img/')
WHERE url LIKE '/assets/img/%';


ALTER TABLE Adoptions
ADD COLUMN name VARCHAR(100),
ADD COLUMN email VARCHAR(100);

ALTER TABLE Adoptions
DROP CONSTRAINT adoptions_user_id_fkey;

ALTER TABLE Adoptions
DROP COLUMN user_id;

DROP TABLE Users;

SELECT 
    conname AS constraint_name, 
    conrelid::regclass AS dependent_table 
FROM 
    pg_constraint 
WHERE 
    confrelid = 'Users'::regclass;

ALTER TABLE Notifications
DROP CONSTRAINT notifications_user_id_fkey;

ALTER TABLE Notifications
DROP CONSTRAINT fk_notifications_user;

ALTER TABLE Admins
DROP CONSTRAINT admins_user_id_fkey;

ALTER TABLE Admins
DROP CONSTRAINT fk_admins_user;

ALTER TABLE Admins
ADD COLUMN name VARCHAR(100) NOT NULL,
ADD COLUMN password VARCHAR(255) NOT NULL;

INSERT INTO Admins (name, email, password, privileges)
VALUES 
('Super Admin', 'admin@example.com', 'hashedpassword123', 'Manage Users, Approve Adoptions'),
('Alice Smith', 'alice.admin@example.com', 'hashedpassword456', 'Approve Adoptions');

ALTER TABLE pets
ADD COLUMN negara VARCHAR(20) NOT NULL DEFAULT 'Indonesia';
