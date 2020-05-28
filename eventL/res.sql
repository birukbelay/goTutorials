create table events (
    id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(255) NOT NULL,
    description text,
    image varchar(255)
); 

create table menus (
    
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(255) NOT NULL,
    price numeric NOT NULL DEFAULT 0,
    description text,
    image varchar(255)

);

create table menu_categories (

    menu_id integer REFERENCES menus(id) ON UPDATE CASCADE ON DELETE CASCADE,
    category_id integer REFERENCES categories(id) ON UPDATE CASCADE,
    CONSTRAINT menu_category_pky PRIMARY KEY (menu_id, category_id) 

);

create table ingredients (

    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(255) NOT NULL,
    description varchar(255)

);

create table menu_ingredients (

    menu_id integer REFERENCES menus(id) ON UPDATE CASCADE ON DELETE CASCADE,
    ingredient_id integer REFERENCES ingredients(id) ON UPDATE CASCADE,

    CONSTRAINT menu_ingredient_pkey PRIMARY KEY (menu_id, ingredient_id)

);

create table users (

    id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    uuid varchar(64) NOT NULL UNIQUE,
    full_name varchar(255),
    email varchar(255) NOT NULL UNIQUE,
    phone varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL

);

create table roles (

    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(255) NOT NULL UNIQUE

);

create table user_roles (

    user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    role_id integer REFERENCES roles(id) ON UPDATE CASCADE,

    CONSTRAINT user_role_pkey PRIMARY KEY (user_id, role_id)

);

create table orders (

    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    placed_at timestamp NOT NULL,
    user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    menu_id integer REFERENCES menus(id) ON UPDATE CASCADE

 );

 create table comments (

     id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
     full_name varchar(255),
     message text NOT NULL,
     phone varchar(255),
     email varchar(255),
     posted_at timestamp NOT NULL

 );