-- Table: users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(40) UNIQUE,
    created_at TIMESTAMP NOT NULL,
    passwd VARCHAR(40) NOT NULL,
    email VARCHAR(50) NOT NULL,
    name VARCHAR(40),
    surname VARCHAR(40),
    phone VARCHAR(12) NOT NULL,
    age INTEGER NOT NULL,
    sex VARCHAR(13) NOT NULL,
    weight INTEGER,
    height INTEGER,
    BMI INTEGER
);

-- Table: recipes
CREATE TABLE recipes (
    id int SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    synopsis VARCHAR(100) NOT NULL,
    time INTEGER NOT NULL, -- Preparation time in minutes
    difficulty INTEGER NOT NULL
);

-- Table: reviews
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    recipe_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    review INTEGER NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: ingredients
CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL
);

-- Table: recipes_ingredients
CREATE TABLE recipes_ingredients (
    recipe_id INTEGER NOT NULL,
    ingredient_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    unit INTEGER DEFAULT 0,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);

-- Table: tags_types
CREATE TABLE tags_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL
);

-- Table: tags
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    type_id INTEGER NOT NULL,
    FOREIGN KEY (type_id) REFERENCES tags_types(id)
);

-- Table: recipes_tags
CREATE TABLE recipes_tags (
    recipe_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

-- Indexes
CREATE INDEX idx_reviews_recipe_id ON reviews (recipe_id);
CREATE INDEX idx_recipes_ingredients_recipe_id ON recipes_ingredients (recipe_id);
CREATE INDEX idx_tags_type_id ON tags (type_id);
CREATE INDEX idx_recipes_tags_recipe_id ON recipes_tags (recipe_id);
